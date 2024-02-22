package controller

import (
	MICS6814 "2024-T0002-EC09-G01/src/internal/mics6814"
	RXWLIB900 "2024-T0002-EC09-G01/src/internal/rxwlib900"
	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type SendData struct {
	Identifier    int                    `json:"identifier"`
	Latitude      float64                `json:"latitude"`
	Longitude     float64                `json:"longitude"`
	CurrentTime   time.Time              `json:"current_time"`
	GasesData     MICS6814.SensorConfig  `json:"gases-values"`
	RadiationData RXWLIB900.SensorConfig `json:"radiation-values"`
}

func RandomValues() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() * 100
}

func (s *SendData) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func Controller(id int) {

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	senddata := SendData{
		Identifier:    id,
		Latitude:      RandomValues(),
		Longitude:     RandomValues(),
		CurrentTime:   time.Now(),
		GasesData:     MICS6814.CreateGasesValues(),
		RadiationData: RXWLIB900.CreateGasesValues(),
	}

	payload, _ := senddata.ToJSON()

	for {

		token := client.Publish("sensors", 0, false, payload)
		token.Wait()
		token.Wait()

		fmt.Printf("Published message: %s\n", payload)

		time.Sleep(1 * time.Second)
	}
}
