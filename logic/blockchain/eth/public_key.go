package eth

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	gojose "gopkg.in/square/go-jose.v2"
	"strconv"
)

type PublicKey struct {
	key *ecdsa.PublicKey
}

func NewPublicKey(hex string) (*PublicKey, error) {
	decode, err := hexutil.Decode(hex)
	if err != nil {
		return nil, err
	}

	pubKey, err := crypto.UnmarshalPubkey(decode)
	if err != nil {
		return nil, err
	}

	return &PublicKey{key: pubKey}, nil
}

func NewPublicKeyFromByte(data []byte) (*PublicKey, error) {
	pubKey, err := crypto.UnmarshalPubkey(data)
	if err != nil {
		return nil, err
	}

	return &PublicKey{key: pubKey}, nil
}

func NewPublicKeyFromSignature(oriData, signature []byte) (*PublicKey, error) {
	hash := crypto.Keccak256Hash(oriData)
	publicKeyByte, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return nil, err
	}

	return NewPublicKeyFromByte(publicKeyByte)
}

func (p *PublicKey) VerifySignature(data, signature []byte) bool {
	hash := crypto.Keccak256Hash(data)

	signatureNoRecoverID := signature[:len(signature)-1]
	return crypto.VerifySignature(crypto.FromECDSAPub(p.key), hash.Bytes(), signatureNoRecoverID)
}

func (p *PublicKey) Hex() []byte {
	return crypto.FromECDSAPub(p.key)
}

func (p *PublicKey) HexString() string {
	return hexutil.Encode(p.Hex())
}

func (p *PublicKey) Address() common.Address {
	return crypto.PubkeyToAddress(*p.key)
}

func (p *PublicKey) VerifySha256JWSSignature(data []byte, signature string) (bool, error) {
	hash := sha256.Sum256(data)

	sigObject, err := gojose.ParseDetached(signature, hash[:])
	if err != nil {
		return false, err
	}

	result, err := sigObject.Verify(p.key)
	if err != nil {
		return false, err
	}

	return bytes.Equal(hash[:], result), nil
}

func (p *PublicKey) EllipticMarshal() []byte {
	return crypto.FromECDSAPub(p.key)
}

func (p *PublicKey) VerifyEcdsaSecp256ETHSignature2022Signature(toByte []byte, signature string) (bool, error) {
	signDataHead := append([]byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(toByte))))
	signHash := crypto.Keccak256Hash(append(signDataHead, toByte...)).String()

	signPublicKey, err := GetPublicKeyUseEthSign(signHash, signature)
	if err != nil {
		return false, err
	}

	return bytes.Equal(signPublicKey.Address().Bytes(), p.Address().Bytes()), nil
}
