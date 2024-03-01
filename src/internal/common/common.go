package common

import (
	"math"
	"math/rand"
	"time"
)

type MaxMin struct {
	MaxValue float64 `json:"max_value"`
	MinValue float64 `json:"min_value"`
}

func RandomValues(rangeValues map[string]MaxMin, gas string) float64 {
	rand.Seed(time.Now().UnixNano()) // Inicializa a semente do gerador de números aleatórios

	maxValue := rangeValues[gas].MaxValue
	minValue := rangeValues[gas].MinValue
	value := rand.Float64()*(maxValue-minValue) + minValue

	return math.Round(value*100) / 100
}
