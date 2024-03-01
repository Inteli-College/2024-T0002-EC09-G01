package rxwlib900

import(
	"testing"
)

func TestCreateGasesValues(t *testing.T) {
	data := CreateGasesValues()
	if data.RadiationValues.Radiation > radiationRange["radiation"].MinValue || data.RadiationValues.Radiation < radiationRange["radiation"].MaxValue {
		t.Errorf("Valor de radiação fora do intervalo esperado")
	}
}