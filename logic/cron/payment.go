package cron

import (
	"go.uber.org/zap"
	"veric-backend/internal/log"
	"veric-backend/logic/db"
	"veric-backend/logic/tasks"
)

func init() {
	forceAddFunc("@every 1m", ProcessDeadPayment)
}

func ProcessDeadPayment() {
	payments, err := db.FindDeadPayment()
	if err != nil {
		log.GetLogger().Warn("[cron ProcessDeadPayment] find payment error", zap.Error(err))
		return
	}

	for _, payment := range payments {
		err = tasks.DefaultManage.NewTask(tasks.TaskTypePaymentTimeout, &tasks.PaymentTimeoutTaskConfig{PaymentId: payment.ID})
		if err != nil {
			log.GetLogger().Warn("[cron ProcessDeadPayment] close payment error", zap.Error(err), zap.Any("payment", payment))
			continue
		}
	}
}
