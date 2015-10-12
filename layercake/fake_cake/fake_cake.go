// This file was generated by counterfeiter
package fake_cake

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/layercake"
	"github.com/docker/docker/image"
	"github.com/docker/docker/pkg/archive"
)

type FakeCake struct {
	DriverNameStub        func() string
	driverNameMutex       sync.RWMutex
	driverNameArgsForCall []struct{}
	driverNameReturns     struct {
		result1 string
	}
	CreateStub        func(containerID, imageID layercake.ID) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		containerID layercake.ID
		imageID     layercake.ID
	}
	createReturns struct {
		result1 error
	}
	RegisterStub        func(img *image.Image, layer archive.ArchiveReader) error
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		img   *image.Image
		layer archive.ArchiveReader
	}
	registerReturns struct {
		result1 error
	}
	GetStub        func(id layercake.ID) (*image.Image, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		id layercake.ID
	}
	getReturns struct {
		result1 *image.Image
		result2 error
	}
	RemoveStub        func(id layercake.ID) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		id layercake.ID
	}
	removeReturns struct {
		result1 error
	}
	PathStub        func(id layercake.ID) (string, error)
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
		id layercake.ID
	}
	pathReturns struct {
		result1 string
		result2 error
	}
	IsLeafStub        func(id layercake.ID) (bool, error)
	isLeafMutex       sync.RWMutex
	isLeafArgsForCall []struct {
		id layercake.ID
	}
	isLeafReturns struct {
		result1 bool
		result2 error
	}
}

func (fake *FakeCake) DriverName() string {
	fake.driverNameMutex.Lock()
	fake.driverNameArgsForCall = append(fake.driverNameArgsForCall, struct{}{})
	fake.driverNameMutex.Unlock()
	if fake.DriverNameStub != nil {
		return fake.DriverNameStub()
	} else {
		return fake.driverNameReturns.result1
	}
}

func (fake *FakeCake) DriverNameCallCount() int {
	fake.driverNameMutex.RLock()
	defer fake.driverNameMutex.RUnlock()
	return len(fake.driverNameArgsForCall)
}

func (fake *FakeCake) DriverNameReturns(result1 string) {
	fake.DriverNameStub = nil
	fake.driverNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCake) Create(containerID layercake.ID, imageID layercake.ID) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		containerID layercake.ID
		imageID     layercake.ID
	}{containerID, imageID})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(containerID, imageID)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeCake) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeCake) CreateArgsForCall(i int) (layercake.ID, layercake.ID) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].containerID, fake.createArgsForCall[i].imageID
}

func (fake *FakeCake) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCake) Register(img *image.Image, layer archive.ArchiveReader) error {
	fake.registerMutex.Lock()
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		img   *image.Image
		layer archive.ArchiveReader
	}{img, layer})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(img, layer)
	} else {
		return fake.registerReturns.result1
	}
}

func (fake *FakeCake) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeCake) RegisterArgsForCall(i int) (*image.Image, archive.ArchiveReader) {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].img, fake.registerArgsForCall[i].layer
}

func (fake *FakeCake) RegisterReturns(result1 error) {
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCake) Get(id layercake.ID) (*image.Image, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(id)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeCake) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeCake) GetArgsForCall(i int) layercake.ID {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].id
}

func (fake *FakeCake) GetReturns(result1 *image.Image, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *image.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) Remove(id layercake.ID) error {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(id)
	} else {
		return fake.removeReturns.result1
	}
}

func (fake *FakeCake) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeCake) RemoveArgsForCall(i int) layercake.ID {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].id
}

func (fake *FakeCake) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCake) Path(id layercake.ID) (string, error) {
	fake.pathMutex.Lock()
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub(id)
	} else {
		return fake.pathReturns.result1, fake.pathReturns.result2
	}
}

func (fake *FakeCake) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeCake) PathArgsForCall(i int) layercake.ID {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return fake.pathArgsForCall[i].id
}

func (fake *FakeCake) PathReturns(result1 string, result2 error) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) IsLeaf(id layercake.ID) (bool, error) {
	fake.isLeafMutex.Lock()
	fake.isLeafArgsForCall = append(fake.isLeafArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.isLeafMutex.Unlock()
	if fake.IsLeafStub != nil {
		return fake.IsLeafStub(id)
	} else {
		return fake.isLeafReturns.result1, fake.isLeafReturns.result2
	}
}

func (fake *FakeCake) IsLeafCallCount() int {
	fake.isLeafMutex.RLock()
	defer fake.isLeafMutex.RUnlock()
	return len(fake.isLeafArgsForCall)
}

func (fake *FakeCake) IsLeafArgsForCall(i int) layercake.ID {
	fake.isLeafMutex.RLock()
	defer fake.isLeafMutex.RUnlock()
	return fake.isLeafArgsForCall[i].id
}

func (fake *FakeCake) IsLeafReturns(result1 bool, result2 error) {
	fake.IsLeafStub = nil
	fake.isLeafReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

var _ layercake.Cake = new(FakeCake)
