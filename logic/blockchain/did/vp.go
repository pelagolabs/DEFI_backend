package did

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"
	"veric-backend/logic/blockchain/eth"
)

type VerifiableCredentialArr []VerifiableCredential

func (v VerifiableCredentialArr) Len() int {
	return len(v)
}

func (v VerifiableCredentialArr) Less(i, j int) bool {
	return v[i].ID < v[j].ID
}

func (v VerifiableCredentialArr) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v VerifiableCredentialArr) ToByte() []byte {
	var convertedBytes bytes.Buffer
	for _, vc := range v {
		convertedBytes.Write(vc.ToByte())
	}
	return convertedBytes.Bytes()
}

func (v VerifiableCredentialArr) Verify(pubKey *eth.PublicKey) (bool, error) {
	for _, vc := range v {
		verify, err := vc.Verify(pubKey)
		if !verify || err != nil {
			return verify, err
		}
	}

	return true, nil
}

type VerifiablePresentation struct {
	Context              []string                `json:"@context" mapstructure:"@context"`
	Type                 []string                `json:"type"`
	VerifiableCredential VerifiableCredentialArr `json:"verifiableCredential"`
	Holder               string                  `json:"holder"`
	Proof                VPProof                 `json:"proof"`
}

func ParseVerifiablePresentationFromJsonStr(jsonDoc []byte) (*VerifiablePresentation, error) {
	vp := &VerifiablePresentation{}
	err := json.Unmarshal(jsonDoc, vp)
	if err != nil {
		return nil, err
	}
	return vp, nil
}

func (vp *VerifiablePresentation) ToJson() []byte {
	jsonDoc, err := json.Marshal(vp)
	if err != nil {
		return nil
	}

	return jsonDoc
}

func (vp *VerifiablePresentation) ToByte() []byte {
	var convertedBytes bytes.Buffer

	for _, credential := range vp.VerifiableCredential {
		credential.init()
	}

	context := cloneAndSort(vp.Context)
	typ := cloneAndSort(vp.Type)

	vca := make(VerifiableCredentialArr, len(vp.VerifiableCredential))
	copy(vca, vp.VerifiableCredential)
	sort.Sort(vca)

	convertedBytes.WriteString(strings.Join(context, ","))
	convertedBytes.WriteString(strings.Join(typ, ","))
	convertedBytes.Write(vca.ToByte())
	convertedBytes.WriteString(vp.Holder)
	convertedBytes.Write(vp.Proof.ToByte())

	return convertedBytes.Bytes()
}

func (vp *VerifiablePresentation) Signature(holderPrivKey *eth.PrivateKey) error {
	vp.Proof.JWSSignature = ""

	signatureData, err := holderPrivKey.Sha256JWSSignature(vp.ToByte())
	if err != nil {
		return err
	}

	vp.Proof.JWSSignature = signatureData
	return nil
}

func (vp *VerifiablePresentation) Verify(holderPubKey *eth.PublicKey, issuerPubKey *eth.PublicKey) (isVerifyVP bool, err error) {
	copiedVP := *vp
	copiedVP.Proof.JWSSignature = ""

	switch vp.Proof.Type {
	case Secp256k1Sig:
		isVerifyVP, err = holderPubKey.VerifySha256JWSSignature(copiedVP.ToByte(), vp.Proof.JWSSignature)
	case Secp256ETHSig:
		isVerifyVP, err = holderPubKey.VerifyEcdsaSecp256ETHSignature2022Signature(copiedVP.ToByte(), vp.Proof.JWSSignature)
	}
	if err != nil {
		return false, err
	}

	if !isVerifyVP {
		return false, nil
	}

	return vp.VerifiableCredential.Verify(issuerPubKey)
}
