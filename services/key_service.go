package services

import (
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
		"github.com/sslab-archive/key_custody_web/model"
	"github.com/ethereum/go-ethereum/common/hexutil"

		"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/share/pvss"
			"fmt"
	"log"
)

func GenerateKeyPair() (model.Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return model.Wallet{}, err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return model.Wallet{}, err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	if err != nil{
		return model.Wallet{}, err
	}

	wallet := model.Wallet{
		PublicKeyAddress: address,
		PrivateKey: hexutil.Encode(privateKeyBytes)[2:],
	}
	return wallet, nil
}

func GeneratePartialKey(privateKey string, divideNumber int){
	// EncShares creates a list of encrypted publicly verifiable PVSS shares for
	// the given secret and the list of public keys X using the sharing threshold
	// t and the base point H. The function returns the list of shares and the
	// public commitment polynomial.

	//suite Suite, H kyber.Point, X []kyber.Point, secret kyber.Scalar, t int
	suite := edwards25519.NewBlakeSHA256Ed25519()
	H := suite.Point().Pick(suite.XOF([]byte("H")))
	n := 512
	t := 2*n/3 + 1
	X := make([]kyber.Point, n)  // trustee public keys
	secret := suite.Scalar().Pick(suite.RandomStream())


	encShares, pubPoly, err := pvss.EncShares(suite, H, X, secret, t)
	if err !=nil {
		log.Println(err)
	}
	fmt.Println(encShares)
	fmt.Println(pubPoly)


}
