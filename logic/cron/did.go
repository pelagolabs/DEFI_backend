package cron

import (
	"go.uber.org/zap"
	"veric-backend/internal/log"
	"veric-backend/logic/db"
	"veric-backend/logic/tasks"
)

func init() {
	forceAddFunc("@every 5m", ProcessCreatedDid)
}

func ProcessCreatedDid() {
	users, userErr := db.GetUserByDidUpStatus(db.DidUpStatusCreated)
	if userErr != nil {
		log.GetLogger().Warn("[cron ProcessCreatedDid] get user error", zap.Error(userErr))
		return
	}

	for _, user := range users {
		if user.Did == "" || user.DidPubKey == "" {
			continue
		}

		taskErr := tasks.DefaultManage.NewTask(tasks.TaskTypeDidRegistry, &tasks.DidRegistryTaskConfig{
			Did:    user.Did,
			PubKey: user.DidPubKey,
		})
		if taskErr != nil {
			log.GetLogger().Warn("[cron ProcessCreatedDid] did registry error", zap.Error(taskErr), zap.Any("did", user.Did))
			continue
		}
	}
}
