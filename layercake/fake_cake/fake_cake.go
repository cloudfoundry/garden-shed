// Code generated by counterfeiter. DO NOT EDIT.
package fake_cake

import (
	"sync"

	"code.cloudfoundry.org/garden-shed/layercake"
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
	driverNameReturnsOnCall map[int]struct {
		result1 string
	}
	CreateStub        func(layerID, parentID layercake.ID, containerID string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		layerID     layercake.ID
		parentID    layercake.ID
		containerID string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
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
	registerReturnsOnCall map[int]struct {
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
	getReturnsOnCall map[int]struct {
		result1 *image.Image
		result2 error
	}
	UnmountStub        func(id layercake.ID) error
	unmountMutex       sync.RWMutex
	unmountArgsForCall []struct {
		id layercake.ID
	}
	unmountReturns struct {
		result1 error
	}
	unmountReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveStub        func(id layercake.ID) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		id layercake.ID
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
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
	pathReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	QuotaedPathStub        func(id layercake.ID, quota int64) (string, error)
	quotaedPathMutex       sync.RWMutex
	quotaedPathArgsForCall []struct {
		id    layercake.ID
		quota int64
	}
	quotaedPathReturns struct {
		result1 string
		result2 error
	}
	quotaedPathReturnsOnCall map[int]struct {
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
	isLeafReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetAllLeavesStub        func() ([]layercake.ID, error)
	getAllLeavesMutex       sync.RWMutex
	getAllLeavesArgsForCall []struct{}
	getAllLeavesReturns     struct {
		result1 []layercake.ID
		result2 error
	}
	getAllLeavesReturnsOnCall map[int]struct {
		result1 []layercake.ID
		result2 error
	}
	AllStub        func() []*image.Image
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []*image.Image
	}
	allReturnsOnCall map[int]struct {
		result1 []*image.Image
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCake) DriverName() string {
	fake.driverNameMutex.Lock()
	ret, specificReturn := fake.driverNameReturnsOnCall[len(fake.driverNameArgsForCall)]
	fake.driverNameArgsForCall = append(fake.driverNameArgsForCall, struct{}{})
	fake.recordInvocation("DriverName", []interface{}{})
	fake.driverNameMutex.Unlock()
	if fake.DriverNameStub != nil {
		return fake.DriverNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.driverNameReturns.result1
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

func (fake *FakeCake) DriverNameReturnsOnCall(i int, result1 string) {
	fake.DriverNameStub = nil
	if fake.driverNameReturnsOnCall == nil {
		fake.driverNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.driverNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCake) Create(layerID layercake.ID, parentID layercake.ID, containerID string) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		layerID     layercake.ID
		parentID    layercake.ID
		containerID string
	}{layerID, parentID, containerID})
	fake.recordInvocation("Create", []interface{}{layerID, parentID, containerID})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(layerID, parentID, containerID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *FakeCake) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeCake) CreateArgsForCall(i int) (layercake.ID, layercake.ID, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].layerID, fake.createArgsForCall[i].parentID, fake.createArgsForCall[i].containerID
}

func (fake *FakeCake) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCake) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCake) Register(img *image.Image, layer archive.ArchiveReader) error {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		img   *image.Image
		layer archive.ArchiveReader
	}{img, layer})
	fake.recordInvocation("Register", []interface{}{img, layer})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(img, layer)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.registerReturns.result1
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

func (fake *FakeCake) RegisterReturnsOnCall(i int, result1 error) {
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCake) Get(id layercake.ID) (*image.Image, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.recordInvocation("Get", []interface{}{id})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
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

func (fake *FakeCake) GetReturnsOnCall(i int, result1 *image.Image, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *image.Image
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *image.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) Unmount(id layercake.ID) error {
	fake.unmountMutex.Lock()
	ret, specificReturn := fake.unmountReturnsOnCall[len(fake.unmountArgsForCall)]
	fake.unmountArgsForCall = append(fake.unmountArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.recordInvocation("Unmount", []interface{}{id})
	fake.unmountMutex.Unlock()
	if fake.UnmountStub != nil {
		return fake.UnmountStub(id)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unmountReturns.result1
}

func (fake *FakeCake) UnmountCallCount() int {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return len(fake.unmountArgsForCall)
}

func (fake *FakeCake) UnmountArgsForCall(i int) layercake.ID {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return fake.unmountArgsForCall[i].id
}

func (fake *FakeCake) UnmountReturns(result1 error) {
	fake.UnmountStub = nil
	fake.unmountReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCake) UnmountReturnsOnCall(i int, result1 error) {
	fake.UnmountStub = nil
	if fake.unmountReturnsOnCall == nil {
		fake.unmountReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmountReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCake) Remove(id layercake.ID) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.recordInvocation("Remove", []interface{}{id})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(id)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeReturns.result1
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

func (fake *FakeCake) RemoveReturnsOnCall(i int, result1 error) {
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

func (fake *FakeCake) Path(id layercake.ID) (string, error) {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.recordInvocation("Path", []interface{}{id})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.pathReturns.result1, fake.pathReturns.result2
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

func (fake *FakeCake) PathReturnsOnCall(i int, result1 string, result2 error) {
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) QuotaedPath(id layercake.ID, quota int64) (string, error) {
	fake.quotaedPathMutex.Lock()
	ret, specificReturn := fake.quotaedPathReturnsOnCall[len(fake.quotaedPathArgsForCall)]
	fake.quotaedPathArgsForCall = append(fake.quotaedPathArgsForCall, struct {
		id    layercake.ID
		quota int64
	}{id, quota})
	fake.recordInvocation("QuotaedPath", []interface{}{id, quota})
	fake.quotaedPathMutex.Unlock()
	if fake.QuotaedPathStub != nil {
		return fake.QuotaedPathStub(id, quota)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.quotaedPathReturns.result1, fake.quotaedPathReturns.result2
}

func (fake *FakeCake) QuotaedPathCallCount() int {
	fake.quotaedPathMutex.RLock()
	defer fake.quotaedPathMutex.RUnlock()
	return len(fake.quotaedPathArgsForCall)
}

func (fake *FakeCake) QuotaedPathArgsForCall(i int) (layercake.ID, int64) {
	fake.quotaedPathMutex.RLock()
	defer fake.quotaedPathMutex.RUnlock()
	return fake.quotaedPathArgsForCall[i].id, fake.quotaedPathArgsForCall[i].quota
}

func (fake *FakeCake) QuotaedPathReturns(result1 string, result2 error) {
	fake.QuotaedPathStub = nil
	fake.quotaedPathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) QuotaedPathReturnsOnCall(i int, result1 string, result2 error) {
	fake.QuotaedPathStub = nil
	if fake.quotaedPathReturnsOnCall == nil {
		fake.quotaedPathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.quotaedPathReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) IsLeaf(id layercake.ID) (bool, error) {
	fake.isLeafMutex.Lock()
	ret, specificReturn := fake.isLeafReturnsOnCall[len(fake.isLeafArgsForCall)]
	fake.isLeafArgsForCall = append(fake.isLeafArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.recordInvocation("IsLeaf", []interface{}{id})
	fake.isLeafMutex.Unlock()
	if fake.IsLeafStub != nil {
		return fake.IsLeafStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isLeafReturns.result1, fake.isLeafReturns.result2
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

func (fake *FakeCake) IsLeafReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsLeafStub = nil
	if fake.isLeafReturnsOnCall == nil {
		fake.isLeafReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isLeafReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) GetAllLeaves() ([]layercake.ID, error) {
	fake.getAllLeavesMutex.Lock()
	ret, specificReturn := fake.getAllLeavesReturnsOnCall[len(fake.getAllLeavesArgsForCall)]
	fake.getAllLeavesArgsForCall = append(fake.getAllLeavesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllLeaves", []interface{}{})
	fake.getAllLeavesMutex.Unlock()
	if fake.GetAllLeavesStub != nil {
		return fake.GetAllLeavesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAllLeavesReturns.result1, fake.getAllLeavesReturns.result2
}

func (fake *FakeCake) GetAllLeavesCallCount() int {
	fake.getAllLeavesMutex.RLock()
	defer fake.getAllLeavesMutex.RUnlock()
	return len(fake.getAllLeavesArgsForCall)
}

func (fake *FakeCake) GetAllLeavesReturns(result1 []layercake.ID, result2 error) {
	fake.GetAllLeavesStub = nil
	fake.getAllLeavesReturns = struct {
		result1 []layercake.ID
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) GetAllLeavesReturnsOnCall(i int, result1 []layercake.ID, result2 error) {
	fake.GetAllLeavesStub = nil
	if fake.getAllLeavesReturnsOnCall == nil {
		fake.getAllLeavesReturnsOnCall = make(map[int]struct {
			result1 []layercake.ID
			result2 error
		})
	}
	fake.getAllLeavesReturnsOnCall[i] = struct {
		result1 []layercake.ID
		result2 error
	}{result1, result2}
}

func (fake *FakeCake) All() []*image.Image {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.allReturns.result1
}

func (fake *FakeCake) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *FakeCake) AllReturns(result1 []*image.Image) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []*image.Image
	}{result1}
}

func (fake *FakeCake) AllReturnsOnCall(i int, result1 []*image.Image) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []*image.Image
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []*image.Image
	}{result1}
}

func (fake *FakeCake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.driverNameMutex.RLock()
	defer fake.driverNameMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.quotaedPathMutex.RLock()
	defer fake.quotaedPathMutex.RUnlock()
	fake.isLeafMutex.RLock()
	defer fake.isLeafMutex.RUnlock()
	fake.getAllLeavesMutex.RLock()
	defer fake.getAllLeavesMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCake) recordInvocation(key string, args []interface{}) {
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

var _ layercake.Cake = new(FakeCake)
