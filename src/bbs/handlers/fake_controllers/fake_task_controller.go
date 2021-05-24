// Code generated by counterfeiter. DO NOT EDIT.
package fake_controllers

import (
	"context"
	"sync"
	"time"

	"code.cloudfoundry.org/diego-release/bbs/handlers"
	"code.cloudfoundry.org/diego-release/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeTaskController struct {
	CancelTaskStub        func(context.Context, lager.Logger, string) error
	cancelTaskMutex       sync.RWMutex
	cancelTaskArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	cancelTaskReturns struct {
		result1 error
	}
	cancelTaskReturnsOnCall map[int]struct {
		result1 error
	}
	CompleteTaskStub        func(context.Context, lager.Logger, string, string, bool, string, string) error
	completeTaskMutex       sync.RWMutex
	completeTaskArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
		arg5 bool
		arg6 string
		arg7 string
	}
	completeTaskReturns struct {
		result1 error
	}
	completeTaskReturnsOnCall map[int]struct {
		result1 error
	}
	ConvergeTasksStub        func(context.Context, lager.Logger, time.Duration, time.Duration, time.Duration) error
	convergeTasksMutex       sync.RWMutex
	convergeTasksArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 time.Duration
		arg4 time.Duration
		arg5 time.Duration
	}
	convergeTasksReturns struct {
		result1 error
	}
	convergeTasksReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteTaskStub        func(context.Context, lager.Logger, string) error
	deleteTaskMutex       sync.RWMutex
	deleteTaskArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	deleteTaskReturns struct {
		result1 error
	}
	deleteTaskReturnsOnCall map[int]struct {
		result1 error
	}
	DesireTaskStub        func(context.Context, lager.Logger, *models.TaskDefinition, string, string) error
	desireTaskMutex       sync.RWMutex
	desireTaskArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.TaskDefinition
		arg4 string
		arg5 string
	}
	desireTaskReturns struct {
		result1 error
	}
	desireTaskReturnsOnCall map[int]struct {
		result1 error
	}
	FailTaskStub        func(context.Context, lager.Logger, string, string) error
	failTaskMutex       sync.RWMutex
	failTaskArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
	}
	failTaskReturns struct {
		result1 error
	}
	failTaskReturnsOnCall map[int]struct {
		result1 error
	}
	RejectTaskStub        func(context.Context, lager.Logger, string, string) error
	rejectTaskMutex       sync.RWMutex
	rejectTaskArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
	}
	rejectTaskReturns struct {
		result1 error
	}
	rejectTaskReturnsOnCall map[int]struct {
		result1 error
	}
	ResolvingTaskStub        func(context.Context, lager.Logger, string) error
	resolvingTaskMutex       sync.RWMutex
	resolvingTaskArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	resolvingTaskReturns struct {
		result1 error
	}
	resolvingTaskReturnsOnCall map[int]struct {
		result1 error
	}
	StartTaskStub        func(context.Context, lager.Logger, string, string) (bool, error)
	startTaskMutex       sync.RWMutex
	startTaskArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
	}
	startTaskReturns struct {
		result1 bool
		result2 error
	}
	startTaskReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	TaskByGuidStub        func(context.Context, lager.Logger, string) (*models.Task, error)
	taskByGuidMutex       sync.RWMutex
	taskByGuidArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	taskByGuidReturns struct {
		result1 *models.Task
		result2 error
	}
	taskByGuidReturnsOnCall map[int]struct {
		result1 *models.Task
		result2 error
	}
	TasksStub        func(context.Context, lager.Logger, string, string) ([]*models.Task, error)
	tasksMutex       sync.RWMutex
	tasksArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
	}
	tasksReturns struct {
		result1 []*models.Task
		result2 error
	}
	tasksReturnsOnCall map[int]struct {
		result1 []*models.Task
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskController) CancelTask(arg1 context.Context, arg2 lager.Logger, arg3 string) error {
	fake.cancelTaskMutex.Lock()
	ret, specificReturn := fake.cancelTaskReturnsOnCall[len(fake.cancelTaskArgsForCall)]
	fake.cancelTaskArgsForCall = append(fake.cancelTaskArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CancelTaskStub
	fakeReturns := fake.cancelTaskReturns
	fake.recordInvocation("CancelTask", []interface{}{arg1, arg2, arg3})
	fake.cancelTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskController) CancelTaskCallCount() int {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return len(fake.cancelTaskArgsForCall)
}

func (fake *FakeTaskController) CancelTaskCalls(stub func(context.Context, lager.Logger, string) error) {
	fake.cancelTaskMutex.Lock()
	defer fake.cancelTaskMutex.Unlock()
	fake.CancelTaskStub = stub
}

func (fake *FakeTaskController) CancelTaskArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	argsForCall := fake.cancelTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTaskController) CancelTaskReturns(result1 error) {
	fake.cancelTaskMutex.Lock()
	defer fake.cancelTaskMutex.Unlock()
	fake.CancelTaskStub = nil
	fake.cancelTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) CancelTaskReturnsOnCall(i int, result1 error) {
	fake.cancelTaskMutex.Lock()
	defer fake.cancelTaskMutex.Unlock()
	fake.CancelTaskStub = nil
	if fake.cancelTaskReturnsOnCall == nil {
		fake.cancelTaskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cancelTaskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) CompleteTask(arg1 context.Context, arg2 lager.Logger, arg3 string, arg4 string, arg5 bool, arg6 string, arg7 string) error {
	fake.completeTaskMutex.Lock()
	ret, specificReturn := fake.completeTaskReturnsOnCall[len(fake.completeTaskArgsForCall)]
	fake.completeTaskArgsForCall = append(fake.completeTaskArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
		arg5 bool
		arg6 string
		arg7 string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	stub := fake.CompleteTaskStub
	fakeReturns := fake.completeTaskReturns
	fake.recordInvocation("CompleteTask", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.completeTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskController) CompleteTaskCallCount() int {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return len(fake.completeTaskArgsForCall)
}

func (fake *FakeTaskController) CompleteTaskCalls(stub func(context.Context, lager.Logger, string, string, bool, string, string) error) {
	fake.completeTaskMutex.Lock()
	defer fake.completeTaskMutex.Unlock()
	fake.CompleteTaskStub = stub
}

func (fake *FakeTaskController) CompleteTaskArgsForCall(i int) (context.Context, lager.Logger, string, string, bool, string, string) {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	argsForCall := fake.completeTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeTaskController) CompleteTaskReturns(result1 error) {
	fake.completeTaskMutex.Lock()
	defer fake.completeTaskMutex.Unlock()
	fake.CompleteTaskStub = nil
	fake.completeTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) CompleteTaskReturnsOnCall(i int, result1 error) {
	fake.completeTaskMutex.Lock()
	defer fake.completeTaskMutex.Unlock()
	fake.CompleteTaskStub = nil
	if fake.completeTaskReturnsOnCall == nil {
		fake.completeTaskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.completeTaskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) ConvergeTasks(arg1 context.Context, arg2 lager.Logger, arg3 time.Duration, arg4 time.Duration, arg5 time.Duration) error {
	fake.convergeTasksMutex.Lock()
	ret, specificReturn := fake.convergeTasksReturnsOnCall[len(fake.convergeTasksArgsForCall)]
	fake.convergeTasksArgsForCall = append(fake.convergeTasksArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 time.Duration
		arg4 time.Duration
		arg5 time.Duration
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.ConvergeTasksStub
	fakeReturns := fake.convergeTasksReturns
	fake.recordInvocation("ConvergeTasks", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.convergeTasksMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskController) ConvergeTasksCallCount() int {
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	return len(fake.convergeTasksArgsForCall)
}

func (fake *FakeTaskController) ConvergeTasksCalls(stub func(context.Context, lager.Logger, time.Duration, time.Duration, time.Duration) error) {
	fake.convergeTasksMutex.Lock()
	defer fake.convergeTasksMutex.Unlock()
	fake.ConvergeTasksStub = stub
}

func (fake *FakeTaskController) ConvergeTasksArgsForCall(i int) (context.Context, lager.Logger, time.Duration, time.Duration, time.Duration) {
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	argsForCall := fake.convergeTasksArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTaskController) ConvergeTasksReturns(result1 error) {
	fake.convergeTasksMutex.Lock()
	defer fake.convergeTasksMutex.Unlock()
	fake.ConvergeTasksStub = nil
	fake.convergeTasksReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) ConvergeTasksReturnsOnCall(i int, result1 error) {
	fake.convergeTasksMutex.Lock()
	defer fake.convergeTasksMutex.Unlock()
	fake.ConvergeTasksStub = nil
	if fake.convergeTasksReturnsOnCall == nil {
		fake.convergeTasksReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.convergeTasksReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) DeleteTask(arg1 context.Context, arg2 lager.Logger, arg3 string) error {
	fake.deleteTaskMutex.Lock()
	ret, specificReturn := fake.deleteTaskReturnsOnCall[len(fake.deleteTaskArgsForCall)]
	fake.deleteTaskArgsForCall = append(fake.deleteTaskArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.DeleteTaskStub
	fakeReturns := fake.deleteTaskReturns
	fake.recordInvocation("DeleteTask", []interface{}{arg1, arg2, arg3})
	fake.deleteTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskController) DeleteTaskCallCount() int {
	fake.deleteTaskMutex.RLock()
	defer fake.deleteTaskMutex.RUnlock()
	return len(fake.deleteTaskArgsForCall)
}

func (fake *FakeTaskController) DeleteTaskCalls(stub func(context.Context, lager.Logger, string) error) {
	fake.deleteTaskMutex.Lock()
	defer fake.deleteTaskMutex.Unlock()
	fake.DeleteTaskStub = stub
}

func (fake *FakeTaskController) DeleteTaskArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.deleteTaskMutex.RLock()
	defer fake.deleteTaskMutex.RUnlock()
	argsForCall := fake.deleteTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTaskController) DeleteTaskReturns(result1 error) {
	fake.deleteTaskMutex.Lock()
	defer fake.deleteTaskMutex.Unlock()
	fake.DeleteTaskStub = nil
	fake.deleteTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) DeleteTaskReturnsOnCall(i int, result1 error) {
	fake.deleteTaskMutex.Lock()
	defer fake.deleteTaskMutex.Unlock()
	fake.DeleteTaskStub = nil
	if fake.deleteTaskReturnsOnCall == nil {
		fake.deleteTaskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteTaskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) DesireTask(arg1 context.Context, arg2 lager.Logger, arg3 *models.TaskDefinition, arg4 string, arg5 string) error {
	fake.desireTaskMutex.Lock()
	ret, specificReturn := fake.desireTaskReturnsOnCall[len(fake.desireTaskArgsForCall)]
	fake.desireTaskArgsForCall = append(fake.desireTaskArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.TaskDefinition
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.DesireTaskStub
	fakeReturns := fake.desireTaskReturns
	fake.recordInvocation("DesireTask", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.desireTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskController) DesireTaskCallCount() int {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return len(fake.desireTaskArgsForCall)
}

func (fake *FakeTaskController) DesireTaskCalls(stub func(context.Context, lager.Logger, *models.TaskDefinition, string, string) error) {
	fake.desireTaskMutex.Lock()
	defer fake.desireTaskMutex.Unlock()
	fake.DesireTaskStub = stub
}

func (fake *FakeTaskController) DesireTaskArgsForCall(i int) (context.Context, lager.Logger, *models.TaskDefinition, string, string) {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	argsForCall := fake.desireTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTaskController) DesireTaskReturns(result1 error) {
	fake.desireTaskMutex.Lock()
	defer fake.desireTaskMutex.Unlock()
	fake.DesireTaskStub = nil
	fake.desireTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) DesireTaskReturnsOnCall(i int, result1 error) {
	fake.desireTaskMutex.Lock()
	defer fake.desireTaskMutex.Unlock()
	fake.DesireTaskStub = nil
	if fake.desireTaskReturnsOnCall == nil {
		fake.desireTaskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.desireTaskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) FailTask(arg1 context.Context, arg2 lager.Logger, arg3 string, arg4 string) error {
	fake.failTaskMutex.Lock()
	ret, specificReturn := fake.failTaskReturnsOnCall[len(fake.failTaskArgsForCall)]
	fake.failTaskArgsForCall = append(fake.failTaskArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.FailTaskStub
	fakeReturns := fake.failTaskReturns
	fake.recordInvocation("FailTask", []interface{}{arg1, arg2, arg3, arg4})
	fake.failTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskController) FailTaskCallCount() int {
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	return len(fake.failTaskArgsForCall)
}

func (fake *FakeTaskController) FailTaskCalls(stub func(context.Context, lager.Logger, string, string) error) {
	fake.failTaskMutex.Lock()
	defer fake.failTaskMutex.Unlock()
	fake.FailTaskStub = stub
}

func (fake *FakeTaskController) FailTaskArgsForCall(i int) (context.Context, lager.Logger, string, string) {
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	argsForCall := fake.failTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTaskController) FailTaskReturns(result1 error) {
	fake.failTaskMutex.Lock()
	defer fake.failTaskMutex.Unlock()
	fake.FailTaskStub = nil
	fake.failTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) FailTaskReturnsOnCall(i int, result1 error) {
	fake.failTaskMutex.Lock()
	defer fake.failTaskMutex.Unlock()
	fake.FailTaskStub = nil
	if fake.failTaskReturnsOnCall == nil {
		fake.failTaskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.failTaskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) RejectTask(arg1 context.Context, arg2 lager.Logger, arg3 string, arg4 string) error {
	fake.rejectTaskMutex.Lock()
	ret, specificReturn := fake.rejectTaskReturnsOnCall[len(fake.rejectTaskArgsForCall)]
	fake.rejectTaskArgsForCall = append(fake.rejectTaskArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.RejectTaskStub
	fakeReturns := fake.rejectTaskReturns
	fake.recordInvocation("RejectTask", []interface{}{arg1, arg2, arg3, arg4})
	fake.rejectTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskController) RejectTaskCallCount() int {
	fake.rejectTaskMutex.RLock()
	defer fake.rejectTaskMutex.RUnlock()
	return len(fake.rejectTaskArgsForCall)
}

func (fake *FakeTaskController) RejectTaskCalls(stub func(context.Context, lager.Logger, string, string) error) {
	fake.rejectTaskMutex.Lock()
	defer fake.rejectTaskMutex.Unlock()
	fake.RejectTaskStub = stub
}

func (fake *FakeTaskController) RejectTaskArgsForCall(i int) (context.Context, lager.Logger, string, string) {
	fake.rejectTaskMutex.RLock()
	defer fake.rejectTaskMutex.RUnlock()
	argsForCall := fake.rejectTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTaskController) RejectTaskReturns(result1 error) {
	fake.rejectTaskMutex.Lock()
	defer fake.rejectTaskMutex.Unlock()
	fake.RejectTaskStub = nil
	fake.rejectTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) RejectTaskReturnsOnCall(i int, result1 error) {
	fake.rejectTaskMutex.Lock()
	defer fake.rejectTaskMutex.Unlock()
	fake.RejectTaskStub = nil
	if fake.rejectTaskReturnsOnCall == nil {
		fake.rejectTaskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rejectTaskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) ResolvingTask(arg1 context.Context, arg2 lager.Logger, arg3 string) error {
	fake.resolvingTaskMutex.Lock()
	ret, specificReturn := fake.resolvingTaskReturnsOnCall[len(fake.resolvingTaskArgsForCall)]
	fake.resolvingTaskArgsForCall = append(fake.resolvingTaskArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ResolvingTaskStub
	fakeReturns := fake.resolvingTaskReturns
	fake.recordInvocation("ResolvingTask", []interface{}{arg1, arg2, arg3})
	fake.resolvingTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskController) ResolvingTaskCallCount() int {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return len(fake.resolvingTaskArgsForCall)
}

func (fake *FakeTaskController) ResolvingTaskCalls(stub func(context.Context, lager.Logger, string) error) {
	fake.resolvingTaskMutex.Lock()
	defer fake.resolvingTaskMutex.Unlock()
	fake.ResolvingTaskStub = stub
}

func (fake *FakeTaskController) ResolvingTaskArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	argsForCall := fake.resolvingTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTaskController) ResolvingTaskReturns(result1 error) {
	fake.resolvingTaskMutex.Lock()
	defer fake.resolvingTaskMutex.Unlock()
	fake.ResolvingTaskStub = nil
	fake.resolvingTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) ResolvingTaskReturnsOnCall(i int, result1 error) {
	fake.resolvingTaskMutex.Lock()
	defer fake.resolvingTaskMutex.Unlock()
	fake.ResolvingTaskStub = nil
	if fake.resolvingTaskReturnsOnCall == nil {
		fake.resolvingTaskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.resolvingTaskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) StartTask(arg1 context.Context, arg2 lager.Logger, arg3 string, arg4 string) (bool, error) {
	fake.startTaskMutex.Lock()
	ret, specificReturn := fake.startTaskReturnsOnCall[len(fake.startTaskArgsForCall)]
	fake.startTaskArgsForCall = append(fake.startTaskArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.StartTaskStub
	fakeReturns := fake.startTaskReturns
	fake.recordInvocation("StartTask", []interface{}{arg1, arg2, arg3, arg4})
	fake.startTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskController) StartTaskCallCount() int {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return len(fake.startTaskArgsForCall)
}

func (fake *FakeTaskController) StartTaskCalls(stub func(context.Context, lager.Logger, string, string) (bool, error)) {
	fake.startTaskMutex.Lock()
	defer fake.startTaskMutex.Unlock()
	fake.StartTaskStub = stub
}

func (fake *FakeTaskController) StartTaskArgsForCall(i int) (context.Context, lager.Logger, string, string) {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	argsForCall := fake.startTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTaskController) StartTaskReturns(result1 bool, result2 error) {
	fake.startTaskMutex.Lock()
	defer fake.startTaskMutex.Unlock()
	fake.StartTaskStub = nil
	fake.startTaskReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskController) StartTaskReturnsOnCall(i int, result1 bool, result2 error) {
	fake.startTaskMutex.Lock()
	defer fake.startTaskMutex.Unlock()
	fake.StartTaskStub = nil
	if fake.startTaskReturnsOnCall == nil {
		fake.startTaskReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.startTaskReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskController) TaskByGuid(arg1 context.Context, arg2 lager.Logger, arg3 string) (*models.Task, error) {
	fake.taskByGuidMutex.Lock()
	ret, specificReturn := fake.taskByGuidReturnsOnCall[len(fake.taskByGuidArgsForCall)]
	fake.taskByGuidArgsForCall = append(fake.taskByGuidArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.TaskByGuidStub
	fakeReturns := fake.taskByGuidReturns
	fake.recordInvocation("TaskByGuid", []interface{}{arg1, arg2, arg3})
	fake.taskByGuidMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskController) TaskByGuidCallCount() int {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return len(fake.taskByGuidArgsForCall)
}

func (fake *FakeTaskController) TaskByGuidCalls(stub func(context.Context, lager.Logger, string) (*models.Task, error)) {
	fake.taskByGuidMutex.Lock()
	defer fake.taskByGuidMutex.Unlock()
	fake.TaskByGuidStub = stub
}

func (fake *FakeTaskController) TaskByGuidArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	argsForCall := fake.taskByGuidArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTaskController) TaskByGuidReturns(result1 *models.Task, result2 error) {
	fake.taskByGuidMutex.Lock()
	defer fake.taskByGuidMutex.Unlock()
	fake.TaskByGuidStub = nil
	fake.taskByGuidReturns = struct {
		result1 *models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskController) TaskByGuidReturnsOnCall(i int, result1 *models.Task, result2 error) {
	fake.taskByGuidMutex.Lock()
	defer fake.taskByGuidMutex.Unlock()
	fake.TaskByGuidStub = nil
	if fake.taskByGuidReturnsOnCall == nil {
		fake.taskByGuidReturnsOnCall = make(map[int]struct {
			result1 *models.Task
			result2 error
		})
	}
	fake.taskByGuidReturnsOnCall[i] = struct {
		result1 *models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskController) Tasks(arg1 context.Context, arg2 lager.Logger, arg3 string, arg4 string) ([]*models.Task, error) {
	fake.tasksMutex.Lock()
	ret, specificReturn := fake.tasksReturnsOnCall[len(fake.tasksArgsForCall)]
	fake.tasksArgsForCall = append(fake.tasksArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.TasksStub
	fakeReturns := fake.tasksReturns
	fake.recordInvocation("Tasks", []interface{}{arg1, arg2, arg3, arg4})
	fake.tasksMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskController) TasksCallCount() int {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return len(fake.tasksArgsForCall)
}

func (fake *FakeTaskController) TasksCalls(stub func(context.Context, lager.Logger, string, string) ([]*models.Task, error)) {
	fake.tasksMutex.Lock()
	defer fake.tasksMutex.Unlock()
	fake.TasksStub = stub
}

func (fake *FakeTaskController) TasksArgsForCall(i int) (context.Context, lager.Logger, string, string) {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	argsForCall := fake.tasksArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTaskController) TasksReturns(result1 []*models.Task, result2 error) {
	fake.tasksMutex.Lock()
	defer fake.tasksMutex.Unlock()
	fake.TasksStub = nil
	fake.tasksReturns = struct {
		result1 []*models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskController) TasksReturnsOnCall(i int, result1 []*models.Task, result2 error) {
	fake.tasksMutex.Lock()
	defer fake.tasksMutex.Unlock()
	fake.TasksStub = nil
	if fake.tasksReturnsOnCall == nil {
		fake.tasksReturnsOnCall = make(map[int]struct {
			result1 []*models.Task
			result2 error
		})
	}
	fake.tasksReturnsOnCall[i] = struct {
		result1 []*models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	fake.deleteTaskMutex.RLock()
	defer fake.deleteTaskMutex.RUnlock()
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	fake.rejectTaskMutex.RLock()
	defer fake.rejectTaskMutex.RUnlock()
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskController) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ handlers.TaskController = new(FakeTaskController)
