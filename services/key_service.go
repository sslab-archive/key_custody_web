package services

import (
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"github.com/sslab-archive/key_custody_web/model"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/share"
	"log"
	)

var suite = edwards25519.NewBlakeSHA256Ed25519()

func GenerateKeyPair() (model.KeyPair, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return model.KeyPair{}, err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	if err != nil {
		return model.KeyPair{}, err
	}

	wallet := model.KeyPair{
		PublicKeyAddress: address,
		PublicKey:        hexutil.Encode(publicKeyBytes)[4:],
		PrivateKey:       hexutil.Encode(privateKeyBytes)[2:],
	}
	return wallet, nil
}

func GeneratePartialKey(privateKey string, providerNum int) []*share.PriShare {

	// Byte Conversion
	privateKeyBytes, err := hexutil.Decode("0x" + privateKey)
	if err != nil{
		log.Println(err)
	}
	threshold := providerNum - 1
	// 다항식 세우기 랜덤으로...
	secret := suite.Scalar().SetBytes(privateKeyBytes)

	priPoly := share.NewPriPoly(suite, threshold, secret, suite.RandomStream())

	// Create secret set of shares
	priShares := priPoly.Shares(providerNum)

	// Partial keys
	return priShares
}

func RestorePartialKey(partialKeys []*share.PriShare, providerNum int, threshold int) string {

	recoveredSecret, err := share.RecoverSecret(suite, partialKeys, threshold, providerNum)
	if err != nil {
		log.Println(err)
	}

	return recoveredSecret.String()
}

func GetRestorePartialKeys(datas map[int]model.RestoreProviderResponseData) []*share.PriShare{

	var priShares = []*share.PriShare{}

	for _, value := range datas {
		priShares = append(priShares, value.PartialKey)
	}
	return priShares
}

//func EncryptedData(partialKeys []model.PartialKeyProviderMappingEntity, responses []model.ProviderResponseMappingEntity){
//	// 데이터를 전부 올림...
//}
