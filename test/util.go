package main

import (
	"github.com/sslab-archive/key_custody_web/services"
	"log"
	"github.com/sslab-archive/key_custody_web/utils"
)

func main(){

	keypair, err := services.GenerateKeyPair()
	if err != nil{
		log.Println(err)
	}
	log.Println(keypair.PrivateKey)

	utils.StoreKey(keypair.PublicKeyAddress, keypair.PublicKeyAddress)
	
}
