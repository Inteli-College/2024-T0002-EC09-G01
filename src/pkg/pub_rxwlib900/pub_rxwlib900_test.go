package pub_rxwlib900

import (
	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	"testing"

	RXWLIB900 "2024-T0002-EC09-G01/src/internal/rxwlib900"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func TestController(t *testing.T) {

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var messageChannel = make(chan mqtt.Message)
	t.Run("TestPublishFields", func(t *testing.T) {
		senddata := SendRadiationData{
			CurrentTime:   time.Now(),
			RadiationData: RXWLIB900.CreateGasesValues(),
		}

		publishpacket := PublishPacket{
			PacketId:   1,
			TopicName:  "sensors/radiation/1",
			Qos:        1,
			RetainFlag: false,
			Payload:    senddata,
			DupFlag:    false,
		}
		fmt.Print("Hi")
		payload, _ := publishpacket.ToJSON()
		
		client.Subscribe("sensors/radiation/1", 1, func(client mqtt.Client, message mqtt.Message) {
			fmt.Print("Subscribed to topic: sensors/radiation/+")
			resultado := message.Payload()
			if messageChannel != nil {
				t.Error("Test")
			}
			fmt.Printf("%s", resultado)

		})
		client.Publish("sensors/radiation/1", 1, false, payload)
		fmt.Print("Hi")

	})
}
