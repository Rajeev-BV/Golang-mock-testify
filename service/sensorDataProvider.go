package service

import "errors"

type sensorData struct {
	VendorID    string
	machineID   string
	sensorType  string
	sensorValue int
}

type sensorDataProvider interface {
	FetchSensorData(VendorID string) ([]sensorData, error)
}

type sensorDataProviderImpl struct {
	sensorData []sensorData
}

func (g sensorDataProviderImpl) FetchSensorData(VendorID string) ([]sensorData, error) {
	//make http call to get data
	g.sensorData = []sensorData{
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
	if len(g.sensorData) == 0 {
		return nil, errors.New("something wrong")
	}
	return g.sensorData, nil
}
