package services

import (
	"github.com/sslab-archive/key_custody_web/model"
	"strconv"
									)

var testProviderData = [...]model.Provider{
	model.Provider{ID: 1, Name: "provider1", Status: "Alive", EndpointUrl: "https://instagram.com"},
	model.Provider{ID: 2, Name: "provider2", Status: "Alive", EndpointUrl: "https://naver.com"},
	model.Provider{ID: 3, Name: "provider3", Status: "Alive", EndpointUrl: "https://facebook.com"},
	model.Provider{ID: 4, Name: "provider4", Status: "Alive", EndpointUrl: "https://facebook.com"},
	model.Provider{ID: 5, Name: "provider5", Status: "Alive", EndpointUrl: "https://facebook.com"},
}

func GetAllProviderList() ([]model.Provider, error) {

	return []model.Provider{}, nil
}

func GetMockProviderList(publicKeyAddress string) ([]model.Provider, error) {
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

func GetProviderListByUserKey() {

}

func SendTransactions(publickKeyAdddress string, privateKeyAddress string){
	//client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	//if err != nil {
	//	log.Println(err)
	//}
	//address := common.HexToAddress(publickKeyAdddress)
	//nonce, err := client.PendingNonceAt(context.Background(), address)
	//if err != nil {
	//	log.Println(err)
	//}
	//privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//value := big.NewInt(0) // in wei (0 eth)
	//gasPrice, err := client.SuggestGasPrice(context.Background())

}
