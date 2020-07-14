// Code generated by counterfeiter. DO NOT EDIT.
package clusterfakes

import (
	"sync"

	"github.ibm.com/coligo/satcon-client/client/actions/cluster"
)

type FakeClusterService struct {
	ClustersByOrgIDStub        func(string, string) (cluster.ClusterList, error)
	clustersByOrgIDMutex       sync.RWMutex
	clustersByOrgIDArgsForCall []struct {
		arg1 string
		arg2 string
	}
	clustersByOrgIDReturns struct {
		result1 cluster.ClusterList
		result2 error
	}
	clustersByOrgIDReturnsOnCall map[int]struct {
		result1 cluster.ClusterList
		result2 error
	}
	DeleteClusterByClusterIDStub        func(string, string, string) (*cluster.DeleteClustersResponseDataDetails, error)
	deleteClusterByClusterIDMutex       sync.RWMutex
	deleteClusterByClusterIDArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	deleteClusterByClusterIDReturns struct {
		result1 *cluster.DeleteClustersResponseDataDetails
		result2 error
	}
	deleteClusterByClusterIDReturnsOnCall map[int]struct {
		result1 *cluster.DeleteClustersResponseDataDetails
		result2 error
	}
	RegisterClusterStub        func(string, cluster.Registration, string) (*cluster.RegisterClusterResponseDataDetails, error)
	registerClusterMutex       sync.RWMutex
	registerClusterArgsForCall []struct {
		arg1 string
		arg2 cluster.Registration
		arg3 string
	}
	registerClusterReturns struct {
		result1 *cluster.RegisterClusterResponseDataDetails
		result2 error
	}
	registerClusterReturnsOnCall map[int]struct {
		result1 *cluster.RegisterClusterResponseDataDetails
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClusterService) ClustersByOrgID(arg1 string, arg2 string) (cluster.ClusterList, error) {
	fake.clustersByOrgIDMutex.Lock()
	ret, specificReturn := fake.clustersByOrgIDReturnsOnCall[len(fake.clustersByOrgIDArgsForCall)]
	fake.clustersByOrgIDArgsForCall = append(fake.clustersByOrgIDArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ClustersByOrgID", []interface{}{arg1, arg2})
	fake.clustersByOrgIDMutex.Unlock()
	if fake.ClustersByOrgIDStub != nil {
		return fake.ClustersByOrgIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.clustersByOrgIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClusterService) ClustersByOrgIDCallCount() int {
	fake.clustersByOrgIDMutex.RLock()
	defer fake.clustersByOrgIDMutex.RUnlock()
	return len(fake.clustersByOrgIDArgsForCall)
}

func (fake *FakeClusterService) ClustersByOrgIDCalls(stub func(string, string) (cluster.ClusterList, error)) {
	fake.clustersByOrgIDMutex.Lock()
	defer fake.clustersByOrgIDMutex.Unlock()
	fake.ClustersByOrgIDStub = stub
}

func (fake *FakeClusterService) ClustersByOrgIDArgsForCall(i int) (string, string) {
	fake.clustersByOrgIDMutex.RLock()
	defer fake.clustersByOrgIDMutex.RUnlock()
	argsForCall := fake.clustersByOrgIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClusterService) ClustersByOrgIDReturns(result1 cluster.ClusterList, result2 error) {
	fake.clustersByOrgIDMutex.Lock()
	defer fake.clustersByOrgIDMutex.Unlock()
	fake.ClustersByOrgIDStub = nil
	fake.clustersByOrgIDReturns = struct {
		result1 cluster.ClusterList
		result2 error
	}{result1, result2}
}

func (fake *FakeClusterService) ClustersByOrgIDReturnsOnCall(i int, result1 cluster.ClusterList, result2 error) {
	fake.clustersByOrgIDMutex.Lock()
	defer fake.clustersByOrgIDMutex.Unlock()
	fake.ClustersByOrgIDStub = nil
	if fake.clustersByOrgIDReturnsOnCall == nil {
		fake.clustersByOrgIDReturnsOnCall = make(map[int]struct {
			result1 cluster.ClusterList
			result2 error
		})
	}
	fake.clustersByOrgIDReturnsOnCall[i] = struct {
		result1 cluster.ClusterList
		result2 error
	}{result1, result2}
}

func (fake *FakeClusterService) DeleteClusterByClusterID(arg1 string, arg2 string, arg3 string) (*cluster.DeleteClustersResponseDataDetails, error) {
	fake.deleteClusterByClusterIDMutex.Lock()
	ret, specificReturn := fake.deleteClusterByClusterIDReturnsOnCall[len(fake.deleteClusterByClusterIDArgsForCall)]
	fake.deleteClusterByClusterIDArgsForCall = append(fake.deleteClusterByClusterIDArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteClusterByClusterID", []interface{}{arg1, arg2, arg3})
	fake.deleteClusterByClusterIDMutex.Unlock()
	if fake.DeleteClusterByClusterIDStub != nil {
		return fake.DeleteClusterByClusterIDStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteClusterByClusterIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClusterService) DeleteClusterByClusterIDCallCount() int {
	fake.deleteClusterByClusterIDMutex.RLock()
	defer fake.deleteClusterByClusterIDMutex.RUnlock()
	return len(fake.deleteClusterByClusterIDArgsForCall)
}

func (fake *FakeClusterService) DeleteClusterByClusterIDCalls(stub func(string, string, string) (*cluster.DeleteClustersResponseDataDetails, error)) {
	fake.deleteClusterByClusterIDMutex.Lock()
	defer fake.deleteClusterByClusterIDMutex.Unlock()
	fake.DeleteClusterByClusterIDStub = stub
}

func (fake *FakeClusterService) DeleteClusterByClusterIDArgsForCall(i int) (string, string, string) {
	fake.deleteClusterByClusterIDMutex.RLock()
	defer fake.deleteClusterByClusterIDMutex.RUnlock()
	argsForCall := fake.deleteClusterByClusterIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClusterService) DeleteClusterByClusterIDReturns(result1 *cluster.DeleteClustersResponseDataDetails, result2 error) {
	fake.deleteClusterByClusterIDMutex.Lock()
	defer fake.deleteClusterByClusterIDMutex.Unlock()
	fake.DeleteClusterByClusterIDStub = nil
	fake.deleteClusterByClusterIDReturns = struct {
		result1 *cluster.DeleteClustersResponseDataDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeClusterService) DeleteClusterByClusterIDReturnsOnCall(i int, result1 *cluster.DeleteClustersResponseDataDetails, result2 error) {
	fake.deleteClusterByClusterIDMutex.Lock()
	defer fake.deleteClusterByClusterIDMutex.Unlock()
	fake.DeleteClusterByClusterIDStub = nil
	if fake.deleteClusterByClusterIDReturnsOnCall == nil {
		fake.deleteClusterByClusterIDReturnsOnCall = make(map[int]struct {
			result1 *cluster.DeleteClustersResponseDataDetails
			result2 error
		})
	}
	fake.deleteClusterByClusterIDReturnsOnCall[i] = struct {
		result1 *cluster.DeleteClustersResponseDataDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeClusterService) RegisterCluster(arg1 string, arg2 cluster.Registration, arg3 string) (*cluster.RegisterClusterResponseDataDetails, error) {
	fake.registerClusterMutex.Lock()
	ret, specificReturn := fake.registerClusterReturnsOnCall[len(fake.registerClusterArgsForCall)]
	fake.registerClusterArgsForCall = append(fake.registerClusterArgsForCall, struct {
		arg1 string
		arg2 cluster.Registration
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("RegisterCluster", []interface{}{arg1, arg2, arg3})
	fake.registerClusterMutex.Unlock()
	if fake.RegisterClusterStub != nil {
		return fake.RegisterClusterStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.registerClusterReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClusterService) RegisterClusterCallCount() int {
	fake.registerClusterMutex.RLock()
	defer fake.registerClusterMutex.RUnlock()
	return len(fake.registerClusterArgsForCall)
}

func (fake *FakeClusterService) RegisterClusterCalls(stub func(string, cluster.Registration, string) (*cluster.RegisterClusterResponseDataDetails, error)) {
	fake.registerClusterMutex.Lock()
	defer fake.registerClusterMutex.Unlock()
	fake.RegisterClusterStub = stub
}

func (fake *FakeClusterService) RegisterClusterArgsForCall(i int) (string, cluster.Registration, string) {
	fake.registerClusterMutex.RLock()
	defer fake.registerClusterMutex.RUnlock()
	argsForCall := fake.registerClusterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClusterService) RegisterClusterReturns(result1 *cluster.RegisterClusterResponseDataDetails, result2 error) {
	fake.registerClusterMutex.Lock()
	defer fake.registerClusterMutex.Unlock()
	fake.RegisterClusterStub = nil
	fake.registerClusterReturns = struct {
		result1 *cluster.RegisterClusterResponseDataDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeClusterService) RegisterClusterReturnsOnCall(i int, result1 *cluster.RegisterClusterResponseDataDetails, result2 error) {
	fake.registerClusterMutex.Lock()
	defer fake.registerClusterMutex.Unlock()
	fake.RegisterClusterStub = nil
	if fake.registerClusterReturnsOnCall == nil {
		fake.registerClusterReturnsOnCall = make(map[int]struct {
			result1 *cluster.RegisterClusterResponseDataDetails
			result2 error
		})
	}
	fake.registerClusterReturnsOnCall[i] = struct {
		result1 *cluster.RegisterClusterResponseDataDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeClusterService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clustersByOrgIDMutex.RLock()
	defer fake.clustersByOrgIDMutex.RUnlock()
	fake.deleteClusterByClusterIDMutex.RLock()
	defer fake.deleteClusterByClusterIDMutex.RUnlock()
	fake.registerClusterMutex.RLock()
	defer fake.registerClusterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClusterService) recordInvocation(key string, args []interface{}) {
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

var _ cluster.ClusterService = new(FakeClusterService)
