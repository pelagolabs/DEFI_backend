package tasks

import (
	"encoding/json"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/did"
	"veric-backend/logic/db"
)

const (
	TaskTypePaymentCreateVC TaskType = "PaymentCreateVC"
)

func init() {
	DefaultManage.register(TaskTypePaymentCreateVC, func(taskId uint, data string) (task, error) {
		conf := &PaymentTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &paymentTaskCreateVC{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type PaymentTaskConfig struct {
	PaymentId uint
}

func (c *PaymentTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

type paymentTaskCreateVC struct {
	taskId uint
	conf   *PaymentTaskConfig
}

func (p *paymentTaskCreateVC) Run(op *taskOperator) error {
	payment, err := db.FindPaymentById(p.conf.PaymentId)
	if err != nil {
		return err
	}

	vcId := util.RandString(32)
	vc := did.CreateVerifiableCredential(vcId, did.IssueDID, &did.VCSubjectDeposit{
		Chain:             payment.Currency.Blockchain.ContractName,
		Currency:          payment.Currency.Symbol,
		Amount:            payment.CollectionAmount.String(),
		PlatformFeeAmount: payment.PlatformFeeAmount.String(),
		PoolFeeAmount:     payment.PoolFeeAmount.String(),
		MerchantAmount:    payment.CollectionAmount.Copy().Sub(payment.PoolFeeAmount).Sub(payment.PlatformFeeAmount).String(),
	})

	err = vc.Signature(did.IssuePriKey)
	if err != nil {
		return err
	}

	err = db.SaveVC(&db.VC{
		PaymentId:  payment.ID,
		MerchantId: payment.MerchantId,
		VCID:       vcId,
		VCContent:  vc.ToJson(),
		VCStatus:   db.VCStatusCreated,
	})
	if err != nil {
		return err
	}

	if payment.Status == db.PaymentStatusSuccess {
		return op.NewSubTask(TaskTypePaymentSuccessCallback, p.conf)
	} else {
		return nil
	}
}

func (p *paymentTaskCreateVC) TaskId() uint {
	return p.taskId
}
