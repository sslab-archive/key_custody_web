package model

import "go.dedis.ch/kyber/v3/share"

type Provider struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	EndpointUrl string `json:"endPointUrl"`
	PublicKey   string `json:"publicKey"`
	// TODO 더 있으면 추가
}

type ProviderAuthDTO struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	EndpointUrl     string `json:"endPointUrl"`
	PartialKey      string `json:"partialKey"`
	PartialKeyIndex int    `json:"partialKeyIndex"`
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
	ProviderId          int    `json:"providerId"`
	Payload             string `json:"payload"`
	CredentialType      string `json:"credentialType"`
	EncryptedPartialKey string `json:"encryptedPartialKey"`
	EncryptedPayload    string `json:"encryptedPayload"`
	SignedByPrivateKey  string `json:"signedByPrivateKey"`
}

type RestoreProviderResponseMappingEntity struct {
	ProviderNumber        int `json:"providerNumber"`
	Threshold             int `json:"threshold"`
	ProviderResponseDatas map[int]RestoreProviderResponseData
}

type RestoreProviderResponseData struct {
	PartialKey *share.PriShare `json:"partialKey"`
}
