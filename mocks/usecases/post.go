// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\Rein\Desktop\my-web-api\usecases\post.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "dot-crud-redis-go-api/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPostUsecase is a mock of PostUsecase interface.
type MockPostUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPostUsecaseMockRecorder
}

// MockPostUsecaseMockRecorder is the mock recorder for MockPostUsecase.
type MockPostUsecaseMockRecorder struct {
	mock *MockPostUsecase
}

// NewMockPostUsecase creates a new mock instance.
func NewMockPostUsecase(ctrl *gomock.Controller) *MockPostUsecase {
	mock := &MockPostUsecase{ctrl: ctrl}
	mock.recorder = &MockPostUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostUsecase) EXPECT() *MockPostUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPostUsecase) Create(post *models.Post) (*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", post)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostUsecaseMockRecorder) Create(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostUsecase)(nil).Create), post)
}

// Delete mocks base method.
func (m *MockPostUsecase) Delete(id int) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockPostUsecaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostUsecase)(nil).Delete), id)
}

// ReadAll mocks base method.
func (m *MockPostUsecase) ReadAll() (*[]models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll")
	ret0, _ := ret[0].(*[]models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockPostUsecaseMockRecorder) ReadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockPostUsecase)(nil).ReadAll))
}

// ReadById mocks base method.
func (m *MockPostUsecase) ReadById(id int) (*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadById", id)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadById indicates an expected call of ReadById.
func (mr *MockPostUsecaseMockRecorder) ReadById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadById", reflect.TypeOf((*MockPostUsecase)(nil).ReadById), id)
}

// Update mocks base method.
func (m *MockPostUsecase) Update(id int, post *models.Post) (*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, post)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPostUsecaseMockRecorder) Update(id, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostUsecase)(nil).Update), id, post)
}
