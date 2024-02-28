package pub_rxwlib900

import (
	RXWLIB900 "2024-T0002-EC09-G01/src/internal/rxwlib900"
	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
)


type SendRadiationData struct {
	CurrentTime   time.Time              `json:"current_time"`
	RadiationData RXWLIB900.SensorConfig `json:"radiation-values"`
}

func (s *SendRadiationData) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}


func ControllerRadiation(id int) {

	client := DefaultClient.CreateClient(DefaultClient.Broker, fmt.Sprintf("publisher-rxwlib900-%s", strconv.Itoa(id)), DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	senddata := SendRadiationData{
		CurrentTime:   time.Now(),
		RadiationData:  RXWLIB900.CreateGasesValues(),
	}

	payload, _ := senddata.ToJSON()

	for {

		token := client.Publish(fmt.Sprintf("sensors/radiation/%s", strconv.Itoa(id)), 1, false, payload)
		token.Wait()
		token.Wait()

		fmt.Printf("Published message in %s: %s\n", fmt.Sprintf("sensors/radiation/%s", strconv.Itoa(id)), payload)

		time.Sleep(2 * time.Second)
	}
}
