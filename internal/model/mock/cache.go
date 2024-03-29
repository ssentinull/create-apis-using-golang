// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/cache.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCacheRepository is a mock of CacheRepository interface.
type MockCacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCacheRepositoryMockRecorder
}

// MockCacheRepositoryMockRecorder is the mock recorder for MockCacheRepository.
type MockCacheRepositoryMockRecorder struct {
	mock *MockCacheRepository
}

// NewMockCacheRepository creates a new mock instance.
func NewMockCacheRepository(ctrl *gomock.Controller) *MockCacheRepository {
	mock := &MockCacheRepository{ctrl: ctrl}
	mock.recorder = &MockCacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheRepository) EXPECT() *MockCacheRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCacheRepository) Delete(ctx context.Context, keys ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheRepositoryMockRecorder) Delete(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCacheRepository)(nil).Delete), varargs...)
}

// Get mocks base method.
func (m *MockCacheRepository) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheRepositoryMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheRepository)(nil).Get), ctx, key)
}

// HashGet mocks base method.
func (m *MockCacheRepository) HashGet(ctx context.Context, hash, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashGet", ctx, hash, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashGet indicates an expected call of HashGet.
func (mr *MockCacheRepositoryMockRecorder) HashGet(ctx, hash, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashGet", reflect.TypeOf((*MockCacheRepository)(nil).HashGet), ctx, hash, key)
}

// HashSet mocks base method.
func (m *MockCacheRepository) HashSet(ctx context.Context, hash, key, val string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashSet", ctx, hash, key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// HashSet indicates an expected call of HashSet.
func (mr *MockCacheRepositoryMockRecorder) HashSet(ctx, hash, key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashSet", reflect.TypeOf((*MockCacheRepository)(nil).HashSet), ctx, hash, key, val)
}

// Set mocks base method.
func (m *MockCacheRepository) Set(ctx context.Context, key, val string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheRepositoryMockRecorder) Set(ctx, key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheRepository)(nil).Set), ctx, key, val)
}
