package mocks

import "github.com/stretchr/testify/mock"

import "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
import "github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
import "github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
import "github.com/GoogleCloudPlatform/kubernetes/pkg/watch"

type KubeServiceInterface struct {
	mock.Mock
}

func (m *KubeServiceInterface) List(selector labels.Selector) (*api.ServiceList, error) {
	ret := m.Called(selector)

	var r0 *api.ServiceList
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.ServiceList)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubeServiceInterface) Get(name string) (*api.Service, error) {
	ret := m.Called(name)

	var r0 *api.Service
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.Service)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubeServiceInterface) Create(srv *api.Service) (*api.Service, error) {
	ret := m.Called(srv)

	var r0 *api.Service
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.Service)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubeServiceInterface) Update(srv *api.Service) (*api.Service, error) {
	ret := m.Called(srv)

	var r0 *api.Service
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.Service)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubeServiceInterface) Delete(name string) error {
	ret := m.Called(name)

	r0 := ret.Error(0)

	return r0
}
func (m *KubeServiceInterface) Watch(label labels.Selector, field fields.Selector, resourceVersion string) (watch.Interface, error) {
	ret := m.Called(label, field, resourceVersion)

	r0 := ret.Get(0).(watch.Interface)
	r1 := ret.Error(1)

	return r0, r1
}
