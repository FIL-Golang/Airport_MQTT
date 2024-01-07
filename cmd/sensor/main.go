package main

import (
	sensor2 "Airport_MQTT/internal/sensor"
	"encoding/json"
	"fmt"
	"os"
)

type SensorData struct {
	Sensors    []sensor2.Sensor `json:"sensors"`
	MQTTConfig sensor2.Config   `json:"mqttConfig"`
}

func main() {
	configFile, err := os.ReadFile("internal/sensor/data.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de sensor:", err)
		return
	}

	var sensorData SensorData
	if err := json.Unmarshal(configFile, &sensorData); err != nil {
		fmt.Println("Erreur lors du parsing du fichier de données sensor:", err)
		return
	}

	for _, sc := range sensorData.Sensors {
		go func(sc sensor2.Sensor) {
			dataChannel := make(chan float64)

			go sensor2.StartCaptureData(sc.AirportCode, sc.Measurement, sc.Frequency, dataChannel)

			for data := range dataChannel {
				fmt.Printf("%s: %f\n", sc.Measurement, data)
			}
		}(sc)
	}

	fmt.Println("Appuyez sur 'Enter' pour arrêter.")
	fmt.Scanln()
}
