package tasks

import (
	"encoding/json"
	"errors"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/db"
)

const (
	TaskTypeFarmSwitch TaskType = "FarmSwitch"
)

func init() {
	DefaultManage.register(TaskTypeFarmSwitch, func(taskId uint, data string) (task, error) {
		conf := &FarmSwitchTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &farmSwitch{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type farmSwitch struct {
	taskId uint
	conf   *FarmSwitchTaskConfig
}

type FarmSwitchTaskConfig struct {
	ChainId   uint   `json:"chain_id"`
	NewStatus string `json:"new_status"`
}

func (c *FarmSwitchTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

func (p *farmSwitch) Run(op *taskOperator) error {
	if p.conf.ChainId == 0 {
		return errors.New("invalid pool id")
	}

	// get chain
	chain, findErr := db.FindBlockchainById(p.conf.ChainId)
	if findErr != nil || chain.ID == 0 || chain.FarmAddress == "" {
		return errors.New("can not found chain or farm contract")
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

	return contract.DefaultContract.FarmSwitch(chain.FarmAddress, action)
}

func (p *farmSwitch) OnFail(op *taskOperator, err error) error {
	return op.RollbackTaskDB()
}

func (p *farmSwitch) TaskId() uint {
	return p.taskId
}
