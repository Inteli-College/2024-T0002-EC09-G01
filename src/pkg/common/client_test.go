package common

import "testing"

func TestClient(t *testing.T) {
	t.Run("Create a Client", func(t *testing.T) {
		client := CreateClient(Broker, "go-mqtt-client", Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}
	})

}