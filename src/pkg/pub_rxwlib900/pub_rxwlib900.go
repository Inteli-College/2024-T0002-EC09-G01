package pub_rxwlib900

import (
	RXWLIB900 "2024-T0002-EC09-G01/src/internal/sensors/rxwlib900"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type PublishPacketRadiation struct {
	PacketId   int               `json:"packet-id"`
	TopicName  string            `json:"topic-name"`
	Qos        int               `json:"qos"`
	RetainFlag bool              `json:"retain-flag"`
	Payload    SendRadiationData `json:"payload"`
	DupFlag    bool              `json:"duplicated-flag"`
}

type SendRadiationData struct {
	CurrentTime   time.Time              `json:"current_time"`
	RadiationData RXWLIB900.SensorConfig `json:"radiation-values"`
}

func (s *PublishPacketRadiation) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func CreatePayloadRadiation(id int) string {
	senddata := SendRadiationData{
		CurrentTime:   time.Now(),
		RadiationData: RXWLIB900.CreateGasesValues(),
	}

	publishpacket := PublishPacketRadiation{
		PacketId:   id,
		TopicName:  fmt.Sprintf("sensor/radiation/%s", strconv.Itoa(id)),
		Qos:        1,
		RetainFlag: false,
		Payload:    senddata,
		DupFlag:    false,
	}

	payload, _ := publishpacket.ToJSON()
	return payload
}
