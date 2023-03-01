package cron

import (
	"github.com/robfig/cron/v3"
	"veric-backend/internal/log"
)

var cronInstance = cron.New()

func Start() {
	log.GetLogger().Info("start cron...")

	cronInstance.Start()
}

func forceAddFunc(spec string, cmd func()) {
	_, err := cronInstance.AddFunc(spec, cmd)
	if err != nil {
		panic(err)
	}
}
