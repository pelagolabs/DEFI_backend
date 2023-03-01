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
	TaskTypeChainSwitch TaskType = "ChainSwitch"
)

func init() {
	DefaultManage.register(TaskTypeChainSwitch, func(taskId uint, data string) (task, error) {
		conf := &ChainSwitchTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &chainSwitch{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type chainSwitch struct {
	taskId uint
	conf   *ChainSwitchTaskConfig
}

type ChainSwitchTaskConfig struct {
	ChainId   uint   `json:"chain_id"`
	NewStatus string `json:"new_status"`
}

func (c *ChainSwitchTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

func (p *chainSwitch) Run(op *taskOperator) error {
	if p.conf.ChainId == 0 {
		return errors.New("invalid chain id")
	}

	// get chain
	chain, findErr := db.FindBlockchainById(p.conf.ChainId)
	if findErr != nil || chain.ID == 0 {
		return errors.New("can not found chain")
	}

	// get all related lp
	pools, findPoolErr := db.FindPoolWithCurrencyByStatus([]string{"available"})
	if findPoolErr != nil {
		return errors.New("get all related lp failed")
	}

	needOperateLpSet := make(map[string]bool, 0)
	for _, pool := range pools {
		if pool.Currency.ChainId != p.conf.ChainId {
			continue
		}

		if _, ok := needOperateLpSet[pool.Address]; !ok {
			needOperateLpSet[pool.Address] = true
		}
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

	oriChainStatus := chain.Status
	chain.Status = p.conf.NewStatus

	saveErr := db.SaveBlockchainWithStatusLock(chain, oriChainStatus, op.tx)
	if saveErr != nil {
		log.GetLogger().Warn("[task ChainSwitch] update pool status failed", zap.Error(saveErr))
		return saveErr
	}

	// operate on chain
	var lpOperateErr error
	for lpAddress, _ := range needOperateLpSet {
		lpOperateErr = contract.DefaultContract.LpSwitch(lpAddress, "", action, false, true)
	}

	lpOperateErr = contract.DefaultContract.FarmSwitch(chain.FarmAddress, action)

	if lpOperateErr != nil {
		return lpOperateErr
	} else {
		return nil
	}
}

func (p *chainSwitch) OnFail(op *taskOperator, err error) error {
	return op.RollbackTaskDB()
}

func (p *chainSwitch) TaskId() uint {
	return p.taskId
}
