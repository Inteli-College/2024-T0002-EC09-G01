package mics6814

import (
	"math"
	"math/rand"
	"time"
)

type GasesValues struct {
	CarbonMonoxide  float64 `json:"carbon_monoxide"`
	NitrogenDioxide float64 `json:"nitrogen_dioxide"`
	Ethanol         float64 `json:"ethanol"`
	Hydrogen        float64 `json:"hydrogen"`
	Ammonia         float64 `json:"ammonia"`
	Methane         float64 `json:"methane"`
	Propane         float64 `json:"propane"`
	IsoButane       float64 `json:"iso_butane"`
}

type SensorConfig struct {
	Sensor          string  `json:"sensor"`
	Unit 		  string  `json:"unit"`
	GasesValues	 GasesValues `json:"gases-values"`
}

type MaxMin struct {
	MaxValue float64 `json:"max_value"`
	MinValue float64 `json:"min_value"`
}

var gasesRange = map[string]MaxMin{
	"carbon_monoxide":  {1, 1000},
	"nitrogen_dioxide": {0.05, 10},
	"ethanol":          {10, 500},
	"hydrogen":         {1, 1000},
	"ammonia":          {1, 500},
	"methane":          {1001, 9999}, // ">1000 ppm"
	"propane":          {1001, 9999}, // ">1000 ppm"
	"iso_butane":       {1001, 9999}, // ">1000 ppm"
}

func RandomValues(gas string) float64 {
	rand.Seed(time.Now().UnixNano()) // Inicializa a semente do gerador de números aleatórios

	maxValue := gasesRange[gas].MaxValue
	minValue := gasesRange[gas].MinValue
	value := rand.Float64()*(maxValue-minValue) + minValue

	return math.Round(value*100) / 100
}

func CreateGasesValues() SensorConfig {
	gasesData := GasesValues{
		CarbonMonoxide:  RandomValues("carbon_monoxide"),
		NitrogenDioxide: RandomValues("nitrogen_dioxide"),
		Ethanol:         RandomValues("ethanol"),
		Hydrogen:        RandomValues("hydrogen"),
		Ammonia:         RandomValues("ammonia"),
		Methane:         RandomValues("methane"),
		Propane:         RandomValues("propane"),
		IsoButane:       RandomValues("iso_butane"),
	}
	sensorData := SensorConfig{
		Sensor: "MiCS-6814",
		Unit: "ppm",
		GasesValues: gasesData,
	}	
	return sensorData
}

// func main() {
// 	data := CreateGasesValues()
// 	fmt.Print(data)
// }
