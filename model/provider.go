package model

type Provider struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	EndpointUrl string `json:"endPointUrl"`
	// TODO 더 있으면 추가
}
