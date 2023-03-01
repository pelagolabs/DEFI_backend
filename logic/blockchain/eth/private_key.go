package eth

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	gojose "gopkg.in/square/go-jose.v2"
	"math/big"
	"strconv"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func NewPrivateKey(hex string) (*PrivateKey, error) {
	priKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		return nil, err
	}

	return &PrivateKey{key: priKey}, nil
}

func NewPrivateKeyFromECDSA(key *ecdsa.PrivateKey) *PrivateKey {
	return &PrivateKey{key: key}
}

func (p *PrivateKey) Signature(data []byte) ([]byte, error) {
	hash := crypto.Keccak256Hash(data)

	return crypto.Sign(hash.Bytes(), p.key)
}

func (p *PrivateKey) PublicKey() *PublicKey {
	pub := p.key.Public()

	if publicKeyECDSA, ok := pub.(*ecdsa.PublicKey); ok {
		return &PublicKey{key: publicKeyECDSA}
	} else {
		return nil
	}
}

func (p *PrivateKey) Hex() []byte {
	return crypto.FromECDSA(p.key)
}

func (p *PrivateKey) HexString() string {
	return hexutil.Encode(p.Hex())
}

func (p *PrivateKey) Address() common.Address {
	return p.PublicKey().Address()
}

func (p *PrivateKey) Sha256JWSSignature(fromByte []byte) (string, error) {
	hash := sha256.Sum256(fromByte)

	signer, err := gojose.NewSigner(gojose.SigningKey{Algorithm: gojose.ES256, Key: p.key}, nil)
	if err != nil {
		return "", err
	}

	signature, err := signer.Sign(hash[:])
	if err != nil {
		return "", err
	}

	compactserialized, err := signature.DetachedCompactSerialize()
	if err != nil {
		return "", err
	}
	return compactserialized, nil
}

func (p *PrivateKey) ETHSignature(fromByte []byte) (string, error) {
	signData := append([]byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(fromByte))))

	signature, err := p.Signature(append(signData, fromByte...))
	if err != nil {
		return "", err
	}

	if len(signature) != 65 {
		return "", fmt.Errorf("sign error")
	}

	signature[64] += 0x1b

	return hexutil.Encode(signature), nil
}

func (p *PrivateKey) MakeTransact(chainId *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(p.key, chainId)
}
