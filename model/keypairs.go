package model

import "go.dedis.ch/kyber/v3/share"

type KeyPair struct {
	PublicKeyAddress string `json:"publicKeyAddress"`
	PublicKey        string `json:"publicKey"`
	PrivateKey       string `json:"privateKey"`
}

type KeyPairDao struct {
	PublicKeyAddress string `json:"publicKeyAddress"`
	IsRegistered     bool   `json:"isRegistered"`
}

type UserPartialKeyManagementEntity struct {
	ProviderNumber             int `json:"providerNumber"`
	Threshold                  int `json:"threshold"`
	PartialKeyProviderEntities map[int]*share.PriShare
}

type ProviderResponseMappingEntity struct {
	ProviderNumber        int `json:"providerNumber"`
	Threshold             int `json:"threshold"`
	ProviderResponseDatas map[int]ProviderResponseData
}

type ProviderResponseData struct {
	ProviderId       int    `json:"providerId"`
	Payload          string `json:"payload"`
	CredentialType   string `json:"credentialType"`
	SignedPartialKey string `json:"signedPartialKey"`
	SignedPayload    string `json:"signedPayload"`
	SignedAllData    string `json:"signedAllData"`
}

type RestoreProviderResponseMappingEntity struct {
	ProviderNumber        int `json:"providerNumber"`
	Threshold             int `json:"threshold"`
	ProviderResponseDatas map[int]RestoreProviderResponseData
}

type RestoreProviderResponseData struct {
	PartialKey string `json:"partialKey"`
}
