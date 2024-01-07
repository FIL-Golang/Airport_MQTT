package model

func NewSensorData(sensorId string, codeIATA string, nature int, value float32) *SensorData {
	return &SensorData{
		SensorId: sensorId,
		CodeIATA: codeIATA,
		Nature:   nature,
		Value:    value,
	}
}

func SensorNatureFromInt(nature int) string {
	switch nature {
	case Temperature:
		return "temperature"
	case Pressure:
		return "pressure"
	case WindSpeed:
		return "wind speed"
	default:
		return ""
	}
}

func SensorNatureFromString(nature string) int {
	switch nature {
	case "temperature":
		return Temperature
	case "pressure":
		return Pressure
	case "wind speed":
		return WindSpeed
	default:
		return -1
	}
}
