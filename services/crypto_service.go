package services

import (
	"crypto/rsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"crypto"
	"crypto/md5"
	"math/big"
	"github.com/sslab-archive/key_custody_web/model"
)

func VerifyProviderData(encryptedPayload string, encryptedPartialKey string, partialKey string, signedByPrivateKey string, providerPublicKey string) (model.UserPartialKeyDto, error) {
	// TODO EncryptedPartial Key를 상대의 pulic Key를 이용해서 partial Key를 복구

	//encryptedPayload := context.Request.URL.Query()["encrypted_payload"][0]
	//credentialType := context.Request.URL.Query()["credential_type"][0]
	//encryptedPartialKey := context.Request.URL.Query()["encrypted_partial_key"][0]
	//
	//payload := context.Request.URL.Query()["payload"][0]
	//credentialType := context.Request.URL.Query()["credential_type"][0]
	//PartialKey := context.Request.URL.Query()["partial_key"][0]
	//
	////[payload, credentialtype, pub(partialkey)]
	//signedByPrivateKey := context.Request.URL.Query()["signed_by_private_key"][0]
	//publicKey := context.Request.URL.Query()["provider_public_key"][0]


	// Public Key로 서명해서 준 데이터를 Verify
	publicKeyBytes, err := hexutil.Decode("0x" + providerPublicKey)
	if err != nil{
		log.Println(err)
	}

	signedByPrivateKeyBytes, err := hexutil.Decode("0x" + signedByPrivateKey)
	if err != nil{
		log.Println(err)
	}

	partialKeyBytes, err := hexutil.Decode("0x" + partialKey)
	if err != nil{
		log.Println(err)
	}
	hash := md5.New()           // 해시 인스턴스 생성
	hash.Write(partialKeyBytes) // 해시 인스턴스에 문자열 추가
	digest := hash.Sum(nil)

	var h2 crypto.Hash
	var bit *big.Int
	bit.SetBytes(publicKeyBytes)
	rsaPublicKey := rsa.PublicKey{N: bit, E: 65526}

	// 풀리면 일단 ㅇㅋ 하는 것으로
	err = rsa.VerifyPKCS1v15(&rsaPublicKey, h2, digest, signedByPrivateKeyBytes)
	if err != nil{
		log.Println(err)
		return model.UserPartialKeyDto{}, err
	}


	// encrypted_partial_key, encryptedPayload



	return model.UserPartialKeyDto{EncryptedPartialKey: encryptedPartialKey, EncryptedPayload: encryptedPayload}, nil
}
