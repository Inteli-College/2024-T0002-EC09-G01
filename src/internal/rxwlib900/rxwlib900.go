package rxwlib900

import (
	"math"
	"math/rand"
	"time"
)

type RadiationValues struct {
	Radiation float64 `json:"radiation"`
}

type SensorConfig struct {
	Sensor          string          `json:"sensor"`
	Unit            string          `json:"unit"`
	RadiationValues RadiationValues `json:"radiation-values"`
}

type MaxMin struct {
	MaxValue float64 `json:"max_value"`
	MinValue float64 `json:"min_value"`
}

var gasesRange = map[string]MaxMin{
	"radiation": {1, 1280},
}

func RandomValues(gas string) float64 {
	rand.Seed(time.Now().UnixNano()) // Inicializa a semente do gerador de números aleatórios

	maxValue := gasesRange[gas].MaxValue
	minValue := gasesRange[gas].MinValue
	value := rand.Float64()*(maxValue-minValue) + minValue

	return math.Round(value*100) / 100
}

func CreateGasesValues() SensorConfig {
	radiationData := RadiationValues{
		Radiation: RandomValues("radiation"),
	}
	sensorData := SensorConfig{
		Sensor:          "RXWLIB900",
		Unit:            "W/m2",
		RadiationValues: radiationData,
	}
	return sensorData
}

// func main() {
// 	fmt.Println(CreateGasesValues())
// }
