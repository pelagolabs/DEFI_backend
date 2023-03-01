package eth

import (
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strconv"
	"strings"
)

// CheckEthSign check whether Ethereum signature correct
func CheckEthSign(ethAddress string, originData string, ethSign string) bool {
	pubKey, err := GetPublicKeyUseEthSign(originData, ethSign)
	if err != nil {
		return false
	}

	// compare the address
	if strings.ToUpper(pubKey.Address().String()) != strings.ToUpper(ethAddress) {
		return false
	} else {
		return true
	}
}

func GetPublicKeyUseEthSign(originData string, ethSign string) (*PublicKey, error) {
	// contact the origin data
	signData := append([]byte("\x19Ethereum Signed Message:\n"+strconv.Itoa(len(originData))), originData...)
	decodeETHSignature, err := hexutil.Decode(ethSign)
	if err != nil {
		return nil, err
	}
	if decodeETHSignature[64] != 27 && decodeETHSignature[64] != 28 {
		return nil, errors.New("invalid signature")
	}
	decodeETHSignature[64] -= 27

	// create public key from sign and origin data
	pubKey, err := NewPublicKeyFromSignature(signData, decodeETHSignature)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}
