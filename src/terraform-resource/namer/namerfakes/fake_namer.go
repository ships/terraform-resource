// Code generated by counterfeiter. DO NOT EDIT.
package namerfakes

import (
	"sync"
	"terraform-resource/namer"
)

type FakeNamer struct {
	RandomNameStub        func() string
	randomNameMutex       sync.RWMutex
	randomNameArgsForCall []struct {
	}
	randomNameReturns struct {
		result1 string
	}
	randomNameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNamer) RandomName() string {
	fake.randomNameMutex.Lock()
	ret, specificReturn := fake.randomNameReturnsOnCall[len(fake.randomNameArgsForCall)]
	fake.randomNameArgsForCall = append(fake.randomNameArgsForCall, struct {
	}{})
	fake.recordInvocation("RandomName", []interface{}{})
	fake.randomNameMutex.Unlock()
	if fake.RandomNameStub != nil {
		return fake.RandomNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.randomNameReturns
	return fakeReturns.result1
}

func (fake *FakeNamer) RandomNameCallCount() int {
	fake.randomNameMutex.RLock()
	defer fake.randomNameMutex.RUnlock()
	return len(fake.randomNameArgsForCall)
}

func (fake *FakeNamer) RandomNameCalls(stub func() string) {
	fake.randomNameMutex.Lock()
	defer fake.randomNameMutex.Unlock()
	fake.RandomNameStub = stub
}

func (fake *FakeNamer) RandomNameReturns(result1 string) {
	fake.randomNameMutex.Lock()
	defer fake.randomNameMutex.Unlock()
	fake.RandomNameStub = nil
	fake.randomNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeNamer) RandomNameReturnsOnCall(i int, result1 string) {
	fake.randomNameMutex.Lock()
	defer fake.randomNameMutex.Unlock()
	fake.RandomNameStub = nil
	if fake.randomNameReturnsOnCall == nil {
		fake.randomNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.randomNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeNamer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.randomNameMutex.RLock()
	defer fake.randomNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNamer) recordInvocation(key string, args []interface{}) {
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

var _ namer.Namer = new(FakeNamer)
