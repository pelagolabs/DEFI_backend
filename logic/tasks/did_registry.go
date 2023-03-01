package tasks

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/db"
)

const (
	TaskTypeDidRegistry TaskType = "DidRegistry"
)

func init() {
	DefaultManage.register(TaskTypeDidRegistry, func(taskId uint, data string) (task, error) {
		conf := &DidRegistryTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &didRegistry{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type didRegistry struct {
	taskId uint
	conf   *DidRegistryTaskConfig
}

func (p *didRegistry) Run(op *taskOperator) error {
	user, userErr := db.FindUserByDid(p.conf.Did)
	if userErr != nil {
		return nil
	}

	if user == nil || user.ID == 0 {
		return errors.New("can not found user")
	}

	if user.DidUpStatus != db.DidUpStatusCreated {
		return nil
	}

	oriDidStatus := user.DidUpStatus
	user.DidUpStatus = db.DidUpStatusPending
	saveErr := db.SaveUserWithDidStatusLock(user, oriDidStatus, op.tx)
	if saveErr != nil {
		log.GetLogger().Warn("[cron DidRegistry] update user did status failed", zap.Error(saveErr))
		return saveErr
	}

	return contract.DefaultContract.DidRegistry(user.Did, user.DidPubKey)
}

func (p *didRegistry) OnFail(op *taskOperator, err error) error {
	return op.RollbackTaskDB()
}

func (p *didRegistry) TaskId() uint {
	return p.taskId
}
