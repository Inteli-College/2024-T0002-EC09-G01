package main

import (
	Mongo "2024-T0002-EC09-G01/src/internal/mongo"
	Controller "2024-T0002-EC09-G01/src/pkg/controller"
	"log"
	"sync"
)

func main() {

	sensors, err := Mongo.GetAllSensors("../../config/.env")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, sensor := range sensors {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			Controller.Controller(id)
		}(sensor.ID)
	}

	wg.Wait()
}
