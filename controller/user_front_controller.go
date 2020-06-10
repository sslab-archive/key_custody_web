package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"fmt"
	"crypto/ecdsa"
	"github.com/sslab-archive/key_custody_web/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

func RegisterUserController(router *gin.Engine) {


	keysRouter := router.Group("keys")
	{
		keysRouter.GET("/", func(context *gin.Context) {
			context.HTML(http.StatusOK, "key_generate.tmpl", gin.H{})
		})

		keysRouter.POST("/generate", func(context *gin.Context) {
			// 지갑 키 생성 해서 리턴해줌.
			privateKey, err := crypto.GenerateKey()
			if err != nil {
				log.Fatal(err)
			}
			privateKeyBytes := crypto.FromECDSA(privateKey)
			fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
			publicKey := privateKey.Public()

			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("error casting public key to ECDSA")
			}

			publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
			fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
			address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
			fmt.Println(address)


			file, err := os.Create("../keys/test_key.json")
			wallet := model.Wallet{
				PublicKeyAddress: address,
				PrivateKey: hexutil.Encode(privateKeyBytes)[2:],
			}
			jsonData, _ := json.MarshalIndent(wallet, "", " ")

			err = ioutil.WriteFile(file.Name(), jsonData, 0644)
			if err != nil{
				log.Println(err)
			}

			context.JSON(http.StatusOK, gin.H{
				"publicKeyAddress": address,
				"privateKey": hexutil.Encode(privateKeyBytes)[2:],
			})
		})

		keysRouter.POST("/store", func(context *gin.Context) {
			//생성된거 저장.
			// 리다이렉트 to request 화면.

		})

		keysRouter.GET("/request", func(context *gin.Context) {
			// provider list 보여주는 뷰 띄울 것
		})

		keysRouter.GET("/divide", func(context *gin.Context) {
			// 나뉘어진 키 값 보여주는 뷰 띄울 것.
		})

		keysRouter.GET("/provider1", func(context *gin.Context) {
			// 1번 프로바이더에게 요청 화면 - 이메일 인증?
		})
		keysRouter.GET("/provider2", func(context *gin.Context) {
			// 2번 프로바이더에게 요청 화면 - 휴대폰 인증
		})
		keysRouter.GET("/provider3", func(context *gin.Context) {
			// 3번 프로바이더에게 요청 화면 - 질문 1
		})
		keysRouter.GET("/provider4", func(context *gin.Context) {
			// 4번 프로바이더에게 요청 화면 - 질문 2
		})
		keysRouter.GET("/provider5", func(context *gin.Context) {
			// 5번 프로바이더에게 요청 화면 - 질문 3
		})
	}

}
