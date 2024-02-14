package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Create new Sensor", func(t *testing.T) {
		sensor := NewSensor("Sensor1", 51.0, 0.0, 0.0, 60, "μg/m³")
		compare := &Sensor{Name: "Sensor1", Latitude: 51.0, Longitude: 0.0, Measurement: 0.0, Rate: 60, Unit: "μg/m³"}

		if !reflect.DeepEqual(sensor, compare) {
			t.Errorf("The sensor was not created successfully...")
		}
	})

	t.Run("Generating JSON file to payload", func(t *testing.T) {
		sensor := NewSensor("SPS30", 51.0, 0.0, 0.0, 1, "μg/m³")

		got, err := sensor.ToJSON()

		var transformed map[string]interface{}

		json.Unmarshal([]byte(got), &transformed)

		if err != nil {
			t.Fatalf("Error generating JSON: %v", err)
		}

		want := map[string]interface{}{
			"Name":        "SPS30",
			"Latitude":    51.0,
			"Longitude":   0.0,
			"Measurement": 0.0,
			"Rate":        1,
			"Unit":        "μg/m³",
		}


		// May change this later. Map comparison is quite confusing (reflect.DeepEqual() returns false)
		if !(fmt.Sprint(transformed) == fmt.Sprint(want)) {
			t.Errorf("Unexpected JSON output.\nGot: %v\nWant: %v", transformed, want)
		}

	})
}

func TestClient(t *testing.T) {
	t.Run("Create a Client", func(t *testing.T) {
		client := CreateClient("broker.hivemq.com:1883", "go-publisher", handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	})

}
