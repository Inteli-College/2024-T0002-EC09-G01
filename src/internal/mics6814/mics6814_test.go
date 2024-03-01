package mics6814

import(
	"testing"
)

func TestCreateGasesValues(t *testing.T) {
	data := CreateGasesValues()
	if data.GasesValues.CarbonMonoxide > gasesRange["carbon_monoxide"].MinValue || data.GasesValues.CarbonMonoxide < gasesRange["carbon_monoxide"].MaxValue {
		t.Errorf("Valor de CO fora do intervalo esperado")
	}
	if data.GasesValues.NitrogenDioxide > gasesRange["nitrogen_dioxide"].MinValue || data.GasesValues.NitrogenDioxide < gasesRange["nitrogen_dioxide"].MaxValue {
		t.Errorf("Valor de NO2 fora do intervalo esperado")
	}
	if data.GasesValues.Ethanol > gasesRange["ethanol"].MinValue || data.GasesValues.Ethanol < gasesRange["ethanol"].MaxValue {
		t.Errorf("Valor de EtOH fora do intervalo esperado")
	}
	if data.GasesValues.Hydrogen > gasesRange["hydrogen"].MinValue || data.GasesValues.Hydrogen < gasesRange["hydrogen"].MaxValue {
		t.Errorf("Valor de H2 fora do intervalo esperado")
	}
	if data.GasesValues.Ammonia > gasesRange["ammonia"].MinValue || data.GasesValues.Ammonia < gasesRange["ammonia"].MaxValue {
		t.Errorf("Valor de NH3 fora do intervalo esperado")
	}
	if data.GasesValues.Methane > gasesRange["methane"].MinValue || data.GasesValues.Methane < gasesRange["methane"].MaxValue {
		t.Errorf("Valor de CH4 fora do intervalo esperado")
	}
	if data.GasesValues.Propane > gasesRange["propane"].MinValue || data.GasesValues.Propane < gasesRange["propane"].MaxValue {
		t.Errorf("Valor de C3H8 fora do intervalo esperado")
	}
	if data.GasesValues.IsoButane > gasesRange["iso_butane"].MinValue || data.GasesValues.IsoButane < gasesRange["iso_butane"].MaxValue {
		t.Errorf("Valor de i-C4H10 fora do intervalo esperado")
	}
}