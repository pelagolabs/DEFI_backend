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
	TaskTypePoolSwitch TaskType = "PoolSwitch"
)

func init() {
	DefaultManage.register(TaskTypePoolSwitch, func(taskId uint, data string) (task, error) {
		conf := &PoolSwitchTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &poolSwitch{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type poolSwitch struct {
	taskId uint
	conf   *PoolSwitchTaskConfig
}

type PoolSwitchTaskConfig struct {
	PoolId    uint   `json:"pool_id"`
	NewStatus string `json:"new_status"`
}

func (c *PoolSwitchTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

func (p *poolSwitch) Run(op *taskOperator) error {
	if p.conf.PoolId == 0 {
		return errors.New("invalid pool id")
	}

	// get pool
	pool, findErr := db.FindPoolWithCurrencyById(p.conf.PoolId)
	if findErr != nil || pool.ID == 0 {
		return errors.New("can not found pool")
	}

	action := ""
	switch p.conf.NewStatus {
	case "available":
		action = "unpause"
	case "unavailable":
		action = "pause"
	default:
		return errors.New("invalid new status")
	}

	oriPoolStatus := pool.Status
	pool.Status = p.conf.NewStatus

	saveErr := db.SavePoolWithStatusLock(pool, oriPoolStatus, op.tx)
	if saveErr != nil {
		log.GetLogger().Warn("[task PoolSwitch] update pool status failed", zap.Error(saveErr))
		return saveErr
	}

	return contract.DefaultContract.LpSwitch(pool.Address, pool.Currency.ContractAddress, action, pool.Currency.IsNative, false)
}

func (p *poolSwitch) OnFail(op *taskOperator, err error) error {
	return op.RollbackTaskDB()
}

func (p *poolSwitch) TaskId() uint {
	return p.taskId
}
