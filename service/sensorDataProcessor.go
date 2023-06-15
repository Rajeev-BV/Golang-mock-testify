package service

type sensorDataProcessor struct {
	_sensorDataProvider sensorDataProvider
	_emailProvider      emailProvider
}

func (g sensorDataProcessor) processSensorData() []string {
	var faultySensorData []string
	sensorData, _ := g._sensorDataProvider.FetchSensorData("5")
	for _, data := range sensorData {
		if (data.sensorType == "H") && (data.sensorValue > 35) {
			faultySensorData = append(faultySensorData, data.sensorType)
		}
	}
	if len(faultySensorData) > 1 {
		g._emailProvider.sendEmail()
	}

	return faultySensorData
}
