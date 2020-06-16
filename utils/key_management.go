package utils

import (
	"os"
	"path/filepath"
		"log"
	"io/ioutil"
	"encoding/json"
	"github.com/sslab-archive/key_custody_web/model"
	)

// Load Key list from File
func LoadKeyList() ([]model.Wallet){

	var files []string
	var keyPairs []model.Wallet

	err := filepath.Walk("./keys", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir(){
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println(err)
		}
		keyPair := model.Wallet{}
		json.Unmarshal([]byte(data), &keyPair)
		keyPairs = append(keyPairs, keyPair)
	}

	return keyPairs
}

func StoreKey(address string, privateKey string){
	file, err := os.Create("./keys/" + address + ".json")
	wallet := model.Wallet{
		PublicKeyAddress: address,
		PrivateKey: privateKey,
	}
	jsonData, _ := json.MarshalIndent(wallet, "", " ")

	err = ioutil.WriteFile(file.Name(), jsonData, 0644)
	if err != nil{
		log.Println(err)
	}
}