package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/taufiqade/gowallet/models"
	httpRequest "github.com/taufiqade/gowallet/models/http/request"
)

// DBUserRepositoryMock is mock of DBUserRepositoryMock interface
// Create Mock for UserRepository
type DBUserRepositoryMock struct {
	ctrl     *gomock.Controller
	recorder *DBUserRecorderMock
}

// DBUserRecorderMock is the mock recorder for DBUserRecorderMock
type DBUserRecorderMock struct {
	mock *DBUserRepositoryMock
}

// NewDBUserRepositoryMock creates a new mock instance
func NewDBUserRepositoryMock(ctrl *gomock.Controller) *DBUserRepositoryMock {
	mock := &DBUserRepositoryMock{ctrl: ctrl}
	mock.recorder = &DBUserRecorderMock{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *DBUserRepositoryMock) EXPECT() *DBUserRecorderMock {
	return m.recorder
}

// GetUserByID mock base method
func (m *DBUserRepositoryMock) GetUserByID(id int) (models.Users, error) {
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *DBUserRecorderMock) GetUserByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*DBUserRepositoryMock)(nil).GetUserByID), id)
}

// Create mock base method
func (m *DBUserRepositoryMock) Create(data *httpRequest.UserRequest) (models.Users, error) {
	ret := m.ctrl.Call(m, "Create", data)
	ret0, _ := ret[0].(models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of GetUserByID
func (mr *DBUserRecorderMock) Create(data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*DBUserRepositoryMock)(nil).Create), data)
}

// GetUserByEmail mock base method
func (m *DBUserRepositoryMock) GetUserByEmail(email string) (models.Users, error) {
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByID
func (mr *DBUserRecorderMock) GetUserByEmail(email interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*DBUserRepositoryMock)(nil).GetUserByEmail), email)
}
