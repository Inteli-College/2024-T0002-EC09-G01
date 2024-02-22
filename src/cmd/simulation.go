package main

import (
	"sync"
	"2024-T0002-EC09-G01/src/pkg/controller"

)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			controller.Controller(i + 1)
		}()
	}
	wg.Wait()
}