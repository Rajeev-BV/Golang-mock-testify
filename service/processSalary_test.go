package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type dbMock0 struct {
	mock.Mock
}

type dbMock2 struct {
	mock.Mock
}

func (d *dbMock0) FetchEmployeeSalaryDetails(empID string) (empSal, error) {
	args := d.Called(empID)
	return args.Get(0).(empSal), args.Error(1)
}

func (d *dbMock0) FetchEmployeeDetails(empID string) (empDetails, error) {
	args := d.Called(empID)
	return args.Get(0).(empDetails), args.Error(1)
}

func (d *dbMock2) CreditSalaryToEmployee(Salary int, AccNumber string) error {
	args := d.Called(Salary, AccNumber)
	return args.Error(0)
}

func TestMockMethodSalaryWithArgs2(t *testing.T) {
	//Arrange
	theDBMock := dbMock0{}
	theDBMockBankAPI := dbMock2{}

	empSal := empSal{basic: 200,
		HRA: 200}
	empDetails := empDetails{empID: 2, empName: "Rajeev", bankAccNumber: "3333",
		ifscCode: "555"}

	theDBMock.On("FetchEmployeeSalaryDetails", "aa").Return(empSal, nil)
	theDBMock.On("FetchEmployeeDetails", "aa").Return(empDetails, nil)
	theDBMockBankAPI.On("CreditSalaryToEmployee", 400, empDetails.bankAccNumber).Return(nil)
	g := FetchSal{&theDBMock, &theDBMockBankAPI}
	//Act
	g.ProcessSalary()
	//Assert
	theDBMockBankAPI.AssertCalled(t, "CreditSalaryToEmployee", 400, empDetails.bankAccNumber)
	theDBMockBankAPI.AssertNumberOfCalls(t, "CreditSalaryToEmployee", 1)
	theDBMockBankAPI.AssertExpectations(t)
}
