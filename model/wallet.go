package model

type Wallet struct {
	PublicKeyAddress string `json:"publicKeyAddress"`
	PrivateKey       string `json:"privateKey"`
}

type KeyPairDao struct {
	PublicKeyAddress string `json:"publicKeyAddress"`
	IsRegistered     bool   `json:"isRegistered"`
}
