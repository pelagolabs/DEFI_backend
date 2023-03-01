package common

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

type ContractSaver struct {
	Chain string
}

func (c *ContractSaver) Save(name string, taskId string, txHash common.Hash, txBinary []byte) error {
	nameSet := strings.Split(name, "::")

	return db.SaveTx(&db.Tx{
		Name:      fmt.Sprintf("[%s]%s", c.Chain, name),
		Method:    nameSet[0],
		TaskId:    taskId,
		Hash:      txHash.String(),
		Status:    db.TxStatusUnknown,
		ChainName: c.Chain,
		TxData:    txBinary,
	})
}

func (c *ContractSaver) SaveTxSuccess(taskId string, txHash common.Hash, gasUsed *types.BigInt, receiptBinary []byte) error {
	tx, err := db.FindTxByHashAndTaskId(txHash.String(), taskId)
	if err != nil {
		return err
	}

	tx.Status = db.TxStatusSuccess
	tx.ReceiptData = receiptBinary
	tx.GasUsed = gasUsed
	return db.SaveTx(tx)
}

func (c *ContractSaver) SaveTxFail(taskId string, txHash common.Hash, gasUsed *types.BigInt, receiptBinary []byte, why error) error {
	tx, err := db.FindTxByHashAndTaskId(txHash.String(), taskId)
	if err != nil {
		return err
	}

	errStr := why.Error()
	tx.Status = db.TxStatusFail
	tx.Why = &errStr
	tx.ReceiptData = receiptBinary
	tx.GasUsed = gasUsed
	return db.SaveTx(tx)
}

func (c *ContractSaver) LoadUnknownTx(name string) (unknownTx <-chan *eth.TransitionLoadItem, err error) {
	status, err := db.FindTxByNameAndStatus(fmt.Sprintf("[%s]%s", c.Chain, name), db.TxStatusUnknown)
	if err != nil {
		return nil, err
	}

	txChan := make(chan *eth.TransitionLoadItem, 100)
	go func() {
		for _, tx := range status {
			txChan <- &eth.TransitionLoadItem{
				TaskId: tx.TaskId,
				Tx:     tx.TxData,
			}
		}
		close(txChan)
	}()

	return txChan, nil
}

func (c *ContractSaver) SaveProcessItem(taskId string, item []byte) error {
	return db.SaveProcessItem(&db.ProcessItem{
		TaskId: taskId,
		Data:   item,
	})
}

func (c *ContractSaver) LoadProcessItem(taskId string) (item []byte, err error) {
	dbItem, err := db.FindProcessItemByTaskId(taskId)
	if err != nil {
		return nil, err
	}

	return dbItem.Data, nil
}

func (c *ContractSaver) DeleteProcessItem(taskId string) error {
	return nil
}
