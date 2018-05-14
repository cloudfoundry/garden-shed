// Code generated by counterfeiter. DO NOT EDIT.
package runruncfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/runrunc"
)

type FakeUserLookupper struct {
	LookupStub        func(rootFsPath string, user string) (*runrunc.ExecUser, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		rootFsPath string
		user       string
	}
	lookupReturns struct {
		result1 *runrunc.ExecUser
		result2 error
	}
	lookupReturnsOnCall map[int]struct {
		result1 *runrunc.ExecUser
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserLookupper) Lookup(rootFsPath string, user string) (*runrunc.ExecUser, error) {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		rootFsPath string
		user       string
	}{rootFsPath, user})
	fake.recordInvocation("Lookup", []interface{}{rootFsPath, user})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(rootFsPath, user)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lookupReturns.result1, fake.lookupReturns.result2
}

func (fake *FakeUserLookupper) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeUserLookupper) LookupArgsForCall(i int) (string, string) {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].rootFsPath, fake.lookupArgsForCall[i].user
}

func (fake *FakeUserLookupper) LookupReturns(result1 *runrunc.ExecUser, result2 error) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 *runrunc.ExecUser
		result2 error
	}{result1, result2}
}

func (fake *FakeUserLookupper) LookupReturnsOnCall(i int, result1 *runrunc.ExecUser, result2 error) {
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 *runrunc.ExecUser
			result2 error
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 *runrunc.ExecUser
		result2 error
	}{result1, result2}
}

func (fake *FakeUserLookupper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserLookupper) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.UserLookupper = new(FakeUserLookupper)
