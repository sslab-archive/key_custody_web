package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-archive/key_custody_web/utils"
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
			keyPairs := utils.LoadKeyList()
			for _, keyPair := range keyPairs {
				if services.CheckKeyIsRegistered(keyPair.PublicKeyAddress) {
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
			utils.StoreKey(keyPair.PublicKeyAddress, keyPair.PrivateKey)

			context.JSON(http.StatusOK, gin.H{
				"publicKeyAddress": keyPair.PublicKeyAddress,
				"privateKey":       keyPair.PrivateKey,
			})
		})

		keysRouter.GET("/divide", func(context *gin.Context) {
			// 나뉘어진 키 값 보여주는 뷰 띄울 것.
		})

	}

	providerViewRouter := router.Group("provider")
	{
		providerViewRouter.GET("list", func(context *gin.Context) {
			// Provider list 보여주는 뷰 띄울 것
			publicKeyAddress := context.Request.URL.Query()["publicKeyAddress"][0]
			// 프로바이더 리스트 뷰로 띄워주기
			// TODO GetProvider List 구현 후 바꿔칠 것.
			providerList, err := services.GetMockProviderList(publicKeyAddress)
			if err != nil {
				log.Println(err)
			}
			context.HTML(http.StatusOK, "provider_list.tmpl", gin.H{
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

			context.HTML(http.StatusOK, "provider_registry.tmpl", gin.H{
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
			log.Println(publicKeyAddress)
			// TODO PublicKeyAddress로 그거 하기.
			providers := context.Request.URL.Query()["providers"][0]
			log.Println(providers)
			ids := strings.Split(providers, ",")
			var intIds []int
			for _, strId := range ids {
				intId, err := strconv.Atoi(strId)
				if err != nil {
					panic(err)
				}
				intIds = append(intIds, intId)
			}

			providerList, err := services.GetProviderListByIds(intIds)
			if err != nil{
				log.Println(err)
			}

			context.HTML(http.StatusOK, "provider_auth_index.tmpl", gin.H{
				"publicKeyAddress": publicKeyAddress,
				"providerList": providerList,
			})
		})
	}

}
