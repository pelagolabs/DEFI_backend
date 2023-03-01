package did

import (
	"bytes"
	"time"
)

type VCProof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	VerificationMethod string `json:"verificationMethod"`
	ProofPurpose       string `json:"proofPurpose"`
	JWSSignature       string `json:"jws"` //signature is created from a hash of the VC
}

func CreateProof(vm string) VCProof {
	vcProof := VCProof{}
	vcProof.Type = Secp256k1Sig
	vcProof.VerificationMethod = vm
	vcProof.JWSSignature = ""
	vcProof.Created = time.Now().Format(time.RFC3339)
	vcProof.ProofPurpose = PurposeAuth
	return vcProof
}

func (v *VCProof) ToByte() []byte {
	var convertedBytes bytes.Buffer
	convertedBytes.WriteString(v.Type)
	convertedBytes.WriteString(v.Created)
	convertedBytes.WriteString(v.VerificationMethod)
	convertedBytes.WriteString(v.ProofPurpose)
	convertedBytes.WriteString(v.JWSSignature)

	return convertedBytes.Bytes()
}
