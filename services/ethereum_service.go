package services

import (
	"github.com/sslab-archive/key_custody_web/model"
											)

var testProviderData = [...]model.Provider{
	model.Provider{ID: 1, Name: "Haesung Company", Status: "Alive", EndpointUrl: "http://141.223.121.111:8888/authentication"},
	model.Provider{ID: 2, Name: "Hoonki Company", Status: "Alive", EndpointUrl: "https://naver.com"},
	model.Provider{ID: 3, Name: "Jeyoung Company", Status: "Alive", EndpointUrl: "https://facebook.com"},
	model.Provider{ID: 4, Name: "Naver Company", Status: "Alive", EndpointUrl: "https://facebook.com"},
	model.Provider{ID: 5, Name: "Woori Company", Status: "Alive", EndpointUrl: "https://facebook.com"},
}

func GetAllProviderList() ([]model.Provider, error) {

	return testProviderData[:], nil
}

func GetProviderListByIds(ids []int) ([]model.Provider, error) {
	var providers = []model.Provider{}
	for i := 0; i < len(testProviderData); i++ {
		if contains(ids, testProviderData[i].ID) {
			providers = append(providers, testProviderData[i])
		}

	}
	return providers[:], nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


// TODO 프로바이더 리스트 그니까 등록된 퍼블릭키 주소에 대해서 ..? 아니면 msg.sender 즉 퍼블릭 주소에 대해서..
func GetProviderListByUserKey(publicKeyAddress string) {

}

// TODO PartialKey 등록하는 인터페이스.
func RegisterPartialKey(dto model.UserPartialKeyDto) (error) {

	//client, err := ethclient.Dial(config.EthereumConfig["ethereuemEndPoint"])
	//if err != nil {
	//	return err
	//}
	//
	//// 쓰려는 기본 private Key
	//basePrivateKey := config.EthereumConfig["privateKeyAddress"]
	//privateKey, err := crypto.HexToECDSA(basePrivateKey)
	//if err != nil {
	//	return err
	//}
	//
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//	return err
	//}
	//
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	return err
	//}
	//
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	return err
	//}
	//
	//auth := bind.NewKeyedTransactor(privateKey)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)     // in wei
	//auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = gasPrice
	//
	//address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	//// TODO 여기에 RegisterPartialKey.
	//instance, err := store.NewStore(address, client)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//key := [32]byte{}
	//value := [32]byte{}
	//copy(key[:], []byte("foo"))
	//copy(value[:], []byte("bar"))
	//
	//// TODO  RegisterPartialKey.
	//tx, err := instance.SetItem(auth, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	//
	//result, err := instance.Items(nil, key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(result[:])) // "bar"
	return nil
}