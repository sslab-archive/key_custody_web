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
)

func RegisterUserController(router *gin.Engine) {

	keysRouter := router.Group("keys")
	{
		keysRouter.GET("/index", func(context *gin.Context) {
			var keyDaos = []model.KeyPairDao{}
			keyPairs := repository.KeyPairs
			for _, keyPair := range keyPairs {
				if repository.CheckKeyIsRegistered(keyPair.PublicKeyAddress) {
					keyDaos = append(keyDaos, model.KeyPairDao{keyPair.PublicKeyAddress, true})
				} else {
					keyDaos = append(keyDaos, model.KeyPairDao{keyPair.PublicKeyAddress, false})
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
			if !result{
				log.Println("NO public Key Param")
			}
			//repository.ProviderResponseMap
		})

		keysRouter.GET("/blockchain/list", func(context *gin.Context) {

			publicKeyAddress := context.Request.URL.Query()["publicKeyAddress"][0]
			keyPair := repository.GetKeyPairByPublicKeyAddress(publicKeyAddress)

			context.HTML(http.StatusOK, "key_blockchain_list.tmpl",gin.H{
				"publicKeyAddress": keyPair.PublicKeyAddress,
				"publicKey": keyPair.PublicKey,
				"privateKey": keyPair.PrivateKey,
			})

		})

		keysRouter.GET("/restore", func(context *gin.Context) {

			var providerIds []int
			var providerData []model.ProviderResponseData
			for key, val := range repository.ProviderResponseMap {
				fmt.Println(key, val)
				for key, val := range val.ProviderResponseDatas{
					providerIds = append(providerIds, key)
					providerData = append(providerData, val)
				}
			}
			providers, err := services.GetProviderListByIds(providerIds)
			if err != nil{
				log.Println(err)
			}

			// 여기서 다보여주는 것은 어떨까?
			context.HTML(http.StatusOK, "key_restore.tmpl",gin.H{
				"providers": providers,
			})
		})

	}

	providerViewRouter := router.Group("providers")
	{
		providerViewRouter.GET("list", func(context *gin.Context) {
			// Provider list 보여주는 뷰 띄울 것
			publicKeyAddress := context.Request.URL.Query()["publicKeyAddress"][0]
			// 프로바이더 리스트 뷰로 띄워주기
			providerList, err := services.GetMockProviderList(publicKeyAddress)
			if err != nil {
				log.Println(err)
			}
			context.HTML(http.StatusOK, "providers_list.tmpl", gin.H{
				"providerList": providerList,
			})
		})

		providerViewRouter.GET("registry", func(context *gin.Context) {
			// 프로바이더 리스트 등록 하기...
			publicKeyAddress := context.Request.URL.Query()["publicKeyAddress"][0]
			log.Println(publicKeyAddress)

			providerList, err := services.GetMockProviderList(publicKeyAddress)
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
				}else {
					for _, val := range value {
						id, _ := strconv.Atoi(val)
						providerIds = append(providerIds, id)
					}
				}
			}
			queryString := "publicKeyAddress=" + publicKeyAddress + "&providers=" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(providerIds)), ","), "[]")

			context.Redirect(http.StatusMovedPermanently, "auth/index?" + queryString)
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

			for i, provider := range providerList{
				tempDto := model.ProviderAuthDTO{ ID: provider.ID, Name: provider.Name, EndpointUrl: provider.EndpointUrl, PartialKey: partialKeys[i].String(), }
				//partialKeys[i].I
				//partialKeys[i].V
				providerAuthDtos = append(providerAuthDtos, tempDto)
			}
			//log.Println(providerAuthDtos)
			//log.Println(len(providerAuthDtos))

			context.HTML(http.StatusOK, "providers_auth_index.tmpl", gin.H{
				"publicKeyAddress": publicKeyAddress,
				"providerDtos": providerAuthDtos,
				"redirectUrl": "http://localhost:8080/callback",
			})
		})

		providerViewRouter.POST("onReady", func(context *gin.Context) {
			// Provider들로부터 응답을 다 받아야함. Response를 관리해야함.
			publicKeyAddress, _ := context.GetPostForm("publicKeyAddress")
			if repository.CheckProviderResponseIsReady(publicKeyAddress){
				context.JSON(http.StatusOK, gin.H{
					"result": true,
				})
			}else{
				context.JSON(http.StatusOK, gin.H{
					// false
					"result": true,
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
				if repository.CheckProviderResponse(publicKeyAddress, intProviderId){
					data := repository.ProviderResponseMap[publicKeyAddress].ProviderResponseDatas[intProviderId]
					context.JSON(http.StatusOK, gin.H{
						"result": true,
						"data": data,
					})
				}else{
					context.JSON(http.StatusOK, gin.H{
						"result": false,
					})
				}
			}else{
				// TODO restore 과정에서 response를 받았는지 체크하는 과정.
				context.JSON(http.StatusOK, gin.H{
					// false
					"result": true,
				})

			}


		})

		providerViewRouter.POST("restorePrivateKey", func(context *gin.Context) {
			//repository.
			context.JSON(http.StatusOK, gin.H{
				// false
				"result": true,
				"privateKey": "5469f91d0542d0486e23d2460cba0ebd2ed38c3477d168543951476ba201fab1",
			})

		})

	}

	callbackRouter := router.Group("callback")
	{
		callbackRouter.GET("", func(context *gin.Context) {
			// Refresh Confirm 추가.

			//TODO purpose[encrypt, decrypt]
			//TODO case encrypt || payload [ex) hea9549] , credential_type [ex) email, phone], signed_partial_key, signed_payload, signed_by_privatekey(credential_type, signed_partial_key, signed_payload)

			//TODO case decrypt || partialKey 넘겨줌
			//TODO partial Key, 블록체인에 등록된 내 provider Public Key로 암호화 했을 떄 저게 나오는지.
		})
	}

}
