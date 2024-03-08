package testing

import (
	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	controller "2024-T0002-EC09-G01/src/pkg/controller"
	// publisher "2024-T0002-EC09-G01/src/pkg/publisher"
	// subscriber "2024-T0002-EC09-G01/src/pkg/subscriber"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"regexp"
	"testing"
	"time"
)

func ReturnRegex(topic string) *regexp.Regexp {
	var re *regexp.Regexp
	switch {
	case topic == "sensor/gases":
		regex := `^\{"packet-id":\d+,"topic-name":"sensor\/gases\/\d+","qos":\d+,"retain-flag":(?:true|false),"payload":\{"current_time":"[^"]+","gases-values":\{"sensor":"[^"]+","unit":"[^"]+","gases-values":\{"carbon_monoxide":\d+\.\d+,"nitrogen_dioxide":\d+\.\d+,"ethanol":\d+\.\d+,"hydrogen":\d+\.\d+,"ammonia":\d+\.\d+,"methane":\d+\.\d+,"propane":\d+\.\d+,"iso_butane":\d+\.\d+\}\}\},"duplicated-flag":(?:true|false)\}$`
		re = regexp.MustCompile(regex)
	case topic == "sensor/radiation":
		regex := `^\{"packet-id":\d+,"topic-name":"sensor\/radiation\/\d+","qos":\d+,"retain-flag":(?:true|false),"payload":\{"current_time":"[^"]+","radiation-values":\{"sensor":"[^"]+","unit":"[^"]+","radiation-values":\{"radiation":\d+\.\d+\}\}\},"duplicated-flag":(?:true|false)\}$`

		re = regexp.MustCompile(regex)
	}
	return re

}

func TestController(t *testing.T) {
	client := DefaultClient.CreateClient(DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	t.Run("TestPublishFields", func(t *testing.T) {

		var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

			resultado := string(msg.Payload())
			topic := msg.Topic()
			fmt.Printf("Recebido: %s do tópico: %s\n", resultado, topic)

			if ReturnRegex(topic).MatchString(resultado) {
				fmt.Print("\nPayload fits all the publish fields\n")
			} else {
				t.Error("\nPayload does not fit all the publish fields\n")
			}
		}
		
		token := client.Subscribe("sensor/#", 1, messageHandler)
		token.Wait()

		client.Publish("sensor/radiation", 1, false, controller.Payload(1, 1))
		client.Publish("sensor/radiation", 1, false, controller.Payload(0, 1))
		client.Disconnect(250)
	})

	t.Run("TestQos", func(t *testing.T) {
		var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

			if msg.Qos() != 1 {
				t.Error("QoS is not 1")
			}
			fmt.Printf("QoS recebido: %d\n", msg.Qos())
		}

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		if token := client.Subscribe("sensor/#", 1, messageHandler); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			return
		}

		time.Sleep(2 * time.Second)
		client.Disconnect(250)

	})

}
