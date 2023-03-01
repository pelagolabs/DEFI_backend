package did

import (
	"encoding/json"
	"github.com/multiformats/go-multibase"
	"time"
	"veric-backend/logic/blockchain/eth"
)

type VerificationMethod struct {
	ID           string `json:"id"`
	MethodType   string `json:"type"`
	Controller   string `json:"controller"`
	MultibaseKey string `json:"publicKeyMultibase"`
}

type Service struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

type DIDDocument struct {
	Context            []string             `json:"@context" mapstructure:"@context"`
	ID                 string               `json:"id"`
	Created            string               `json:"created"`
	Updated            string               `json:"updated"`
	Version            int                  `json:"version"`
	VerificationMethod []VerificationMethod `json:"verificationMethod"`
	Authentication     string               `json:"authentication"`
	Service            []Service            `json:"service"`
	Address            string               `json:"address"`
}

func CreateUserDID(userPbKey string, issuePriKey *eth.PrivateKey) (*DIDDocument, error) {
	documentId := DIDPrefix + userPbKey
	currentTime := time.Now().Format(time.RFC3339)
	pubData, err := multibase.Encode(multibase.Base58BTC, issuePriKey.PublicKey().EllipticMarshal())
	if err != nil {
		return nil, err
	}

	return &DIDDocument{
		ID:             documentId,
		Context:        []string{ContextSecp256k1, ContextDID},
		Created:        currentTime,
		Updated:        currentTime,
		Version:        1,
		Authentication: documentId + "#verification",
		Address:        userPbKey,
		VerificationMethod: []VerificationMethod{
			{
				ID:           documentId + "#verification",
				Controller:   documentId,
				MethodType:   Secp256k1Key,
				MultibaseKey: pubData,
			},
		},
	}, nil
}

func ParseUserDIDFromJsonStr(jsonDoc []byte) (*DIDDocument, error) {
	document := &DIDDocument{}
	err := json.Unmarshal(jsonDoc, document)
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (d *DIDDocument) ToJson() []byte {
	jsonDoc, err := json.Marshal(d)
	if err != nil {
		return nil
	}

	return jsonDoc
}

func (d *DIDDocument) RetrieveVerificationMethod(vmID string) (VerificationMethod, error) {
	for _, vm := range d.VerificationMethod {
		if vm.ID == vmID {
			return vm, nil
		}
	}
	return VerificationMethod{}, ErrMissingVM
}
