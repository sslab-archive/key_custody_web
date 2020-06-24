package main

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"fmt"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/share"
	"log"
)

func main() {
	privateKeyString := "92d04a2aa5a2e24de0a0311907ae00f4bc9f4ca67acb50d013ef6881d4a03200"
	privateKeyBytes, err := hexutil.Decode("0x" + privateKeyString)
	if err != nil {
		log.Println(err)
	}

	suite := edwards25519.NewBlakeSHA256Ed25519()

	providerNum := 5
	threshold := providerNum - 1
	// 다항식 세우기 랜덤으로...
	secret := suite.Scalar().SetBytes(privateKeyBytes)

	priPoly := share.NewPriPoly(suite, threshold, secret, suite.RandomStream())
	// Create secret set of shares
	// 비밀키 값을 나눴음..

	priShares := priPoly.Shares(providerNum)
	fmt.Println(priShares[0].V)
	partialString := priShares[0].V.String()
	fmt.Println("partial String", partialString)
	partialBytes, _ := hexutil.Decode("0x" + partialString)
	fmt.Println("Decoded Partial Byte", partialBytes)
	s := share.PriShare{I: 0, V: suite.Scalar().SetBytes(partialBytes)}
	fmt.Println("V", s.V)

	recoveredSecret, err := share.RecoverSecret(suite, priShares, threshold, providerNum)
	if err != nil {
		fmt.Println(err)
	}
	if !recoveredSecret.Equal(priPoly.Secret()) {
		fmt.Println("recovered secret does not match initial value")
	}

	fmt.Println(recoveredSecret.String())

}
