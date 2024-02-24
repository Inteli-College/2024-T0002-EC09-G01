package common

import (
	"testing"

	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
)

func TestClient(t *testing.T) {
	t.Run("Create a Client", func(t *testing.T) {
		client := DefaultClient.CreateClient(DefaultClient.Broker, "go-mqtt-client", DefaultClient.Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}
	})

}