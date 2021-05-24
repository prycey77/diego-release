// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/consuladapter"
)

type FakeLock struct {
	LockStub        func(stopCh <-chan struct{}) (lostLock <-chan struct{}, err error)
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		stopCh <-chan struct{}
	}
	lockReturns struct {
		result1 <-chan struct{}
		result2 error
	}
}

func (fake *FakeLock) Lock(stopCh <-chan struct{}) (lostLock <-chan struct{}, err error) {
	fake.lockMutex.Lock()
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		stopCh <-chan struct{}
	}{stopCh})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		return fake.LockStub(stopCh)
	} else {
		return fake.lockReturns.result1, fake.lockReturns.result2
	}
}

func (fake *FakeLock) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeLock) LockArgsForCall(i int) <-chan struct{} {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return fake.lockArgsForCall[i].stopCh
}

func (fake *FakeLock) LockReturns(result1 <-chan struct{}, result2 error) {
	fake.LockStub = nil
	fake.lockReturns = struct {
		result1 <-chan struct{}
		result2 error
	}{result1, result2}
}

var _ consuladapter.Lock = new(FakeLock)
