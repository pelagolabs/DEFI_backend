package tasks

import (
	"fmt"
	"go.uber.org/zap"
	"veric-backend/internal/log"
	"veric-backend/internal/util"
	"veric-backend/logic/db"
)

var DefaultManage = NewManage(1)

func init() {
	log.GetLogger().Info("init task manage...")
}

type task interface {
	Run(*taskOperator) error
	TaskId() uint
}

type onFail interface {
	OnFail(*taskOperator, error) error
}

type taskCreate func(taskId uint, data string) (task, error)
type TaskType string

type TaskData interface {
	data() (string, error)
}

type Manage struct {
	taskTypes   util.SyncedMap[TaskType, taskCreate]
	parallelNum uint
	taskChan    chan task
}

func NewManage(parallelNum uint) *Manage {
	m := &Manage{
		parallelNum: parallelNum,
		taskChan:    make(chan task, 1000),
	}

	return m
}

func (m *Manage) register(taskType TaskType, create taskCreate) {
	m.taskTypes.Store(taskType, create)
}

func (m *Manage) NewTask(taskType TaskType, data TaskData) error {
	taskInstance, err := m.createTaskInstance(taskType, data, 0)
	if err != nil {
		return err
	}

	m.taskChan <- taskInstance
	return nil
}

func (m *Manage) NewTaskWithTx(taskType TaskType, data TaskData, f func(tx db.Options) error) error {
	var taskInstance task
	err := db.WithTx(func(tx db.Options) (err error) {
		taskInstance, err = m.createTaskInstance(taskType, data, 0, tx)
		if err != nil {
			return err
		}

		return f(tx)
	})
	if err != nil {
		return err
	}

	m.taskChan <- taskInstance
	return nil
}

func (m *Manage) createTaskInstance(taskType TaskType, data TaskData, parentTaskId uint, options ...db.Options) (task, error) {
	taskData, err := data.data()
	if err != nil {
		return nil, err
	}

	dbTask := &db.Task{
		ParentTaskId: parentTaskId,
		TaskType:     string(taskType),
		TaskData:     taskData,
		Status:       db.TaskStatusCreated,
	}

	err = db.SaveTask(dbTask, options...)
	if err != nil {
		return nil, err
	}

	var taskInstance task
	if taskCreateInstance, ok := m.taskTypes.Load(taskType); !ok {
		return nil, fmt.Errorf("task type %s not found", taskType)
	} else {
		taskInstance, err = taskCreateInstance(dbTask.ID, taskData)
		if err != nil {
			dbTask.TaskResult = "[init fail]" + err.Error()
			_ = db.UpdateTaskStatus(dbTask, dbTask.Status, db.TaskStatusFailed, options...)
			return nil, err
		}
	}

	return taskInstance, nil
}

func (m *Manage) Start(runUncompletedTask bool) {
	log.GetLogger().Info("start task manage...")

	for i := 0; i < int(m.parallelNum); i++ {
		go m.taskProcess()
	}

	if runUncompletedTask {
		tasks, err := db.QueryAllUncompletedTask()
		if err != nil {
			log.GetLogger().Warn("[task manage] init fail, db query error", zap.Error(err))
			return
		}

		for _, dbTask := range tasks {
			if taskCreateInstance, ok := m.taskTypes.Load(TaskType(dbTask.TaskType)); !ok {
				log.GetLogger().Warn("[task manage] init fail, dbTask type not found", zap.String("taskType", dbTask.TaskType))
				_ = db.UpdateTaskStatus(dbTask, dbTask.Status, db.TaskStatusFailed)
				continue
			} else {
				instance, err := taskCreateInstance(dbTask.ID, dbTask.TaskData)
				if err != nil {
					dbTask.TaskResult = "[init fail]" + err.Error()
					_ = db.UpdateTaskStatus(dbTask, dbTask.Status, db.TaskStatusFailed)
					continue
				}

				m.taskChan <- instance
			}
		}
	}
}

func (m *Manage) taskProcess() {
	waitProcessTask := make([]task, 0)
	for {
		var taskInstance task
		if len(waitProcessTask) > 0 {
			taskInstance = waitProcessTask[0]
			waitProcessTask = waitProcessTask[1:]
		} else {
			taskInstance = <-m.taskChan
		}

		_ = db.WithTx(func(tx db.Options) error {
			dbTask, dbErr := db.QueryTaskById(taskInstance.TaskId())
			if dbErr != nil {
				return dbErr
			}

			if dbTask.Status == db.TaskStatusFailed || dbTask.Status == db.TaskStatusCompleted {
				return nil
			}

			dbErr = db.UpdateTaskStatus(dbTask, dbTask.Status, db.TaskStatusPending)
			if dbErr != nil {
				return dbErr
			}

			op := &taskOperator{
				taskId: taskInstance.TaskId(),
				tx:     tx,
				manage: m,
			}

			dbErr = db.SavePoint("task_init", tx)
			if dbErr != nil {
				log.GetLogger().Warn(
					"[task manage] start tx fail",
					zap.Uint("taskId", taskInstance.TaskId()),
					zap.NamedError("dbErr", dbErr),
				)
				return nil
			}

			err := taskInstance.Run(op)
			if err != nil {
				if fail, ok := taskInstance.(onFail); ok {
					_ = fail.OnFail(op, err)
				}

				dbTask.TaskResult = err.Error()
				dbErr = db.UpdateTaskStatus(dbTask, dbTask.Status, db.TaskStatusFailed)
			} else {
				dbErr = db.UpdateTaskStatus(dbTask, dbTask.Status, db.TaskStatusCompleted)

				if len(op.subTasks) > 0 {
					waitProcessTask = append(waitProcessTask, op.subTasks...)
				}
			}

			if dbErr != nil {
				log.GetLogger().Warn(
					"[task manage] update task status fail",
					zap.Uint("taskId", taskInstance.TaskId()),
					zap.NamedError("dbErr", dbErr),
					zap.NamedError("taskErr", err),
				)
				return nil
			}

			return nil
		})
	}
}
