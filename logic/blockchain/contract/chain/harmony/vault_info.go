package harmony

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type VaultInfo struct {
	RandomId           string
	ChainName, Address string
}

func (e *VaultInfo) MarshalBinary() (data []byte, err error) {
	return rlp.EncodeToBytes(e)
}

func (e *VaultInfo) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, &e)
}
