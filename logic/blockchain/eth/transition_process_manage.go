package eth

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"reflect"
	"time"
	"veric-backend/internal/util"
)

type ProcessItem interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type ProcessItemSaver interface {
	SaveProcessItem(taskId string, item []byte) error
	LoadProcessItem(taskId string) (item []byte, err error)
	DeleteProcessItem(taskId string) error
}

type TransitionProcessStep[T ProcessItem] struct {
	stepName    string
	isFirstStep bool
	isLastStep  bool
	processFunc func(T, *types.Transaction, *types.Receipt, *bind.TransactOpts) (next *types.Transaction, err error)
}

type syncWaitItem[T ProcessItem] struct {
	C   chan T
	err error
}

type TransitionProcessManage[T ProcessItem] struct {
	tm          *TransitionManage
	saver       ProcessItemSaver
	processName string
	steps       []TransitionProcessStep[T]
	syncWait    util.SyncedMap[string, *syncWaitItem[T]]
}

type TransitionProcessStart[T ProcessItem] interface {
	FirstStep(stepName string, f func(T, *bind.TransactOpts) (next *types.Transaction, err error)) TransitionProcessMaker[T]
	OneStep(stepName string, f func(T) (err error)) (TransitionProcessCaller[T], error)
}

type TransitionProcessMaker[T ProcessItem] interface {
	NextStep(stepName string, f func(T, *types.Transaction, *types.Receipt, *bind.TransactOpts) (next *types.Transaction, err error)) TransitionProcessMaker[T]
	LastStep(stepName string, f func(T, *types.Transaction, *types.Receipt) (err error)) (TransitionProcessCaller[T], error)
}

type TransitionProcessCaller[T ProcessItem] interface {
	RunAsync(ctx T) error
	RunSyncWithTimeout(timeout time.Duration, ctx T) (err error)
	RunSyncResultWithTimeout(timeout time.Duration, ctx T) (result T, err error)
}

func NewTransitionProcessManage[T ProcessItem](tm *TransitionManage, saver ProcessItemSaver, processName string) TransitionProcessStart[T] {
	return &TransitionProcessManage[T]{
		tm:          tm,
		processName: processName,
		saver:       saver,
	}
}

func (m *TransitionProcessManage[T]) makeFullStepName(stepName string) string {
	return fmt.Sprintf("%s::%s", m.processName, stepName)
}

func (m *TransitionProcessManage[T]) addStep(stepName string, f func(T, *types.Transaction, *types.Receipt, *bind.TransactOpts) (next *types.Transaction, err error)) *TransitionProcessManage[T] {
	m.steps = append(m.steps, TransitionProcessStep[T]{
		stepName:    stepName,
		processFunc: f,
	})
	return m
}

func (m *TransitionProcessManage[T]) addLastStep(stepName string, f func(T, *types.Transaction, *types.Receipt) (err error)) (*TransitionProcessManage[T], error) {
	m.addStep(stepName, func(t T, tx *types.Transaction, r *types.Receipt, opts *bind.TransactOpts) (next *types.Transaction, err error) {
		return nil, f(t, tx, r)
	})

	if len(m.steps) > 0 {
		m.steps[0].isFirstStep = true
		m.steps[len(m.steps)-1].isLastStep = true

		for i, step := range m.steps {
			err := m.tm.TransitionRegister(m.makeFullStepName(step.stepName), m.makeNextStepFunc(i))
			if err != nil {
				return nil, err
			}
		}
	}

	return m, nil
}

func (m *TransitionProcessManage[T]) FirstStep(stepName string, f func(T, *bind.TransactOpts) (next *types.Transaction, err error)) TransitionProcessMaker[T] {
	return m.addStep(stepName, func(t T, _ *types.Transaction, _ *types.Receipt, opts *bind.TransactOpts) (next *types.Transaction, err error) {
		return f(t, opts)
	})
}

func (m *TransitionProcessManage[T]) OneStep(stepName string, f func(T) (err error)) (TransitionProcessCaller[T], error) {
	return m.addLastStep(stepName, func(t T, _ *types.Transaction, _ *types.Receipt) (err error) {
		return f(t)
	})
}

func (m *TransitionProcessManage[T]) NextStep(stepName string, f func(T, *types.Transaction, *types.Receipt, *bind.TransactOpts) (next *types.Transaction, err error)) TransitionProcessMaker[T] {
	return m.addStep(stepName, f)
}

func (m *TransitionProcessManage[T]) LastStep(stepName string, f func(T, *types.Transaction, *types.Receipt) (err error)) (TransitionProcessCaller[T], error) {
	return m.addLastStep(stepName, f)
}

func (m *TransitionProcessManage[T]) commitAndRunSync(taskId string, ctx T) error {
	ctxBinary, err := ctx.MarshalBinary()
	if err != nil {
		return err
	}

	err = m.saver.SaveProcessItem(taskId, ctxBinary)
	if err != nil {
		return err
	}

	if len(m.steps) > 0 {
		firstStep := m.steps[0]
		return m.tm.TransitionRequest(m.makeFullStepName(firstStep.stepName), taskId, func(taskId string, opts *bind.TransactOpts) (tx *types.Transaction, err error) {
			err = m.useCtx(taskId, func(ctx T) error {
				tx, err = firstStep.processFunc(ctx, nil, nil, opts)
				return err
			})

			return tx, err
		})
	}

	return nil
}

func (m *TransitionProcessManage[T]) RunAsync(ctx T) error {
	taskId := uuid.NewString()
	return m.commitAndRunSync(taskId, ctx)
}

func (m *TransitionProcessManage[T]) RunSyncWithTimeout(timeout time.Duration, ctx T) (err error) {
	_, err = m.RunSyncResultWithTimeout(timeout, ctx)
	return err
}

func (m *TransitionProcessManage[T]) RunSyncResultWithTimeout(timeout time.Duration, ctx T) (result T, err error) {
	taskId := uuid.NewString()
	waitChan := make(chan T, 1)
	defer func() {
		m.syncWait.Delete(taskId)
		close(waitChan)
	}()

	item := &syncWaitItem[T]{C: waitChan}
	if _, loaded := m.syncWait.LoadOrStore(taskId, item); !loaded {
		err = m.commitAndRunSync(taskId, ctx)
		if err != nil {
			return result, err
		}
	} else {
		return result, errors.New("store task error, try again later")
	}

	select {
	case resultCtx := <-waitChan:
		return resultCtx, item.err
	case <-time.After(timeout):
		return result, errors.New("timeout")
	}
}

func (m *TransitionProcessManage[T]) useCtx(taskId string, fn func(ctx T) error) (err error) {
	oriByte, err := m.saver.LoadProcessItem(taskId)
	if err != nil {
		return err
	}

	ctx := reflect.New(reflect.TypeOf(new(T)).Elem().Elem()).Interface().(T)
	err = ctx.UnmarshalBinary(oriByte)

	err = fn(ctx)
	if err != nil {
		return err
	}

	ctxBinary, err := ctx.MarshalBinary()
	if err != nil {
		return err
	}

	if bytes.Compare(oriByte, ctxBinary) != 0 {
		err = m.saver.SaveProcessItem(taskId, ctxBinary)
	}

	return err
}

func (m *TransitionProcessManage[T]) fireWait(taskId string, result T, err error) {
	if item, ok := m.syncWait.Load(taskId); ok {
		item.err = err
		select {
		case item.C <- result:
		default:
		}
	}
}

func (m *TransitionProcessManage[T]) makeNextStepFunc(i int) TransitionEvent {
	return func(taskId string, tx *types.Transaction, receipt *types.Receipt) error {
		return m.useCtx(taskId, func(ctx T) (err error) {
			if receipt.Status == 0 {
				m.fireWait(taskId, ctx, errors.New("receipt status is 0"))
				return nil
			}

			nextStep := i + 1
			if len(m.steps) > nextStep {
				step := m.steps[nextStep]
				if step.isLastStep {
					_, err = step.processFunc(ctx, tx, receipt, nil)
					m.fireWait(taskId, ctx, err)
					if err != nil {
						return err
					}

					_ = m.saver.DeleteProcessItem(taskId)
				} else {
					err = m.tm.TransitionRequest(m.makeFullStepName(step.stepName), taskId, func(taskId string, opts *bind.TransactOpts) (tx *types.Transaction, err error) {
						return step.processFunc(ctx, tx, receipt, opts)
					})
					if err != nil {
						m.fireWait(taskId, ctx, err)
						return err
					}
				}
			}
			return nil
		})
	}
}
