package tasks

import (
	"veric-backend/logic/db"
)

type taskOperator struct {
	taskId   uint
	tx       db.Options
	subTasks []task
	manage   *Manage
}

func (t *taskOperator) NewSubTask(taskType TaskType, data TaskData) error {
	taskInstance, err := t.manage.createTaskInstance(taskType, data, t.taskId)
	if err != nil {
		return err
	}

	t.subTasks = append(t.subTasks, taskInstance)
	return nil
}

func (t *taskOperator) RollbackTaskDB() error {
	return db.RollbackTo("task_init", t.tx)
}
