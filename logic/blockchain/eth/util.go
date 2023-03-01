package eth

import "github.com/ethereum/go-ethereum/common"

func IsValidEthAddress(address string) bool {
	return common.IsHexAddress(address)
}

func ToEthAddress(address string) string {
	return common.HexToAddress(address).String()
}
