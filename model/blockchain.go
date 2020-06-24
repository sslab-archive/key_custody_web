package model

type UserPartialKeyDto struct {
	EncryptedPartialKey string `json:"encryptedPartialKey"`
	EncryptedPayload string `json:"encryptedPayload"`
}
