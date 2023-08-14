package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockTransactionLeaveRepo struct {
	mock.Mock
}

func (m *MockTransactionLeaveRepo) Create(payload model.TransactionLeave) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *MockTransactionLeaveRepo) GetByEmployeeID(employeeID string) ([]model.TransactionLeave, error) {
	args := m.Called(employeeID)
	return args.Get(0).([]model.TransactionLeave), args.Error(1)
}

func (m *MockTransactionLeaveRepo) GetByID(id string) (model.TransactionLeave, error) {
	args := m.Called(id)
	return args.Get(0).(model.TransactionLeave), args.Error(1)
}

func (m *MockTransactionLeaveRepo) GetByIdTxNonDto(id string) (model.TransactionLeave, error) {
	args := m.Called(id)
	return args.Get(0).(model.TransactionLeave), args.Error(1)
}

func (m *MockTransactionLeaveRepo) Paging(requestPagung dto.PaginationParam) ([]dto.TransactionResponseDto, dto.Paging, error) {
	args := m.Called(requestPagung)
	return args.Get(0).([]dto.TransactionResponseDto), args.Get(1).(dto.Paging), args.Error(2)
}

func (m *MockTransactionLeaveRepo) UpdateStatus(transactionID string, statusID string) error {
	args := m.Called(transactionID, statusID)
	return args.Error(0)
}

type MockEmployeeUseCase struct {
	mock.Mock
}

func (*MockEmployeeUseCase) FindAllEmpl() ([]model.Employee, error) {
	panic("unimplemented")
}

func (*MockEmployeeUseCase) FindByIdEmpl(id string) (model.Employee, error) {
	panic("unimplemented")
}

func (*MockEmployeeUseCase) PaternityLeave(id string, availableDays int) error {
	panic("unimplemented")
}

func (*MockEmployeeUseCase) RegisterNewEmpl(payload model.Employee) error {
	panic("unimplemented")
}

func (*MockEmployeeUseCase) UpdateAnnualLeave(id string, availableDays int) error {
	panic("unimplemented")
}

func (*MockEmployeeUseCase) UpdateEmpl(payload model.Employee) error {
	panic("unimplemented")
}

func (*MockEmployeeUseCase) UpdateMarriageLeave(id string, availableDays int) error {
	panic("unimplemented")
}

func (*MockEmployeeUseCase) UpdateMaternityLeave(id string, availableDays int) error {
	panic("unimplemented")
}

func (*MockEmployeeUseCase) UpdateMenstrualLeave(id string, availableDays int) error {
	panic("unimplemented")
}

type MockLeaveTypeUseCase struct {
	mock.Mock
}

func (*MockLeaveTypeUseCase) DeleteLeaveType(id string) error {
	panic("unimplemented")
}

func (*MockLeaveTypeUseCase) FindAllLeaveType() ([]model.LeaveType, error) {
	panic("unimplemented")
}

func (*MockLeaveTypeUseCase) FindByIdLeaveType(id string) (model.LeaveType, error) {
	panic("unimplemented")
}

func (*MockLeaveTypeUseCase) FindRoleNameId(id string) (model.Role, error) {
	panic("unimplemented")
}

func (*MockLeaveTypeUseCase) GetByName(name string) (model.LeaveType, error) {
	panic("unimplemented")
}

func (*MockLeaveTypeUseCase) RegisterNewLeaveType(payload model.LeaveType) error {
	panic("unimplemented")
}

func (*MockLeaveTypeUseCase) UpdateLeaveType(payload model.LeaveType) error {
	panic("unimplemented")
}

type MockStatusLeaveUseCase struct {
	mock.Mock
}

func (*MockStatusLeaveUseCase) DeleteStatusLeave(id string) error {
	panic("unimplemented")
}

func (*MockStatusLeaveUseCase) FindAllStatusLeave() ([]model.StatusLeave, error) {
	panic("unimplemented")
}

func (*MockStatusLeaveUseCase) FindByIdStatusLeave(id string) (model.StatusLeave, error) {
	panic("unimplemented")
}

func (*MockStatusLeaveUseCase) FindByNameStatusLeave(statusName string) (model.StatusLeave, error) {
	panic("unimplemented")
}

func (*MockStatusLeaveUseCase) RegisterNewStatusLeave(payload model.StatusLeave) error {
	panic("unimplemented")
}

func (*MockStatusLeaveUseCase) UpdateStatusLeave(payload model.StatusLeave) error {
	panic("unimplemented")
}

type TransactionLeaveUseCaseSuite struct {
	suite.Suite
	MockTransactionLeaveRepo *MockTransactionLeaveRepo
	employeeUC               *MockEmployeeUseCase
	leaveTypeUC              *MockLeaveTypeUseCase
	statusLeaveUC            *MockStatusLeaveUseCase
	transactionUC                     TransactionLeaveUseCase
}

func (suite *TransactionLeaveUseCaseSuite) SetupTest() {
	suite.MockTransactionLeaveRepo = new(MockTransactionLeaveRepo)
	suite.employeeUC = new(MockEmployeeUseCase)
	suite.leaveTypeUC = new(MockLeaveTypeUseCase)
	suite.statusLeaveUC = new(MockStatusLeaveUseCase)

	suite.transactionUC = NewTransactionLeaveUseCase(suite.MockTransactionLeaveRepo, suite.employeeUC, suite.leaveTypeUC, suite.statusLeaveUC)
}

// code test

func TestTransactionLeaveUseCaseSuite(t *testing.T) {
	suite.Run(t, new(StatusLeaveUseCaseSuite))
}
