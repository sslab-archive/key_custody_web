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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EndpointUrl string `json:"endPointUrl"`
	PartialKey string `json:"partialKey"`
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
	Index      int    `json:"index"`
	PartialKey string `json:"partialKey"`
}
