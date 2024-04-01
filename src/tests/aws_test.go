package testing

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	appURL := "http://100.26.209.229:3000"

	for i := 0; i < 30; i++ {
		resp, err := http.Get(appURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			fmt.Println("A aplicação está em execução na nuvem!")
			os.Exit(0)
		}

		time.Sleep(time.Second)
	}

	fmt.Println("Falha ao conectar à aplicação na nuvem.")
	os.Exit(1)
}
