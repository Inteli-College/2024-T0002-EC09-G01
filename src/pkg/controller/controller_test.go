package controller

import (
	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	publisher "2024-T0002-EC09-G01/src/pkg/publisher"
	subscriber "2024-T0002-EC09-G01/src/pkg/subscriber"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	// "sync"
	"regexp"
	"testing"
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

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	t.Run("TestPublishFields", func(t *testing.T) {

		subscriber.Subscribe("sensor/+", client, func(client mqtt.Client, msg mqtt.Message) {
			resultado := string(msg.Payload())
			topic := msg.Topic()
			fmt.Printf("Recebido: %s do tópico: %s\n", resultado, topic)

			if ReturnRegex(topic).MatchString(resultado) {
				fmt.Print("\nPayload above fits all the publish fields\n")
			} else {
				t.Error("\nPayload above does not fit all the publish fields\n")
			}
		})
		publisher.Publish(Payload(1, 1), "sensor/radiation", client)
		publisher.Publish(Payload(0, 1), "sensor/gases", client)
	})
}
