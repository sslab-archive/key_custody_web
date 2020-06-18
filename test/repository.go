package main

import (
	"github.com/sslab-archive/key_custody_web/repository"
	"github.com/sslab-archive/key_custody_web/services"
	"fmt"
)

//func main(){
//	repository.LoadKeyList()
//	var providerIds = []int{1, 2, 3, 4, 5}
//	publicKeyAddress := "0x0f84B2DEc6f292C5CE8c4cbf2A91D44C9b71154e"
//	providerNumber := len(providerIds)
//	fmt.Println(providerNumber)
//	privateKey := repository.GetPrivateKeyByPublicKey(publicKeyAddress)
//	fmt.Println(privateKey)
//
//	partialKeys := services.GeneratePartialKey(privateKey, providerNumber)
//
//	repository.StorePartialKey(publicKeyAddress, providerIds, partialKeys)
//	fmt.Println(repository.PartialKeyMap)
//}

func main(){
	repository.LoadKeyList()
	var providerIds = []int{1, 2, 3, 4, 5}
	publicKeyAddress := "0x0f84B2DEc6f292C5CE8c4cbf2A91D44C9b71154e"
	providerNumber := len(providerIds)
	fmt.Println(providerNumber)
	privateKey := repository.GetPrivateKeyByPublicKey(publicKeyAddress)
	fmt.Println(privateKey)

	partialKeys := services.GeneratePartialKey(privateKey, providerNumber)

	repository.StorePartialKey(publicKeyAddress, providerIds, partialKeys)
	fmt.Println(repository.UserPartialKeyMap)
	//fmt.Println(repository.PartialKeyMap)

	repository.StoreProviderResponse(publicKeyAddress, 1, "sibal")
	fmt.Println(repository.ProviderResponseMap)
	repository.StoreProviderResponse(publicKeyAddress, 2, "hoonki")
	fmt.Println(repository.ProviderResponseMap)

}
