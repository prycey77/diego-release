// Code generated by counterfeiter. DO NOT EDIT.
package taskworkpoolfakes

import (
	"sync"

	"code.cloudfoundry.org/diego-release/bbs/db"
	"code.cloudfoundry.org/diego-release/bbs/events"
	"code.cloudfoundry.org/diego-release/bbs/models"
	"code.cloudfoundry.org/diego-release/bbs/taskworkpool"
)

type FakeTaskCompletionClient struct {
	SubmitStub        func(db.TaskDB, events.Hub, *models.Task)
	submitMutex       sync.RWMutex
	submitArgsForCall []struct {
		arg1 db.TaskDB
		arg2 events.Hub
		arg3 *models.Task
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskCompletionClient) Submit(arg1 db.TaskDB, arg2 events.Hub, arg3 *models.Task) {
	fake.submitMutex.Lock()
	fake.submitArgsForCall = append(fake.submitArgsForCall, struct {
		arg1 db.TaskDB
		arg2 events.Hub
		arg3 *models.Task
	}{arg1, arg2, arg3})
	stub := fake.SubmitStub
	fake.recordInvocation("Submit", []interface{}{arg1, arg2, arg3})
	fake.submitMutex.Unlock()
	if stub != nil {
		fake.SubmitStub(arg1, arg2, arg3)
	}
}

func (fake *FakeTaskCompletionClient) SubmitCallCount() int {
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	return len(fake.submitArgsForCall)
}

func (fake *FakeTaskCompletionClient) SubmitCalls(stub func(db.TaskDB, events.Hub, *models.Task)) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = stub
}

func (fake *FakeTaskCompletionClient) SubmitArgsForCall(i int) (db.TaskDB, events.Hub, *models.Task) {
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	argsForCall := fake.submitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTaskCompletionClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskCompletionClient) recordInvocation(key string, args []interface{}) {
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

var _ taskworkpool.TaskCompletionClient = new(FakeTaskCompletionClient)
