package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type LpSwitchInfo struct {
	LpAddress        common.Address `json:"lp_address"`
	CurrencyAddress  common.Address `json:"currency_address"`
	IsNativeCurrency bool           `json:"is_native_currency"`
	IsGlobal         bool           `json:"is_global"`
	Action           string         `json:"action"`
}

func (e *LpSwitchInfo) MarshalBinary() (data []byte, err error) {
	return rlp.EncodeToBytes(e)
}

func (e *LpSwitchInfo) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, &e)
}
