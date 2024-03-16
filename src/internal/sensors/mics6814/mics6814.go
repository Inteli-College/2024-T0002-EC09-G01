package mics6814

import (
	Common "2024-T0002-EC09-G01/src/internal/sensors/common"
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

var gasesRange = map[string]Common.MaxMin{
	"carbon_monoxide":  {1, 1000},
	"nitrogen_dioxide": {0.05, 10},
	"ethanol":          {10, 500},
	"hydrogen":         {1, 1000},
	"ammonia":          {1, 500},
	"methane":          {1001, 9999}, // ">1000 ppm"
	"propane":          {1001, 9999}, // ">1000 ppm"
	"iso_butane":       {1001, 9999}, // ">1000 ppm"
}

func CreateGasesValues() SensorConfig {
	gasesData := GasesValues{
		CarbonMonoxide:  Common.RandomValues(gasesRange, "carbon_monoxide"),
		NitrogenDioxide: Common.RandomValues(gasesRange, "nitrogen_dioxide"),
		Ethanol:         Common.RandomValues(gasesRange, "ethanol"),
		Hydrogen:        Common.RandomValues(gasesRange, "hydrogen"),
		Ammonia:         Common.RandomValues(gasesRange, "ammonia"),
		Methane:         Common.RandomValues(gasesRange, "methane"),
		Propane:         Common.RandomValues(gasesRange, "propane"),
		IsoButane:       Common.RandomValues(gasesRange, "iso_butane"),
	}
	sensorData := SensorConfig{
		Sensor: "MiCS-6814",
		Unit: "ppm",
		GasesValues: gasesData,
	}	
	return sensorData
}
