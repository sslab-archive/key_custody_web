package services

import (
	"github.com/sslab-archive/key_custody_web/model"
	"strconv"
	"math/rand"
)

var testProviderData = [...]model.Provider{
	model.Provider{ID: 1, Name: "provider1", Status: "Alive", EndpointUrl: "https://instagram.com"},
	model.Provider{ID: 2, Name: "provider2", Status: "Alive", EndpointUrl: "https://naver.com"},
	model.Provider{ID: 3, Name: "provider3", Status: "Alive", EndpointUrl: "https://facebook.com"},
	model.Provider{ID: 4, Name: "provider4", Status: "Alive", EndpointUrl: "https://facebook.com"},
	model.Provider{ID: 5, Name: "provider5", Status: "Alive", EndpointUrl: "https://facebook.com"},
}

func CheckKeyIsRegistered(publicKeyAddress string) bool{
	randomNumber := rand.Intn(100)

	if randomNumber > 50 {
		return true
	}

	return false
}

func GetAllProviderList() ([]model.Provider, error){

	return []model.Provider{}, nil
}


func GetMockProviderList(publicKeyAddress string) ([]model.Provider, error){
	var providers = []model.Provider{}

	for i := 1; i <= 3; i++ {
		providers = append(providers, model.Provider{ID: i, Name: "provider" + strconv.Itoa(i), Status: "Alive", EndpointUrl: "https://instagram.com"})
	}
	for i := 4; i <= 5; i++ {
		providers = append(providers, model.Provider{ID: i, Name: "provider" + strconv.Itoa(i), Status: "Dead", EndpointUrl: "https://naver.com"})
	}

	return providers, nil
}

func GetProviderListByIds(ids []int) ([]model.Provider, error) {
	return testProviderData[:], nil
}


func GetProviderListByUserKey(){

}