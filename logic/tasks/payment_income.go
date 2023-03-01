package tasks

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

const (
	TaskTypePaymentIncome TaskType = "PaymentIncome"
)

func init() {
	DefaultManage.register(TaskTypePaymentIncome, func(taskId uint, data string) (task, error) {
		conf := &PaymentIncomeTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &paymentIncomeTask{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type PaymentIncomeTaskConfig struct {
	From      common.Address
	To        common.Address
	Amount    *types.BigInt
	UniqueKey string
	Raw       []byte
}

func (c *PaymentIncomeTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

type paymentIncomeTask struct {
	taskId uint
	conf   *PaymentIncomeTaskConfig
}

func (p *paymentIncomeTask) Run(op *taskOperator) error {
	payment, err := db.FindPaymentByCollectionAddress(p.conf.To.String())
	if err != nil {
		return err
	}

	paymentTx := &db.PaymentTx{
		TxAddress: p.conf.From.String(),
		Amount:    p.conf.Amount,
		Hash:      p.conf.UniqueKey,
		EventData: p.conf.Raw,
	}

	if payment != nil && payment.ID > 0 {
		oriStatus := payment.Status
		if oriStatus == db.PaymentStatusCreated || oriStatus == db.PaymentStatusPending {
			paymentTx.PaymentId = &payment.ID

			// p.conf.Value * payment.CurrencyPriceInUSD * 100 / (10 ^ payment.Currency.DecimalCount)
			paymentTx.AmountInCent = types.NewBigFloatUseBigInt(p.conf.Amount).
				Mul(types.NewBigFloatFast(payment.CurrencyPriceInUSD)).
				Mul(types.NewBigFloatFast(100)).
				Div(types.NewBigFloatUseBigInt(types.NewBigIntFast(10).Pow(types.NewBigIntFast(int64(payment.Currency.DecimalCount))))).
				RoundToInt().
				Uint64()

			payment.CollectionAmount.Add(p.conf.Amount)

			// payment.CollectionAmount * payment.CurrencyPriceInUSD * 100 / (10 ^ payment.Currency.DecimalCount)
			payment.CollectionAmountInCent = types.NewBigFloatUseBigInt(payment.CollectionAmount).
				Mul(types.NewBigFloatFast(payment.CurrencyPriceInUSD)).
				Mul(types.NewBigFloatFast(100)).
				Div(types.NewBigFloatUseBigInt(types.NewBigIntFast(10).Pow(types.NewBigIntFast(int64(payment.Currency.DecimalCount))))).
				RoundToInt().
				Uint64()

			// round(payment.Amount * (100 - payment.Slippage) / 100)
			minPayAmount := types.NewBigFloatUseBigInt(payment.Amount).
				Mul(types.NewBigFloatFast((100 - payment.Slippage) / 100)).
				RoundToInt()

			if payment.CollectionAmount.Cmp(minPayAmount) >= 0 {
				payment.Status = db.PaymentStatusSuccess
				payment.PlatformFeeAmount = payment.CollectionAmount.Copy().Mul(types.NewBigIntFast(5)).Div(types.NewBigIntFast(10000))
				payment.PoolFeeAmount = payment.CollectionAmount.Copy().Mul(types.NewBigIntFast(25)).Div(types.NewBigIntFast(10000))

				// round(payment.Amount * (100 + payment.Slippage) / 100)
				maxPayAmount := types.NewBigFloatUseBigInt(payment.Amount).
					Mul(types.NewBigFloatFast((100 + payment.Slippage) / 100)).
					RoundToInt()

				if payment.CollectionAmount.Cmp(maxPayAmount) > 0 {
					payment.AmountStatus = db.AmountStatusOverpaid
				} else {
					payment.AmountStatus = db.AmountStatusPaid
				}

				err = contract.DefaultContract.ReturnAddressToPool(payment.Currency, payment.CollectionAddress)
				if err != nil {
					log.GetLogger().Warn("address return address pool fail", zap.Error(err))
					return nil
				}

				err = op.NewSubTask(TaskTypePaymentFee, &PaymentFeeTaskConfig{
					ApplicationId:     payment.ApplicationId,
					Date:              time.Now().Format("2006-01-02"),
					CurrencyId:        payment.CurrencyId,
					MerchantAmount:    payment.CollectionAmount.Copy().Sub(payment.PoolFeeAmount).Sub(payment.PlatformFeeAmount).String(),
					PlatformFeeAmount: payment.PlatformFeeAmount.String(),
					PoolFeeAmount:     payment.PoolFeeAmount.String(),
				})
				if err != nil {
					return err
				}

				err = op.NewSubTask(TaskTypePaymentCreateVC, &PaymentTaskConfig{PaymentId: payment.ID})
				if err != nil {
					return err
				}
			} else {
				payment.Status = db.PaymentStatusPending
				payment.AmountStatus = db.AmountStatusUnderpaid
			}

			err = db.SavePaymentWithStatusLock(payment, oriStatus, op.tx)
			if err != nil {
				return err
			}
		}
	}

	return db.SavePaymentTx(paymentTx, op.tx)
}

func (p *paymentIncomeTask) TaskId() uint {
	return p.taskId
}
