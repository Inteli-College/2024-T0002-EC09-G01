package main

import (
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
