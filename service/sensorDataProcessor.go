package service

import "errors"

type sensorDataProcessor struct {
	_sensorDataProvider sensorDataProvider
	_emailProvider      emailProvider
}

func (g sensorDataProcessor) processSensorData() ([]string, error) {
	var faultySensorData []string
	sensorData, err := g._sensorDataProvider.FetchSensorData("5")
	if err != nil {
		return faultySensorData, errors.New("not able to fetch data")
	}
	for _, data := range sensorData {
		if (data.sensorType == "H") && (data.sensorValue > 35) {
			faultySensorData = append(faultySensorData, data.sensorType)
		}
	}
	if len(faultySensorData) > 1 {
		g._emailProvider.sendEmail()
	}

	return faultySensorData, nil
}
