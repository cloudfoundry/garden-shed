// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/layercake"
	"github.com/cloudfoundry-incubator/garden-shed/layercake/cleaner"
	"github.com/pivotal-golang/lager"
)

type FakeThreshold struct {
	ExceededStub        func(log lager.Logger, cake layercake.Cake) bool
	exceededMutex       sync.RWMutex
	exceededArgsForCall []struct {
		log  lager.Logger
		cake layercake.Cake
	}
	exceededReturns struct {
		result1 bool
	}
}

func (fake *FakeThreshold) Exceeded(log lager.Logger, cake layercake.Cake) bool {
	fake.exceededMutex.Lock()
	fake.exceededArgsForCall = append(fake.exceededArgsForCall, struct {
		log  lager.Logger
		cake layercake.Cake
	}{log, cake})
	fake.exceededMutex.Unlock()
	if fake.ExceededStub != nil {
		return fake.ExceededStub(log, cake)
	} else {
		return fake.exceededReturns.result1
	}
}

func (fake *FakeThreshold) ExceededCallCount() int {
	fake.exceededMutex.RLock()
	defer fake.exceededMutex.RUnlock()
	return len(fake.exceededArgsForCall)
}

func (fake *FakeThreshold) ExceededArgsForCall(i int) (lager.Logger, layercake.Cake) {
	fake.exceededMutex.RLock()
	defer fake.exceededMutex.RUnlock()
	return fake.exceededArgsForCall[i].log, fake.exceededArgsForCall[i].cake
}

func (fake *FakeThreshold) ExceededReturns(result1 bool) {
	fake.ExceededStub = nil
	fake.exceededReturns = struct {
		result1 bool
	}{result1}
}

var _ cleaner.Threshold = new(FakeThreshold)