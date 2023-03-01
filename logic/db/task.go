package db

const (
	TaskStatusCreated   = "created"
	TaskStatusPending   = "pending"
	TaskStatusCompleted = "completed"
	TaskStatusFailed    = "failed"
)

type Task struct {
	CommonModel

	ParentTaskId uint
	TaskType     string `gorm:"size:50;not null"`
	TaskData     string `gorm:"type:LONGBLOB;not null"`
	TaskResult   string `gorm:"type:BLOB;not null"`
	Status       string `gorm:"size:50;index:idx_status;not null"`
}

func QueryAllUncompletedTask() (tasks []*Task, err error) {
	err = db.
		Where("status in ?", []string{TaskStatusCreated, TaskStatusPending}).
		Order("id asc").
		Find(&tasks).
		Error
	return
}

func QueryTaskById(id uint) (task *Task, err error) {
	err = db.Where("id = ?", id).Find(&task).Error
	return
}

func UpdateTaskStatus(task *Task, oriStatus string, newStatus string, option ...Options) error {
	task.Status = newStatus

	tx := useOptions(option...).Model(&Task{}).Where("id = ? and status = ?", task.ID, oriStatus).Updates(task)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}

func SaveTask(task *Task, option ...Options) error {
	return useOptions(option...).Save(task).Error
}
