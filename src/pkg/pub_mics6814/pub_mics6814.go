package pub_mics6814

import (
	MICS6814 "2024-T0002-EC09-G01/src/internal/mics6814"
	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
)

type SendGasesData struct {
	CurrentTime   time.Time              `json:"current_time"`
	GasesData     MICS6814.SensorConfig  `json:"gases-values"`
}

func (s *SendGasesData) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}


func ControllerGases(id int) {

	client := DefaultClient.CreateClient(DefaultClient.Broker, fmt.Sprintf("publisher-mics6814-%s", strconv.Itoa(id)), DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	senddata := SendGasesData{
		CurrentTime:   time.Now(),
		GasesData:     MICS6814.CreateGasesValues(),
	}

	payload, _ := senddata.ToJSON()

	for {

		if client.IsAuthorized("sensors", 1) {
			token := client.Publish(fmt.Sprintf("sensors/gases/%s", strconv.Itoa(id)), 1, false, payload)
			token.Wait()
			token.Wait()

			fmt.Printf("Published message in %s: %s\n", fmt.Sprintf("sensors/gases/%s", strconv.Itoa(id)), payload)
		} else {
			fmt.Println("Client not authorized.")
		}

		time.Sleep(2 * time.Second)
	}
}
