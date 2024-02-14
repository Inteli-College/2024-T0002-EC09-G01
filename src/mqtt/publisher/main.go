package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

const maxSensorRange = 1.0
const minSensorRange = 0.03

const broker = "broker.hivemq.com:1883"
const id = "go-mqtt-sensor"

type Sensor struct {
	Name        string
	Latitude    float64
	Longitude   float64
	Measurement float64
	Rate        int
	Unit        string
}

func NewSensor(
	name string,
	latitude float64,
	longitude float64,
	measurement float64,
	rate int,
	unit string) *Sensor {

	s := &Sensor{
		Name:        name,
		Latitude:    latitude,
		Longitude:   longitude,
		Measurement: measurement,
		Rate:        rate,
		Unit:        unit,
	}

	return s

}

func (s *Sensor) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func main() {

	// Creating an Instance for SPS30 sensor
	// Which can read Mass Concentration within an area
	// With a 1s rate

	sensor := NewSensor("SPS30", 51.0, 0.0, 0.0, 1, "μg/m³")

	// Creating an Instance for MiCS-6814 sensor
	// Which can read substance concentration
	// With a 1s rate

	sensor2 := NewSensor("MiCS-6814", 10.0, 1.0, 0.0, 1, "NO2 - ppm")

	// Creating an Array of Sensors

	var sensors []Sensor
	sensors = append(sensors, *sensor, *sensor2)

	client := CreateClient(broker, id, handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		for _, sensor := range sensors {
			
			topic := "sensors/" + sensor.Name

			sensor.Measurement = (rand.Float64() * (maxSensorRange - minSensorRange)) + minSensorRange

			payload, _ := sensor.ToJSON()

			token := client.Publish(topic, 0, false, payload)

			token.Wait()

			fmt.Printf("Published message: %s\n", payload)

			time.Sleep(time.Duration(sensor.Rate) * time.Second)

		}
	}
}
