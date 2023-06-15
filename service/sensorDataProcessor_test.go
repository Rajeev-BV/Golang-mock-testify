package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSensorDataProvider struct {
	mock.Mock
}

func (d *mockSensorDataProvider) FetchSensorData(vendorID string) ([]sensorData, error) {
	args := d.Called(vendorID)
	return args.Get(0).([]sensorData), args.Error(1)
}

type mockEmailProvider struct {
	mock.Mock
}

func (d *mockEmailProvider) sendEmail() error {
	args := d.Called()
	return args.Error(0)
}

func Test_When_FaultySensors_GreaterThan1_SendEmail(t *testing.T) {
	//Arrange
	mockSensorData := []sensorData{
		{
			VendorID:    "5",
			machineID:   "6",
			sensorType:  "H",
			sensorValue: 60,
		},
		{
			VendorID:    "5",
			machineID:   "6",
			sensorType:  "H",
			sensorValue: 60,
		},
	}
	expectedFaultySensorList := make([]string, 0, 2)
	expectedFaultySensorList = append(expectedFaultySensorList, "H", "H")

	_mockSensorDataProvider := mockSensorDataProvider{}
	_mockSensorDataProvider.On("FetchSensorData", "5").Return(mockSensorData, nil)

	_mockEmailProvider := mockEmailProvider{}
	_mockEmailProvider.On("sendEmail").Return(nil)

	//Act
	g := sensorDataProcessor{&_mockSensorDataProvider, &_mockEmailProvider}
	sensorData := g.processSensorData()
	//Assert
	assert.Equal(t, sensorData, expectedFaultySensorList)
	_mockEmailProvider.AssertCalled(t, "sendEmail")
	_mockEmailProvider.AssertNumberOfCalls(t, "sendEmail", 1)
	_mockEmailProvider.AssertExpectations(t)

}
