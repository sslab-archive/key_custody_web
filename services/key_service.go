package services

import (
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"github.com/sslab-archive/key_custody_web/model"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"go.dedis.ch/kyber/v3/group/edwards25519"
	"fmt"
	"go.dedis.ch/kyber/v3/share"
	"log"
	)

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

	// 암호 알고리즘인듯
	fmt.Println("시작이다 Partial Key 생성~~ ", privateKey)
	suite := edwards25519.NewBlakeSHA256Ed25519()
	threshold := providerNum - 1
	// 다항식 세우기 랜덤으로...
	secret := suite.Scalar().SetBytes([]byte(privateKey))

	priPoly := share.NewPriPoly(suite, threshold, secret, suite.RandomStream())

	// Create secret set of shares
	priShares := priPoly.Shares(providerNum)

	// Partial keys
	return priShares
}

func RestorePartialKey(partialKeys []*share.PriShare) string {
 	return "Hello"
}

//func EncryptedData(partialKeys []model.PartialKeyProviderMappingEntity, responses []model.ProviderResponseMappingEntity){
//	// 데이터를 전부 올림...
//}
