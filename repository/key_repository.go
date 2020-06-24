package repository

import (
	"github.com/sslab-archive/key_custody_web/model"
	"path/filepath"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"go.dedis.ch/kyber/v3/share"
	"path"
	"strings"
)


var KeyPairs = []model.KeyPair{}

var UserPartialKeyMap = map[string]model.UserPartialKeyManagementEntity{}
var ProviderResponseMap = map[string]model.ProviderResponseMappingEntity{}
var RestoreProviderResponseMap = map[string]model.RestoreProviderResponseMappingEntity{}


func CheckKeyIsRegistered(publicKeyAddress string) bool {
	_, result := UserPartialKeyMap[publicKeyAddress]
	return result
}

// Load Key list from File
func LoadKeyList() {

	var files []string
	var keyPairs []model.KeyPair

	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/sslab-archive/key_custody_web")

	err := filepath.Walk(ap + "/keys", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.Contains(info.Name(), ".json") {
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
		keyPair := model.KeyPair{}
		json.Unmarshal([]byte(data), &keyPair)
		keyPairs = append(keyPairs, keyPair)
	}
	KeyPairs = keyPairs
}

func GetCurrentPublicKeyAddress() string{
	for k := range UserPartialKeyMap {
		return k
	}
	return ""
}

func GetKeyPairByPublicKeyAddress(publicKeyAddress string) model.KeyPair {
	for _, val := range KeyPairs {
		if val.PublicKeyAddress == publicKeyAddress {
			return val
		}
	}
	return model.KeyPair{}
}

func StoreKey(publicKeyAddress string, publicKey,privateKey string) {
	file, err := os.Create("./keys/" + publicKeyAddress + ".json")
	wallet := model.KeyPair{
		PublicKeyAddress: publicKeyAddress,
		PublicKey: publicKey,
		PrivateKey:       privateKey,
	}
	jsonData, _ := json.MarshalIndent(wallet, "", " ")

	err = ioutil.WriteFile(file.Name(), jsonData, 0644)
	if err != nil {
		log.Println(err)
	}
	KeyPairs = append(KeyPairs, wallet)
}

// Initial이라 상관 없음.
func StorePartialKey(address string, providerIds []int, partialKeys []*share.PriShare) {
	// 공개키에 대해 프로바이더 리스트랑 다 저장...
	var tempEntity = model.UserPartialKeyManagementEntity{}

	tempEntity.ProviderNumber = len(providerIds)
	tempEntity.Threshold = len(providerIds) - 1
	tempEntity.PartialKeyProviderEntities = make(map[int]*share.PriShare)

	for i := 0; i < len(providerIds); i++{
		tempEntity.PartialKeyProviderEntities[providerIds[i]] = partialKeys[i]
	}

	UserPartialKeyMap[address] = tempEntity
	// 데이터 셋팅
	setupProviderResponseInformation(address, providerIds)
	setupRestoreProvderResponseInformation(address, providerIds)
}
func setupProviderResponseInformation(address string, providerIds []int){
	var tempEntity = model.ProviderResponseMappingEntity{}
	tempEntity.ProviderNumber = len(providerIds)
	tempEntity.Threshold = len(providerIds) - 1
	tempEntity.ProviderResponseDatas = make(map[int]model.ProviderResponseData)
	ProviderResponseMap[address] = tempEntity
}

func setupRestoreProvderResponseInformation(address string, providerIds []int){
	var tempEntity = model.RestoreProviderResponseMappingEntity{}
	tempEntity.ProviderNumber = len(providerIds)
	tempEntity.Threshold = len(providerIds) - 1
	tempEntity.ProviderResponseDatas = make(map[int]model.RestoreProviderResponseData)
	RestoreProviderResponseMap[address] = tempEntity
}


func StoreProviderResponse(address string, providerId int, providerResponseData model.ProviderResponseData){
	ProviderResponseMap[address].ProviderResponseDatas[providerId] = providerResponseData
}

func CheckProviderResponseIsReady(address string) bool {
	_, result := ProviderResponseMap[address]
	return result
}

func CheckProviderResponse(address string, providerId int) bool {
	providerResponse, result := ProviderResponseMap[address]
	if !result{
		return false
	}
	_, isDataExist := providerResponse.ProviderResponseDatas[providerId]
	return isDataExist
}

func CheckRestoreProviderResponse(address string, providerId int) bool{
	providerResponse, result := RestoreProviderResponseMap[address]
	if !result{
		return false
	}
	_, isDataExist := providerResponse.ProviderResponseDatas[providerId]
	return isDataExist
}
