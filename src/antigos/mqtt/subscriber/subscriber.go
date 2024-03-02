package main

import (
	"fmt"
	DefaultClient "2024-T0002-EC09-G01/src/antigos/mqtt/common"
)

func main() {
	
	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DefaultClient.Handler)
	
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("sensors/SPS30", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber running...")
	select {}
}