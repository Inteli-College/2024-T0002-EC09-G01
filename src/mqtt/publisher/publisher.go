package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Sensor struct {
	name        string
	latitude    float64
	longitude   float64
	measurement float64
	rate        int
	mqtt.Client
}

func NewSensor(
	name string,
	latitude float64,
	longitude float64,
	measurement float64,
	rate int) *Sensor {

	s := &Sensor{
		name:        name,
		latitude:    latitude,
		longitude:   longitude,
		measurement: measurement,
		rate:        rate,
	}

	s.Client = mqtt.NewClient(mqtt.NewClientOptions())

	return s

}

func (s *Sensor) OnMessageReceived(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())
}

func main() {
	sensor := NewSensor("SPS30", 51.0, 0.0, 0.0, 1)
	opts := mqtt.NewClientOptions().AddBroker("broker.hivemq.com:1883")
	opts.SetClientID("go-mqtt-sensor")
	opts.SetDefaultPublishHandler(sensor.OnMessageReceived)

	sensor.Client = mqtt.NewClient(opts)
	if token := sensor.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "sensors/" + sensor.name

	for {
		sensor.measurement = rand.Float64()*5
		payload := strconv.FormatFloat(sensor.measurement, 'f', 2, 64)
		token := sensor.Publish(topic, 0, false, payload)
		token.Wait()
		fmt.Printf("Published message: %s\n", payload)
		time.Sleep(time.Duration(sensor.rate) * time.Second)
	}
}
