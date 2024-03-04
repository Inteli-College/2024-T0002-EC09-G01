package testing

import (
	"io"
	"net/http"
	"testing"
	"time"
	"bytes"
	"encoding/json"
)

const url string = "http://localhost:8000/"

func testPostsensor(t *testing.T) {
	endpoint := url + "sensors"

	message := map[string]interface{}{
		"latitude": -23.555734690062174,
		"longitude": -46.73388952459531,
		"sensor": "Sensor-Inteli",
		"code": 123,
		"manufacturer": "SmarTopia",
		"author": "Luana Parra",
		"date": time.Now(),
	}

	payload, err := json.Marshal(message)

	if err != nil {
		t.Errorf("Error marshalling message: %v", err)
	}

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("Error making POST request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	t.Logf("Response body: %s", body)
}

func testPostgas(t *testing.T) {
	endpoint := url + "gases"

	message := map[string]interface{}{
		"sensorId": 123,
		"sensorName": "Sensor-Inteli",
		"unit": "ppm",
		"time": time.Now(),
		"NH3": 0.5,
		"CO": 0.5,
		"C2H5OH": 0.5,
		"H2": 0.5,
		"iC4H10": 0.5,
		"CH4": 0.5,
		"NO2": 0.5,
		"C3H8": 0.5,
	}

	payload, err := json.Marshal(message)
	if err != nil {
		t.Errorf("Error marshalling message: %v", err)
	}

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("Error making POST request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	t.Logf("Response body: %s", body)
}

func testPostradiation(t *testing.T) {
	endpoint := url + "radiations"

	message := map[string]interface{}{
		"sensorId": 123,
		"sensorName": "Sensor-Inteli",
		"unit": "uSv/h",
		"time": time.Now(),
		"radiation": 0.5,
	}

	payload, err := json.Marshal(message)
	if err != nil {
		t.Errorf("Error marshalling message: %v", err)
	}

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("Error making POST request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	t.Logf("Response body: %s", body)
}


func testGet(t *testing.T) {
	sensor_endpoint := url + "sensors"
	gas_endpoint := url + "gases"
	radiation_endpoint := url + "radiations"

	endpoints := []string{sensor_endpoint, gas_endpoint, radiation_endpoint}

	for _, endpoint := range endpoints {
		response, err := http.Get(endpoint)
		if err != nil {
			t.Errorf("Error making GET request: %v", err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Errorf("Error reading response body: %v", err)
		}

		t.Logf("Response body: %s", body)
	}
}

func TestApi(t *testing.T) {
	testPostsensor(t)
	testPostgas(t)
	testPostradiation(t)
	testGet(t)
}