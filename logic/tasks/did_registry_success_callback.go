package tasks

import (
	"encoding/json"
	"errors"
	"veric-backend/logic/db"
)

const (
	TaskTypeDidRegistrySuccessCallback TaskType = "DidRegistrySuccessCallback"
)

func init() {
	DefaultManage.register(TaskTypeDidRegistrySuccessCallback, func(taskId uint, data string) (task, error) {
		conf := &DidRegistryTaskConfig{}
		err := json.Unmarshal([]byte(data), conf)
		if err != nil {
			return nil, err
		}

		return &didRegistrySuccessCallback{
			taskId: taskId,
			conf:   conf,
		}, nil
	})
}

type DidRegistryTaskConfig struct {
	Did    string `json:"did"`
	PubKey string `json:"pub_key"`
	UpTx   string `json:"up_tx"`
}

func (c *DidRegistryTaskConfig) data() (string, error) {
	marshal, err := json.Marshal(c)
	return string(marshal), err
}

type didRegistrySuccessCallback struct {
	taskId uint
	conf   *DidRegistryTaskConfig
}

func (p *didRegistrySuccessCallback) Run(op *taskOperator) error {
	user, userErr := db.FindUserByDid(p.conf.Did)
	if userErr != nil {
		return nil
	}

	if user == nil || user.ID == 0 {
		return errors.New("can not found user")
	}

	if user.DidUpStatus != db.DidUpStatusPending {
		return nil
	}

	user.DidUpStatus = db.DidUpStatusSuccess
	user.DidUpTx = p.conf.UpTx

	saveErr := db.SaveUser(user)
	if saveErr != nil {
		return errors.New("update user failed. err: " + saveErr.Error())
	}

	return nil
}

func (p *didRegistrySuccessCallback) TaskId() uint {
	return p.taskId
}
