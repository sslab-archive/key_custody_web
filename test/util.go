package main

import (
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/share"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	)

func main() {
	privateKey, err := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	suite := edwards25519.NewBlakeSHA256Ed25519()

	providerNum := 5
	threshold := providerNum - 1
	// 다항식 세우기 랜덤으로...
	fmt.Println("프라이빗 키 바이트 길이: ", len(privateKeyBytes))
	fmt.Println("프라이빗 키 바이트 ", privateKeyBytes)

	secret := suite.Scalar().SetBytes(privateKeyBytes)
	fmt.Println(secret)


	priPoly := share.NewPriPoly(suite, threshold, secret, suite.RandomStream())

	// Create secret set of shares
	// 비밀키 값을 나눴음..

	priShares := priPoly.Shares(providerNum)

	recoveredSecret, err := share.RecoverSecret(suite, priShares, threshold, providerNum)
	if err != nil {
		fmt.Println(err)
	}
	if !recoveredSecret.Equal(priPoly.Secret()) {
		fmt.Println("recovered secret does not match initial value")
	}

	fmt.Println("리코버된 프라이빗 키 바이트 : ", []byte(recoveredSecret.String()))

}

