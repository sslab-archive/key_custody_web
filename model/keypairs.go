package model

type KeyPair struct {
	PublicKeyAddress string `json:"publicKeyAddress"`
	PublicKey        string `json:"publicKey"`
	PrivateKey       string `json:"privateKey"`
}

type KeyPairDao struct {
	PublicKeyAddress string `json:"publicKeyAddress"`
	PublicKey        string `json:"publicKey"`
	PrivateKey       string `json:"privateKey"`
	IsRegistered     bool   `json:"isRegistered"`
}


