package did

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"
	"veric-backend/logic/blockchain/eth"
)

type VerifiableCredentialCredentialSubject interface {
	CanToByte
	CanToJson
	Name() string
}

type VerifiableCredential struct {
	credentialSubject VerifiableCredentialCredentialSubject

	Context           []string        `json:"@context" mapstructure:"@context"`
	ID                string          `json:"id"`
	Type              []string        `json:"type"`
	Issuer            string          `json:"issuer"`
	IssuanceDate      string          `json:"issuanceDate"`
	ExpirationDate    string          `json:"expirationDate"`
	Description       string          `json:"description"`
	CredentialSubject json.RawMessage `json:"credentialSubject"`
	Proof             VCProof         `json:"proof"`
}

func CreateVerifiableCredential(id string, issueDID *DIDDocument, credentialSubject VerifiableCredentialCredentialSubject) *VerifiableCredential {
	loc, _ := time.LoadLocation("UTC")

	return &VerifiableCredential{
		credentialSubject: credentialSubject,
		ID:                id,
		Context:           []string{ContextCredential, ContextSecp256k1},
		Type:              []string{TypeCredential, TypeVericDeposit, credentialSubject.Name()},
		Issuer:            issueDID.ID,
		IssuanceDate:      time.Now().In(loc).Format(time.RFC3339),
		ExpirationDate:    time.Now().In(loc).AddDate(10, 0, 0).Format(time.RFC3339),
		Description:       "Veric Deposit",
		CredentialSubject: credentialSubject.ToJson(),
		Proof:             CreateProof(issueDID.Authentication),
	}
}

func ParseVerifiableCredentialFromJsonStr(jsonDoc []byte) (*VerifiableCredential, error) {
	vc := &VerifiableCredential{}
	err := json.Unmarshal(jsonDoc, vc)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

func (vc *VerifiableCredential) init() {
	typ := vc.Type[len(vc.Type)-1]
	switch typ {
	case "VCSubjectDeposit":
		credential, err := ParseVCSubjectDepositFromVerifiableCredential(vc)
		if err != nil {
			return
		}

		vc.credentialSubject = credential
	}
}

func (vc *VerifiableCredential) ToJson() []byte {
	jsonDoc, err := json.Marshal(vc)
	if err != nil {
		return nil
	}

	return jsonDoc
}

func (vc *VerifiableCredential) ToByte() []byte {
	var convertedBytes bytes.Buffer

	context := cloneAndSort(vc.Context)
	typ := cloneAndSort(vc.Type)

	if vc.credentialSubject == nil {
		vc.init()
	}

	convertedBytes.WriteString(strings.Join(context, ","))
	convertedBytes.WriteString(vc.ID)
	convertedBytes.WriteString(strings.Join(typ, ","))
	convertedBytes.WriteString(vc.Issuer)
	convertedBytes.WriteString(vc.IssuanceDate)
	convertedBytes.WriteString(vc.ExpirationDate)
	convertedBytes.WriteString(vc.Description)
	if vc.credentialSubject != nil {
		convertedBytes.Write(vc.credentialSubject.ToByte())
	}
	convertedBytes.Write(vc.Proof.ToByte())

	return convertedBytes.Bytes()
}

func (vc *VerifiableCredential) Signature(issuerPrivKey *eth.PrivateKey) error {
	vc.Proof.JWSSignature = ""

	signatureData, err := issuerPrivKey.Sha256JWSSignature(vc.ToByte())
	if err != nil {
		return err
	}

	vc.Proof.JWSSignature = signatureData
	return nil
}

func (vc *VerifiableCredential) Verify(issuerPubKey *eth.PublicKey) (bool, error) {
	copiedVC := *vc
	copiedVC.Proof.JWSSignature = ""

	result, err := issuerPubKey.VerifySha256JWSSignature(copiedVC.ToByte(), vc.Proof.JWSSignature)
	if err != nil {
		return false, err
	}
	return result, nil
}
