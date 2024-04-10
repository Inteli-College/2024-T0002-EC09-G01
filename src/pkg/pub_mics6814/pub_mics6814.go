package pub_mics6814

import (
	MICS6814 "2024-T0002-EC09-G01/src/internal/sensors/mics6814"
	"encoding/json"
	"fmt"
	"time"
)

type PublishPacketGases struct {
	PacketId   string        `json:"packet-id"`
	TopicName  string        `json:"topic-name"`
	Qos        int           `json:"qos"`
	RetainFlag bool          `json:"retain-flag"`
	Payload    SendGasesData `json:"payload"`
	DupFlag    bool          `json:"duplicated-flag"`
}

type SendGasesData struct {
	CurrentTime time.Time             `json:"current_time"`
	GasesData   MICS6814.SensorConfig `json:"gases-values"`
}

func (s *PublishPacketGases) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func CreatePayloadGases(id string) string {
	senddata := SendGasesData{
		CurrentTime: time.Now(),
		GasesData:   MICS6814.CreateGasesValues(),
	}

	publishpacket := PublishPacketGases{
		PacketId:   id,
		TopicName:  fmt.Sprintf("sensor/gases/%s", id),
		Qos:        1,
		RetainFlag: false,
		Payload:    senddata,
		DupFlag:    false,
	}

	payload, _ := publishpacket.ToJSON()
	return payload
}
