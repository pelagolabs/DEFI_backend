package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type FarmSwitchInfo struct {
	FarmAddress common.Address `json:"farm_address"`
	Action      string         `json:"action"`
}

func (e *FarmSwitchInfo) MarshalBinary() (data []byte, err error) {
	return rlp.EncodeToBytes(e)
}

func (e *FarmSwitchInfo) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, &e)
}
