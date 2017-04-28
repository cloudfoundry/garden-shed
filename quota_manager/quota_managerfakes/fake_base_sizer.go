// This file was generated by counterfeiter
package quota_managerfakes

import (
	"sync"

	"code.cloudfoundry.org/garden-shed/quota_manager"
	"code.cloudfoundry.org/lager"
)

type FakeBaseSizer struct {
	BaseSizeStub        func(logger lager.Logger, rootfsPath string) (uint64, error)
	baseSizeMutex       sync.RWMutex
	baseSizeArgsForCall []struct {
		logger     lager.Logger
		rootfsPath string
	}
	baseSizeReturns struct {
		result1 uint64
		result2 error
	}
	baseSizeReturnsOnCall map[int]struct {
		result1 uint64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBaseSizer) BaseSize(logger lager.Logger, rootfsPath string) (uint64, error) {
	fake.baseSizeMutex.Lock()
	ret, specificReturn := fake.baseSizeReturnsOnCall[len(fake.baseSizeArgsForCall)]
	fake.baseSizeArgsForCall = append(fake.baseSizeArgsForCall, struct {
		logger     lager.Logger
		rootfsPath string
	}{logger, rootfsPath})
	fake.recordInvocation("BaseSize", []interface{}{logger, rootfsPath})
	fake.baseSizeMutex.Unlock()
	if fake.BaseSizeStub != nil {
		return fake.BaseSizeStub(logger, rootfsPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.baseSizeReturns.result1, fake.baseSizeReturns.result2
}

func (fake *FakeBaseSizer) BaseSizeCallCount() int {
	fake.baseSizeMutex.RLock()
	defer fake.baseSizeMutex.RUnlock()
	return len(fake.baseSizeArgsForCall)
}

func (fake *FakeBaseSizer) BaseSizeArgsForCall(i int) (lager.Logger, string) {
	fake.baseSizeMutex.RLock()
	defer fake.baseSizeMutex.RUnlock()
	return fake.baseSizeArgsForCall[i].logger, fake.baseSizeArgsForCall[i].rootfsPath
}

func (fake *FakeBaseSizer) BaseSizeReturns(result1 uint64, result2 error) {
	fake.BaseSizeStub = nil
	fake.baseSizeReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseSizer) BaseSizeReturnsOnCall(i int, result1 uint64, result2 error) {
	fake.BaseSizeStub = nil
	if fake.baseSizeReturnsOnCall == nil {
		fake.baseSizeReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 error
		})
	}
	fake.baseSizeReturnsOnCall[i] = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseSizer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.baseSizeMutex.RLock()
	defer fake.baseSizeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBaseSizer) recordInvocation(key string, args []interface{}) {
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

var _ quota_manager.BaseSizer = new(FakeBaseSizer)
