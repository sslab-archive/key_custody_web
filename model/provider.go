package model

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
