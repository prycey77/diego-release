// Code generated by counterfeiter. DO NOT EDIT.
package eventfakes

import (
	"sync"

	"code.cloudfoundry.org/diego-release/bbs/events"
	"github.com/vito/go-sse/sse"
)

type FakeRawEventSource struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	NextStub        func() (sse.Event, error)
	nextMutex       sync.RWMutex
	nextArgsForCall []struct {
	}
	nextReturns struct {
		result1 sse.Event
		result2 error
	}
	nextReturnsOnCall map[int]struct {
		result1 sse.Event
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRawEventSource) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRawEventSource) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeRawEventSource) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeRawEventSource) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRawEventSource) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRawEventSource) Next() (sse.Event, error) {
	fake.nextMutex.Lock()
	ret, specificReturn := fake.nextReturnsOnCall[len(fake.nextArgsForCall)]
	fake.nextArgsForCall = append(fake.nextArgsForCall, struct {
	}{})
	stub := fake.NextStub
	fakeReturns := fake.nextReturns
	fake.recordInvocation("Next", []interface{}{})
	fake.nextMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRawEventSource) NextCallCount() int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	return len(fake.nextArgsForCall)
}

func (fake *FakeRawEventSource) NextCalls(stub func() (sse.Event, error)) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = stub
}

func (fake *FakeRawEventSource) NextReturns(result1 sse.Event, result2 error) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	fake.nextReturns = struct {
		result1 sse.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeRawEventSource) NextReturnsOnCall(i int, result1 sse.Event, result2 error) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	if fake.nextReturnsOnCall == nil {
		fake.nextReturnsOnCall = make(map[int]struct {
			result1 sse.Event
			result2 error
		})
	}
	fake.nextReturnsOnCall[i] = struct {
		result1 sse.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeRawEventSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRawEventSource) recordInvocation(key string, args []interface{}) {
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

var _ events.RawEventSource = new(FakeRawEventSource)
