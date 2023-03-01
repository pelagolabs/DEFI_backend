package tasks

import (
	"encoding/json"
	"go.uber.org/zap"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

const (
	TaskTypePaymentTimeout TaskType = "PaymentTimeout"
)

func init() {
	DefaultManage.register(TaskTypePaymentTimeout, func(taskId uint, data string) (task, error) {
		conf := &PaymentTimeoutTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &paymentTimeoutTask{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type PaymentTimeoutTaskConfig struct {
	PaymentId uint
}

func (c *PaymentTimeoutTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

type paymentTimeoutTask struct {
	taskId uint
	conf   *PaymentTimeoutTaskConfig
}

func (p *paymentTimeoutTask) Run(op *taskOperator) error {
	payment, err := db.FindPaymentById(p.conf.PaymentId)
	if err != nil {
		return err
	}

	oriStatus := payment.Status
	payment.Status = db.PaymentStatusClosed

	if !payment.CollectionAmount.IsZero() {
		payment.PlatformFeeAmount = payment.CollectionAmount.Copy().Mul(types.NewBigIntFast(5)).Div(types.NewBigIntFast(10000))
		payment.PoolFeeAmount = payment.CollectionAmount.Copy().Mul(types.NewBigIntFast(25)).Div(types.NewBigIntFast(10000))

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
	}

	err = db.SavePaymentWithStatusLock(payment, oriStatus, op.tx)
	if err != nil {
		log.GetLogger().Warn("[cron ProcessDeadPayment] close payment error", zap.Error(err), zap.Any("payment", payment), zap.String("oriStatus", oriStatus))
		return err
	}

	return contract.DefaultContract.ReturnAddressToPool(payment.Currency, payment.CollectionAddress)
}

func (p *paymentTimeoutTask) TaskId() uint {
	return p.taskId
}
