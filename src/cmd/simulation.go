package main

import (
	// MICS6814 "2024-T0002-EC09-G01/src/pkg/pub_mics6814"
	// RXWLIB900 "2024-T0002-EC09-G01/src/pkg/pub_rxwlib900"
	Controller "2024-T0002-EC09-G01/src/pkg/controller"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numControllersRad := 5 // ou o número desejado de controladores

	for i := 1; i <= numControllersRad; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Controller.Controller(id)
		}(i)
	}

	wg.Wait()
}
