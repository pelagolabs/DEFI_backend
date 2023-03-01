package dynamic

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type AddRewardsInfo struct {
	BlockchainEndpoint string
	BlockchainName     string
	AddRewardsNative   bool
	AddRewardsAmount   *big.Int
	CurrencyAddress    common.Address
	WalletAddress      common.Address
	LpAddress          common.Address
}

func (e *AddRewardsInfo) MarshalBinary() (data []byte, err error) {
	return rlp.EncodeToBytes(e)
}

func (e *AddRewardsInfo) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, &e)
}
