package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sslab-archive/key_custody_web/controller"
	"github.com/sslab-archive/key_custody_web/repository"
	"github.com/sslab-archive/key_custody_web/services"
	"github.com/sslab-archive/key_custody_web/model"
	"fmt"
)

func main() {
	//repository.LoadKeyList()
	setMockData()
	startClientWebServer()
}
func startClientWebServer() {
	defaultRouter := gin.Default()
	defaultRouter.LoadHTMLGlob("templates/*")
	defaultRouter.Use(CORSMiddleware())
	controller.RegisterUserController(defaultRouter)
	defaultRouter.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	}
}

func setMockData() {
	repository.LoadKeyList()
	publicKeyAddress := "0x727985460763bA7BFfCe6BE84E5705068FB8DA43"
	var providerIds = []int{1, 2, 3, 4, 5}
	providerNumber := len(providerIds)
	keyPair := repository.GetKeyPairByPublicKeyAddress(publicKeyAddress)

	partialKeys := services.GeneratePartialKey(keyPair.PrivateKey, providerNumber)
	fmt.Println(partialKeys)

	repository.StorePartialKey(publicKeyAddress, providerIds, partialKeys)

	data1 := model.RestoreProviderResponseData{PartialKey: "8d2fe2e6805f368926844f6b6cba7529a60b2b05f624ab5c4b0cafa37efefc08", Index: 0}
	data2 := model.RestoreProviderResponseData{PartialKey: "23994158e914aae762b958c0057ab74544739cb2265b17b9d7a3bce386242f03", Index: 1}
	data3 := model.RestoreProviderResponseData{PartialKey: "fb23d8094fd37fa82aadc55a02451b5bcae764e3cf15501eeb183c9395d9b309", Index: 2}
	data4 := model.RestoreProviderResponseData{PartialKey: "f56a3370d381c20290f52794f585593d6b7a48ccb4fc0fc5b7ced70353e47507", Index: 3}
	data5 := model.RestoreProviderResponseData{PartialKey: "92d04a2aa5a2e24de0a0311907ae00f4bc9f4ca67acb50d013ef6881d4a03200", Index: 4}
	repository.RestoreProviderResponseMap[publicKeyAddress].ProviderResponseDatas[1] = data1
	repository.RestoreProviderResponseMap[publicKeyAddress].ProviderResponseDatas[2] = data2
	repository.RestoreProviderResponseMap[publicKeyAddress].ProviderResponseDatas[3] = data3
	repository.RestoreProviderResponseMap[publicKeyAddress].ProviderResponseDatas[4] = data4
	repository.RestoreProviderResponseMap[publicKeyAddress].ProviderResponseDatas[5] = data5

}
