package tasks

import (
	"encoding/json"
	"time"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

const (
	TaskTypePaymentFee TaskType = "PaymentFee"
)

func init() {
	DefaultManage.register(TaskTypePaymentFee, func(taskId uint, data string) (task, error) {
		conf := &PaymentFeeTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &paymentTaskFee{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type PaymentFeeTaskConfig struct {
	ApplicationId     uint
	Date              string
	CurrencyId        uint
	MerchantAmount    string
	PlatformFeeAmount string
	PoolFeeAmount     string
}

func (c *PaymentFeeTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

type paymentTaskFee struct {
	taskId uint
	conf   *PaymentFeeTaskConfig
}

func (p *paymentTaskFee) Run(op *taskOperator) error {
	err := p.updateFee(op, db.PaymentFeeFeeTypePool, p.conf.PoolFeeAmount)
	if err != nil {
		return err
	}

	err = p.updateFee(op, db.PaymentFeeFeeTypePlatform, p.conf.PlatformFeeAmount)
	if err != nil {
		return err
	}

	err = p.balanceRecord(op)
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentTaskFee) OnFail(op *taskOperator, err error) error {
	return op.RollbackTaskDB()
}

func (p *paymentTaskFee) balanceRecord(op *taskOperator) error {
	balances, err := db.FindPaymentBalanceByApplicationIdAndDateAndCurrencyId(p.conf.ApplicationId, p.conf.Date, p.conf.CurrencyId)
	if err != nil {
		return err
	}

	if feeAmount, err := types.NewBigIntString(p.conf.MerchantAmount); err == nil {
		if balances.ID == 0 {
			dateIns, err := time.Parse("2006-01-02", p.conf.Date)
			if err != nil {
				return err
			}
			return db.SavePaymentBalance(&db.PaymentBalance{
				Date:          dateIns,
				ApplicationId: p.conf.ApplicationId,
				CurrencyId:    p.conf.CurrencyId,
				TotalAmount:   feeAmount,
				ClaimedAmount: types.NewBigIntFast(0),
			})
		} else {
			return db.UpdatePaymentBalanceTotalAmount(balances, balances.TotalAmount.Copy().Add(feeAmount), op.tx)
		}
	} else {
		return err
	}
}

func (p *paymentTaskFee) updateFee(op *taskOperator, typ string, amount string) error {
	fees, err := db.FindPaymentFeeByDateAndFeeTypeAndCurrencyId(typ, p.conf.Date, p.conf.CurrencyId)
	if err != nil {
		return err
	}

	if feeAmount, err := types.NewBigIntString(amount); err == nil {
		if fees.ID == 0 {
			dateIns, err := time.Parse("2006-01-02", p.conf.Date)
			if err != nil {
				return err
			}
			return db.SavePaymentFee(&db.PaymentFee{
				Date:          dateIns,
				FeeType:       typ,
				CurrencyId:    p.conf.CurrencyId,
				TotalAmount:   feeAmount,
				FrozenAmount:  types.NewBigIntFast(0),
				ClaimedAmount: types.NewBigIntFast(0),
			})
		} else {
			return db.UpdatePaymentFeeTotalAmount(fees, fees.TotalAmount.Copy().Add(feeAmount), op.tx)
		}
	} else {
		return err
	}
}

func (p *paymentTaskFee) TaskId() uint {
	return p.taskId
}
