package did

import (
	"bytes"
	"encoding/json"
	"errors"
)

type VCSubjectDeposit struct {
	Chain             string `json:"chain"`
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	MerchantAmount    string `json:"merchant_amount"`
	PlatformFeeAmount string `json:"platform_fee_amount"`
	PoolFeeAmount     string `json:"pool_fee_amount"`
}

func ParseVCSubjectDepositFromVerifiableCredential(vc *VerifiableCredential) (*VCSubjectDeposit, error) {
	if vc.Type[len(vc.Type)-1] != "VCSubjectDeposit" {
		return nil, errors.New("not VCSubjectDeposit object")
	}
	obj := &VCSubjectDeposit{}
	err := json.Unmarshal(vc.CredentialSubject, obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (v *VCSubjectDeposit) Name() string {
	return "VCSubjectDeposit"
}

func (v *VCSubjectDeposit) ToJson() []byte {
	jsonDoc, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	return jsonDoc
}

func (v *VCSubjectDeposit) ToByte() []byte {
	var convertedBytes bytes.Buffer

	convertedBytes.WriteString(v.Chain)
	convertedBytes.WriteString(v.Currency)
	convertedBytes.WriteString(v.Amount)
	convertedBytes.WriteString(v.PlatformFeeAmount)
	convertedBytes.WriteString(v.PoolFeeAmount)
	convertedBytes.WriteString(v.MerchantAmount)

	return convertedBytes.Bytes()
}
