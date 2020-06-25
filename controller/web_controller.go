package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-archive/key_custody_web/repository"
	"github.com/sslab-archive/key_custody_web/services"
	"log"
	"github.com/sslab-archive/key_custody_web/model"
	"fmt"
	"strconv"
	"strings"
	"go.dedis.ch/kyber/v3/share"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

func RegisterUserController(router *gin.Engine) {

	keysRouter := router.Group("keys")
	{
		keysRouter.GET("/index", func(context *gin.Context) {
			var keyDaos = []model.KeyPairDao{}
			keyPairs := repository.KeyPairs
			for _, keyPair := range keyPairs {
				if repository.CheckKeyIsRegistered(keyPair.PublicKeyAddress) {
					keyDaos = append(keyDaos, model.KeyPairDao{keyPair.PublicKeyAddress, keyPair.PublicKey[0:40] + "...", keyPair.PrivateKey, true})
				} else {
					keyDaos = append(keyDaos, model.KeyPairDao{keyPair.PublicKeyAddress, keyPair.PublicKey[0:40] + "...", keyPair.PrivateKey, false})
				}
			}

			context.HTML(http.StatusOK, "index.tmpl", gin.H{
				"keyDaos": keyDaos,
			})
		})

		keysRouter.GET("/generate", func(context *gin.Context) {
			keyPair, err := services.GenerateKeyPair()
			if err != nil {
				log.Println(err)
			}

			context.HTML(http.StatusOK, "key_generate.tmpl", gin.H{
				"publicKeyAddress": keyPair.PublicKeyAddress,
				"publicKey":        keyPair.PublicKey,
				"privateKey":       keyPair.PrivateKey,
			})
		})

		keysRouter.POST("/generate", func(context *gin.Context) {
			//생성된거 저장.
			// 리다이렉트 to request 화면.
			keyPair, err := services.GenerateKeyPair()
			if err != nil {
				log.Println(err)
			}
			// 키 저장.
			repository.StoreKey(keyPair.PublicKeyAddress, keyPair.PublicKey, keyPair.PrivateKey)

			context.JSON(http.StatusOK, gin.H{
				"publicKeyAddress": keyPair.PublicKeyAddress,
				"privateKey":       keyPair.PrivateKey,
			})
		})

		keysRouter.POST("/submit", func(context *gin.Context) {
			// 데이터 받아서 저장
			_, result := context.GetPostForm("publicKeyAddress")
			if !result {
				log.Println("NO public Key Param")
			}
			//repository.ProviderResponseMap
		})

		keysRouter.GET("/blockchain/list", func(context *gin.Context) {

			publicKeyAddress := context.Request.URL.Query()["publicKeyAddress"][0]
			keyPair := repository.GetKeyPairByPublicKeyAddress(publicKeyAddress)
			var partialKeys []string
			for _, val := range repository.UserPartialKeyMap[publicKeyAddress].PartialKeyProviderEntities {
				partialKeys = append(partialKeys, val.V.String())
				fmt.Println(val.V.String())
			}

			context.HTML(http.StatusOK, "key_blockchain_list.tmpl", gin.H{
				"publicKeyAddress": keyPair.PublicKeyAddress,
				"publicKey":        keyPair.PublicKey,
				"privateKey":       keyPair.PrivateKey,
				"partialKeys":      partialKeys,
			})

		})

		keysRouter.GET("/restore", func(context *gin.Context) {
			publicKeyAddress := repository.GetCurrentPublicKeyAddress()
			keyPair := repository.GetKeyPairByPublicKeyAddress(publicKeyAddress)
			var providerIds []int
			var providerData []model.ProviderResponseData
			for key, val := range repository.ProviderResponseMap {
				fmt.Println(key, val)
				for key, val := range val.ProviderResponseDatas {
					providerIds = append(providerIds, key)
					providerData = append(providerData, val)
				}
			}
			providers, err := services.GetAllProviderList()
			log.Println(providers)
			if err != nil {
				log.Println(err)
			}

			var partialKeys []string
			for _, value := range repository.RestoreProviderResponseMap[publicKeyAddress].ProviderResponseDatas {
				partialKeys = append(partialKeys, value.PartialKey.V.String())
			}

			// 여기서 다보여주는 것은 어떨까?
			context.HTML(http.StatusOK, "key_restore.tmpl", gin.H{
				"publicKeyAddress": publicKeyAddress,
				"publicKey":   keyPair.PublicKey,
				"providers":   providers,
				"partialKeys": partialKeys,
				"redirectUrl": "http://localhost:8080/callback",
			})
		})

	}

	providerViewRouter := router.Group("providers")
	{
		providerViewRouter.GET("list", func(context *gin.Context) {
			// Provider list 보여주는 뷰 띄울 것
			//publicKeyAddress := context.Request.URL.Query()["publicKeyAddress"][0]
			// 프로바이더 리스트 뷰로 띄워주기
			providerList, err := services.GetAllProviderList()
			if err != nil {
				log.Println(err)
			}
			context.HTML(http.StatusOK, "providers_list.tmpl", gin.H{
				"providerList": providerList,
			})
		})

		providerViewRouter.GET("registry", func(context *gin.Context) {
			// 프로바이더 리스트 등록 하기...

			providerList, err := services.GetAllProviderList()
			if err != nil {
				log.Println(err)
			}

			context.HTML(http.StatusOK, "providers_registry.tmpl", gin.H{
				"providerList": providerList,
			})
		})

		providerViewRouter.POST("auth", func(context *gin.Context) {
			// 프로바이더 리스트 등록 하기...
			context.Request.ParseForm()
			var publicKeyAddress string
			var providerIds []int
			for key, value := range context.Request.PostForm {
				fmt.Println(key, value)
				if key == "publicKeyAddress" {
					publicKeyAddress = value[0]
				} else {
					for _, val := range value {
						id, _ := strconv.Atoi(val)
						providerIds = append(providerIds, id)
					}
				}
			}
			queryString := "publicKeyAddress=" + publicKeyAddress + "&providers=" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(providerIds)), ","), "[]")

			context.Redirect(http.StatusMovedPermanently, "auth/index?"+queryString)
		})

		providerViewRouter.GET("auth/index", func(context *gin.Context) {
			publicKeyAddress := context.Request.URL.Query()["publicKeyAddress"][0]
			providers := context.Request.URL.Query()["providers"][0]
			ids := strings.Split(providers, ",")
			var providerIds []int
			for _, strId := range ids {
				intId, err := strconv.Atoi(strId)
				if err != nil {
					panic(err)
				}
				providerIds = append(providerIds, intId)
			}

			providerNumber := len(providerIds)
			keyPair := repository.GetKeyPairByPublicKeyAddress(publicKeyAddress)
			partialKeys := services.GeneratePartialKey(keyPair.PrivateKey, providerNumber)

			repository.StorePartialKey(publicKeyAddress, providerIds, partialKeys)

			var providerAuthDtos = []model.ProviderAuthDTO{}
			providerList, _ := services.GetProviderListByIds(providerIds)

			for i, provider := range providerList {
				tempDto := model.ProviderAuthDTO{ID: provider.ID, Name: provider.Name, EndpointUrl: provider.EndpointUrl, PartialKey: partialKeys[i].V.String(), PartialKeyIndex: partialKeys[i].I}
				providerAuthDtos = append(providerAuthDtos, tempDto)
			}

			context.HTML(http.StatusOK, "providers_auth_index.tmpl", gin.H{
				"publicKey":        keyPair.PublicKey,
				"publicKeyAddress": publicKeyAddress,
				"providerDtos":     providerAuthDtos,
				"redirectUrl":      "http://localhost:8080/callback",
			})
		})

		providerViewRouter.POST("onReady", func(context *gin.Context) {
			// Provider들로부터 응답을 다 받아야함. Response를 관리해야함.
			publicKeyAddress, _ := context.GetPostForm("publicKeyAddress")
			if repository.CheckProviderResponseIsReady(publicKeyAddress) {
				context.JSON(http.StatusOK, gin.H{
					"result": true,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					// false
					"result": false,
				})
			}
		})

		providerViewRouter.POST("checkResponse", func(context *gin.Context) {
			publicKeyAddress, _ := context.GetPostForm("publicKeyAddress")
			providerId, _ := context.GetPostForm("providerId")
			intProviderId, _ := strconv.Atoi(providerId)

			currentType, _ := context.GetPostForm("type")
			//generate, restore인지 구분.
			if currentType == "generate" {
				if repository.CheckProviderResponse(publicKeyAddress, intProviderId) {
					data := repository.ProviderResponseMap[publicKeyAddress].ProviderResponseDatas[intProviderId]
					context.JSON(http.StatusOK, gin.H{
						"result": true,
						"data":   data,
					})
				} else {
					context.JSON(http.StatusOK, gin.H{
						"result": false,
					})
				}
			} else {
				if repository.CheckRestoreProviderResponse(publicKeyAddress, intProviderId) {
					data := repository.RestoreProviderResponseMap[publicKeyAddress].ProviderResponseDatas[intProviderId]
					context.JSON(http.StatusOK, gin.H{
						"result": true,
						"partialKey":   data.PartialKey.V.String(),
					})
				} else {
					context.JSON(http.StatusOK, gin.H{
						"result": false,
					})
				}
			}
		})

		providerViewRouter.POST("restorePrivateKey", func(context *gin.Context) {
			//repository.
			providerNumber := repository.RestoreProviderResponseMap[repository.GetCurrentPublicKeyAddress()].ProviderNumber
			log.Println(providerNumber)
			threshold := repository.RestoreProviderResponseMap[repository.GetCurrentPublicKeyAddress()].Threshold
			log.Println(threshold)

			prishares := services.GetRestorePartialKeys(repository.RestoreProviderResponseMap[repository.GetCurrentPublicKeyAddress()].ProviderResponseDatas)
			if len(prishares) >= threshold {
				privateKey := services.RestorePartialKey(prishares, providerNumber, threshold)
				fmt.Println(privateKey)
				context.JSON(http.StatusOK, gin.H{
					// false
					"result":     true,
					"privateKey": privateKey,
				})
			}else{
				context.JSON(http.StatusOK, gin.H{
					// false
					"result":     false,
				})
			}
		})

	}

	callbackRouter := router.Group("callback")
	{
		callbackRouter.GET("", func(context *gin.Context) {
			// Refresh Confirm 추가.
			purpose := context.Request.URL.Query()["purpose"][0]
			if purpose == "encrypt" {
				encryptedPayload := context.Request.URL.Query()["encrypted_payload"][0]
				payload := context.Request.URL.Query()["payload"][0]
				credentialType := context.Request.URL.Query()["credential_type"][0]
				//partialKey := context.Request.URL.Query()["partial_key"][0]
				encryptedPartialKey := context.Request.URL.Query()["encrypted_partial_key"][0]
				signedByPrivateKey := context.Request.URL.Query()["signed_by_private_key"][0]
				//providerPublicKey := context.Request.URL.Query()["public_key"][0]
				providerId := context.Request.URL.Query()["provider_id"][0]

				intProviderId, _ := strconv.Atoi(providerId)
				// 여기서 payload 블록체인에 올릴까..?
				tempResponseData := model.ProviderResponseData{
					ProviderId:          intProviderId,
					Payload:             payload,
					CredentialType:      credentialType,
					EncryptedPartialKey: encryptedPartialKey,
					SignedByPrivateKey:  signedByPrivateKey,
					EncryptedPayload:    encryptedPayload,
				}
				data := model.UserPartialKeyDto{EncryptedPartialKey: encryptedPartialKey, EncryptedPayload: encryptedPayload}
				repository.ProviderResponseMap[repository.GetCurrentPublicKeyAddress()].ProviderResponseDatas[intProviderId] = tempResponseData
				err := services.RegisterPartialKey(data)
				if err != nil {
					log.Println(err)
					return
				}
			} else {
				providerId := context.Request.URL.Query()["provider_id"][0]
				intProviderId, _ := strconv.Atoi(providerId)
				partialKey := context.Request.URL.Query()["partial_key"][0]
				partialKeyIndex := context.Request.URL.Query()["partial_key_index"][0]
				convertedIndex, err := strconv.Atoi(partialKeyIndex)
				if err != nil {
					log.Println(err)
					return
				}

				convertedPartialKey, err := hexutil.Decode("0x" + partialKey)
				if err != nil {
					log.Println(err)
				}

				var suite = edwards25519.NewBlakeSHA256Ed25519()
				scalarPartialKey := suite.Scalar().SetBytes(convertedPartialKey)
				log.Println(scalarPartialKey)
				// 데이터 저장
				repository.RestoreProviderResponseMap[repository.GetCurrentPublicKeyAddress()].ProviderResponseDatas[intProviderId] = model.RestoreProviderResponseData{PartialKey: &share.PriShare{I: convertedIndex, V: scalarPartialKey}}
			}
			context.HTML(http.StatusOK, "callback.tmpl", gin.H{})
		})
	}

}
