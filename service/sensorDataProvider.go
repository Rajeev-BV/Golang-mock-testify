package service

type sensorData struct {
	VendorID    string
	machineID   string
	sensorType  string
	sensorValue int
}

type sensorDataProvider interface {
	FetchSensorData(VendorID string) ([]sensorData, error)
}
