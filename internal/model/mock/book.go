// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/book.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ssentinull/create-apis-using-golang/internal/model"
)

// MockBookUsecase is a mock of BookUsecase interface.
type MockBookUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockBookUsecaseMockRecorder
}

// MockBookUsecaseMockRecorder is the mock recorder for MockBookUsecase.
type MockBookUsecaseMockRecorder struct {
	mock *MockBookUsecase
}

// NewMockBookUsecase creates a new mock instance.
func NewMockBookUsecase(ctrl *gomock.Controller) *MockBookUsecase {
	mock := &MockBookUsecase{ctrl: ctrl}
	mock.recorder = &MockBookUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookUsecase) EXPECT() *MockBookUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBookUsecase) Create(ctx context.Context, input *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBookUsecaseMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookUsecase)(nil).Create), ctx, input)
}

// DeleteByID mocks base method.
func (m *MockBookUsecase) DeleteByID(ctx context.Context, ID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockBookUsecaseMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockBookUsecase)(nil).DeleteByID), ctx, ID)
}

// FindAll mocks base method.
func (m *MockBookUsecase) FindAll(ctx context.Context, query model.GetBooksQueryParams) ([]*model.Book, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, query)
	ret0, _ := ret[0].([]*model.Book)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAll indicates an expected call of FindAll.
func (mr *MockBookUsecaseMockRecorder) FindAll(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockBookUsecase)(nil).FindAll), ctx, query)
}

// FindByID mocks base method.
func (m *MockBookUsecase) FindByID(ctx context.Context, ID int64) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, ID)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockBookUsecaseMockRecorder) FindByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockBookUsecase)(nil).FindByID), ctx, ID)
}

// Update mocks base method.
func (m *MockBookUsecase) Update(ctx context.Context, input *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, input)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBookUsecaseMockRecorder) Update(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookUsecase)(nil).Update), ctx, input)
}

// MockBookRepository is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository.
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// CountAll mocks base method.
func (m *MockBookRepository) CountAll(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAll", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAll indicates an expected call of CountAll.
func (mr *MockBookRepositoryMockRecorder) CountAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAll", reflect.TypeOf((*MockBookRepository)(nil).CountAll), ctx)
}

// Create mocks base method.
func (m *MockBookRepository) Create(ctx context.Context, input *model.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBookRepositoryMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookRepository)(nil).Create), ctx, input)
}

// DeleteByID mocks base method.
func (m *MockBookRepository) DeleteByID(ctx context.Context, ID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockBookRepositoryMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockBookRepository)(nil).DeleteByID), ctx, ID)
}

// FindAll mocks base method.
func (m *MockBookRepository) FindAll(ctx context.Context, query model.GetBooksQueryParams) ([]*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, query)
	ret0, _ := ret[0].([]*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockBookRepositoryMockRecorder) FindAll(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockBookRepository)(nil).FindAll), ctx, query)
}

// FindByID mocks base method.
func (m *MockBookRepository) FindByID(ctx context.Context, ID int64) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, ID)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockBookRepositoryMockRecorder) FindByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockBookRepository)(nil).FindByID), ctx, ID)
}

// Update mocks base method.
func (m *MockBookRepository) Update(ctx context.Context, input *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, input)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBookRepositoryMockRecorder) Update(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookRepository)(nil).Update), ctx, input)
}
