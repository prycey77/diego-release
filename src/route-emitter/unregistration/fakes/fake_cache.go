// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/diego-release/route-emitter/routingtable"
	"code.cloudfoundry.org/diego-release/route-emitter/unregistration"
)

type FakeCache struct {
	AddStub        func([]routingtable.RegistryMessage) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 []routingtable.RegistryMessage
	}
	addReturns struct {
		result1 error
	}
	addReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func() []*unregistration.Message
	listMutex       sync.RWMutex
	listArgsForCall []struct {
	}
	listReturns struct {
		result1 []*unregistration.Message
	}
	listReturnsOnCall map[int]struct {
		result1 []*unregistration.Message
	}
	RemoveStub        func([]routingtable.RegistryMessage) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 []routingtable.RegistryMessage
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCache) Add(arg1 []routingtable.RegistryMessage) error {
	var arg1Copy []routingtable.RegistryMessage
	if arg1 != nil {
		arg1Copy = make([]routingtable.RegistryMessage, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 []routingtable.RegistryMessage
	}{arg1Copy})
	fake.recordInvocation("Add", []interface{}{arg1Copy})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addReturns
	return fakeReturns.result1
}

func (fake *FakeCache) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeCache) AddCalls(stub func([]routingtable.RegistryMessage) error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeCache) AddArgsForCall(i int) []routingtable.RegistryMessage {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCache) AddReturns(result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCache) AddReturnsOnCall(i int, result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCache) List() []*unregistration.Message {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
	}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1
}

func (fake *FakeCache) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeCache) ListCalls(stub func() []*unregistration.Message) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeCache) ListReturns(result1 []*unregistration.Message) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []*unregistration.Message
	}{result1}
}

func (fake *FakeCache) ListReturnsOnCall(i int, result1 []*unregistration.Message) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []*unregistration.Message
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []*unregistration.Message
	}{result1}
}

func (fake *FakeCache) Remove(arg1 []routingtable.RegistryMessage) error {
	var arg1Copy []routingtable.RegistryMessage
	if arg1 != nil {
		arg1Copy = make([]routingtable.RegistryMessage, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 []routingtable.RegistryMessage
	}{arg1Copy})
	fake.recordInvocation("Remove", []interface{}{arg1Copy})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeReturns
	return fakeReturns.result1
}

func (fake *FakeCache) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeCache) RemoveCalls(stub func([]routingtable.RegistryMessage) error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *FakeCache) RemoveArgsForCall(i int) []routingtable.RegistryMessage {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCache) RemoveReturns(result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCache) RemoveReturnsOnCall(i int, result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCache) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCache) recordInvocation(key string, args []interface{}) {
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

var _ unregistration.Cache = new(FakeCache)
