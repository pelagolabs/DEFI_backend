package did

import (
	"errors"
	"sort"
)

const (
	PurposeAuth   = "Authentication"
	Secp256k1Sig  = "EcdsaSecp256k1Signature2019"
	Secp256ETHSig = "EcdsaSecp256ETHSignature2022"
	Secp256k1Key  = "EcdsaSecp256k1VerificationKey2019"

	ContextDID        = "https://w3id.org/did/v1"
	ContextCredential = "https://www.w3.org/2018/credentials/v1"
	ContextSecp256k1  = "https://ns.did.ai/suites/secp256k1-2019/v1/"

	TypeCredential   = "VerifiableCredential"
	TypeVericDeposit = "VericDeposit"

	DIDPrefix = "did:veric:"
)

var ErrUnknownProofType = errors.New("unable to verify unknown proof type")
var ErrSecp256k1WrongVMType = errors.New("must use a verification method with a type of 'EcdsaSecp256k1VerificationKey2019' to verify a 'EcdsaSecp256k1Signature2019' proof")
var ErrMissingVM = errors.New("failed to find verification method")

type CanToByte interface {
	ToByte() []byte
}

type CanToJson interface {
	ToJson() []byte
}

func cloneAndSort(from []string) []string {
	to := make([]string, len(from))
	copy(to, from)
	sort.Strings(to)
	return to
}
