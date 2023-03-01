package db

type ProcessItem struct {
	CommonModel

	TaskId string `gorm:"size:64;index:idx_taskid;not null"`
	Data   []byte `gorm:"type:BLOB;not null"`
}

func FindProcessItemByTaskId(taskId string) (tx *ProcessItem, err error) {
	err = db.Where("task_id = ?", taskId).Find(&tx).Error
	return
}

func SaveProcessItem(tx *ProcessItem) (err error) {
	return db.Save(tx).Error
}

func DeleteProcessItem(tx *ProcessItem) (err error) {
	return db.Delete(tx).Error
}
