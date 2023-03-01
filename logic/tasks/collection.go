package tasks

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/blockchain/contract/chain/ethereum"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

const (
	TaskTypeCollection TaskType = "Collection"
)

func init() {
	DefaultManage.register(TaskTypeCollection, func(taskId uint, data string) (task, error) {
		conf := &CollectionTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &collectionTask{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type CollectionTaskConfig struct {
	ManagerId uint
	Remark    string
	FeeType   string

	FeeIds      []uint
	ToAddress   common.Address
	TotalAmount *big.Int
	Chain       string
	Currency    string
}

func (c *CollectionTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

type collectionTask struct {
	taskId uint
	conf   *CollectionTaskConfig
}

func (p *collectionTask) Run(op *taskOperator) error {
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
		FeeType:    p.conf.FeeType,
		CurrencyId: currency.ID,
		ToAddress:  p.conf.ToAddress.String(),
		Amount:     types.NewBigInt(p.conf.TotalAmount),
		ManagerId:  p.conf.ManagerId,
		Remark:     p.conf.Remark,
	})
	if err != nil {
		return err
	}

	if currency.IsNative {
		return contract.DefaultContract.WithdrawNative(&ethereum.WithdrawInfo{
			WalletAddress:   common.HexToAddress(wallet),
			WithdrawAddress: p.conf.ToAddress,
			WithdrawAmount:  p.conf.TotalAmount,
		})
	} else {
		return contract.DefaultContract.Withdraw(&ethereum.WithdrawInfo{
			WalletAddress:   common.HexToAddress(wallet),
			WithdrawAddress: p.conf.ToAddress,
			CurrencyAddress: common.HexToAddress(currency.ContractAddress),
			WithdrawAmount:  p.conf.TotalAmount,
		})
	}
}

func (p *collectionTask) OnFail(op *taskOperator, err error) error {
	err = op.RollbackTaskDB()
	if err != nil {
		return err
	}

	return db.UpdatePaymentFeesFrozenAmountToZero(p.conf.FeeIds, op.tx)
}

func (p *collectionTask) TaskId() uint {
	return p.taskId
}
