package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type WithdrawInfo struct {
	WalletAddress   common.Address
	WithdrawAddress common.Address
	CurrencyAddress common.Address
	WithdrawAmount  *big.Int
	WithdrawId      uint
}

func (e *WithdrawInfo) MarshalBinary() (data []byte, err error) {
	return rlp.EncodeToBytes(e)
}

func (e *WithdrawInfo) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, &e)
}
