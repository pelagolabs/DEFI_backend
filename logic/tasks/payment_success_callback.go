package tasks

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"veric-backend/logic/db"
)

const (
	TaskTypePaymentSuccessCallback TaskType = "PaymentSuccessCallback"
)

func init() {
	DefaultManage.register(TaskTypePaymentSuccessCallback, func(taskId uint, data string) (task, error) {
		conf := &PaymentTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &paymentSuccessTaskCallback{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type paymentSuccessTaskCallback struct {
	taskId uint
	conf   *PaymentTaskConfig
}

func (p *paymentSuccessTaskCallback) Run(op *taskOperator) error {
	payment, err := db.FindPaymentById(p.conf.PaymentId)
	if err != nil {
		return err
	}

	if payment == nil || payment.ID == 0 {
		return errors.New("can not found payment")
	}

	app := payment.Application
	callbackUrl := payment.Application.CallbackUrl
	if callbackUrl == "" {
		return errors.New("not set callback url")
	}

	payment.Application = nil
	callbackData, err := json.Marshal(app)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, callbackUrl, bytes.NewReader(callbackData))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("status code is not 200")
	}

	return nil
}

func (p *paymentSuccessTaskCallback) TaskId() uint {
	return p.taskId
}
