package tasks

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/blockchain/contract/chain/dynamic"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

const (
	TaskTypeAddRewards TaskType = "AddRewards"
)

func init() {
	DefaultManage.register(TaskTypeAddRewards, func(taskId uint, data string) (task, error) {
		conf := &AddRewardsTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &AddRewardsTask{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type AddRewardsTaskConfig struct {
	ManagerId uint
	Remark    string
	FeeType   string

	FeeIds      []uint
	TotalAmount *big.Int
	Chain       string
	Currency    string
}

func (c *AddRewardsTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

type AddRewardsTask struct {
	taskId uint
	conf   *AddRewardsTaskConfig
}

func (p *AddRewardsTask) Run(op *taskOperator) error {
	blockchain, err := db.FindBlockchainByContractName(p.conf.Chain)
	if err != nil {
		return err
	}
	if blockchain.ID == 0 {
		return nil
	}

	currency, err := db.FindCurrencyByChainAndSymbol(blockchain.ID, p.conf.Currency)
	if err != nil {
		return err
	}
	if currency.ID == 0 {
		return nil
	}

	wallets, err := db.QueryWalletUseChainName(p.conf.Chain)
	if err != nil {
		return err
	}

	pool, err := db.FindPoolByCurrencyId(currency.ID)
	if err != nil || pool.ID == 0 {
		return errors.New("can not found pool")
	}

	var wallet string
	if currency.IsNative {
		wallet, err = contract.DefaultContract.QueryEnoughWalletsMaxNativeAmount(wallets, false, currency, p.conf.TotalAmount)
	} else {
		wallet, err = contract.DefaultContract.QueryEnoughWalletsMaxAmount(wallets, false, currency, p.conf.TotalAmount)
	}

	if err != nil {
		return err
	}

	fees, err := db.FindPaymentFeesById(p.conf.FeeIds)
	if err != nil {
		return err
	}

	for _, fee := range fees {
		if err = db.AddPaymentFeeClaimedAmountAndSubFrozenAmount(fee, op.tx); err != nil {
			return err
		}
	}

	err = db.SaveFeeWithdrawLog(&db.FeeWithdrawLog{
		FeeType:    db.PaymentFeeFeeTypePool,
		CurrencyId: currency.ID,
		ToAddress:  pool.Address,
		Amount:     types.NewBigInt(p.conf.TotalAmount),
		ManagerId:  p.conf.ManagerId,
		Remark:     p.conf.Remark,
	})
	if err != nil {
		return err
	}

	return contract.DefaultContract.AddRewards(&dynamic.AddRewardsInfo{
		LpAddress:          common.HexToAddress(pool.Address),
		WalletAddress:      common.HexToAddress(wallet),
		CurrencyAddress:    common.HexToAddress(currency.ContractAddress),
		AddRewardsAmount:   p.conf.TotalAmount,
		AddRewardsNative:   currency.IsNative,
		BlockchainName:     blockchain.Name,
		BlockchainEndpoint: blockchain.Endpoint,
	})
}

func (p *AddRewardsTask) OnFail(op *taskOperator, err error) error {
	err = op.RollbackTaskDB()
	if err != nil {
		return err
	}

	return db.UpdatePaymentFeesFrozenAmountToZero(p.conf.FeeIds, op.tx)
}

func (p *AddRewardsTask) TaskId() uint {
	return p.taskId
}
