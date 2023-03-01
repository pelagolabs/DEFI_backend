package tasks

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/blockchain/contract/chain/ethereum"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

const (
	TaskTypeWithdraw TaskType = "Withdraw"
)

func init() {
	DefaultManage.register(TaskTypeWithdraw, func(taskId uint, data string) (task, error) {
		conf := &WithdrawTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &withdrawTask{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type WithdrawTaskConfig struct {
	VCId        []uint
	ToAddress   common.Address
	TotalAmount *big.Int
	Chain       string
	Currency    string
}

func (c *WithdrawTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

type withdrawTask struct {
	taskId uint
	conf   *WithdrawTaskConfig
}

func (p *withdrawTask) Run(op *taskOperator) error {
	blockchain, err := db.FindBlockchainByContractName(p.conf.Chain)
	if err != nil {
		return err
	}
	if blockchain.ID == 0 {
		return nil
	}

	if blockchain.Status == "unavailable" {
		return errors.New("blockchain was suspend")
	}

	currency, err := db.FindCurrencyByChainAndSymbol(blockchain.ID, p.conf.Currency)
	if err != nil {
		return err
	}
	if currency.ID == 0 {
		return errors.New("currency not found")
	}

	if currency.Status == "unavailable" {
		return errors.New("currency was suspend")
	}

	wallets, err := db.QueryWalletUseChainName(p.conf.Chain)
	if err != nil {
		return err
	}

	vcs, err := db.FindVCsById(p.conf.VCId)
	if err != nil {
		return err
	}

	if len(vcs) == 0 {
		return errors.New("no available vc")
	}

	for _, vc := range vcs {
		vc.VCStatus = db.VCStatusWithdraw
		if err = db.UpdateVC(vc, db.VCStatusProcessing, op.tx); err != nil {
			return err
		}
	}

	firstVC, err := db.FindVCById(vcs[0].ID, db.WithPreload("Payment.Currency"))
	if err != nil {
		return err
	}

	withdrawData := &db.Withdraw{
		MerchantId:    firstVC.Payment.MerchantId,
		ApplicationId: firstVC.Payment.ApplicationId,
		CurrencyId:    firstVC.Payment.CurrencyId,
		WithdrawNum:   util.RandString(16),
		Amount:        p.conf.TotalAmount.String(),
		Status:        db.WithdrawStatusCreated,
	}

	amountInCent, err := exchange.DefaultManage.ExchangeCoinToUSD(p.conf.TotalAmount, firstVC.Payment.Currency)
	if err == nil {
		amountFloat, _ := amountInCent.Float64()
		withdrawData.AmountInCent = uint64(amountFloat * 100)
	}

	err = db.SaveWithdraw(withdrawData, op.tx)
	if err != nil {
		return err
	}

	var wallet string
	if currency.IsNative {
		wallet, err = contract.DefaultContract.QueryEnoughWalletsMaxNativeAmount(wallets, true, currency, p.conf.TotalAmount)
	} else {
		wallet, err = contract.DefaultContract.QueryEnoughWalletsMaxAmount(wallets, true, currency, p.conf.TotalAmount)
	}
	if err != nil {
		return err
	}

	err = p.balanceRecord(firstVC.Payment.ApplicationId, firstVC.Payment.CurrencyId, p.conf.TotalAmount, op)
	if err != nil {
		return err
	}

	if currency.IsNative {
		return contract.DefaultContract.WithdrawNative(&ethereum.WithdrawInfo{
			WithdrawId:      withdrawData.ID,
			WalletAddress:   common.HexToAddress(wallet),
			WithdrawAddress: p.conf.ToAddress,
			WithdrawAmount:  p.conf.TotalAmount,
		})
	} else {
		return contract.DefaultContract.Withdraw(&ethereum.WithdrawInfo{
			WithdrawId:      withdrawData.ID,
			WalletAddress:   common.HexToAddress(wallet),
			WithdrawAddress: p.conf.ToAddress,
			CurrencyAddress: common.HexToAddress(currency.ContractAddress),
			WithdrawAmount:  p.conf.TotalAmount,
		})
	}
}

func (p *withdrawTask) balanceRecord(applicationId uint, currencyId uint, amount *big.Int, op *taskOperator) error {
	dateIns := time.Now()
	date := dateIns.Format("2006-01-02")
	balances, err := db.FindPaymentBalanceByApplicationIdAndDateAndCurrencyId(applicationId, date, currencyId)
	if err != nil {
		return err
	}

	if balances.ID == 0 {
		return db.SavePaymentBalance(&db.PaymentBalance{
			Date:          dateIns,
			ApplicationId: applicationId,
			CurrencyId:    currencyId,
			TotalAmount:   types.NewBigIntFast(0),
			ClaimedAmount: types.NewBigInt(amount),
		})
	} else {
		return db.UpdatePaymentBalanceClaimedAmount(balances, balances.ClaimedAmount.Copy().Add(types.NewBigInt(amount)), op.tx)
	}
}

func (p *withdrawTask) OnFail(op *taskOperator, err error) error {
	return op.RollbackTaskDB()
}

func (p *withdrawTask) TaskId() uint {
	return p.taskId
}
