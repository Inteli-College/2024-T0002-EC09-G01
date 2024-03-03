---
label: "Criação de dados simulados"
---

# Criação de dados simulados


Para a execução do projeto demanda-se a criação de um simulador para dispositivos IoT. Esse simulador é capaz de enviar informações para tópicos em um formato de dados consistente, semelhante aos dados captados por sensores reais. Até o momento, os sensores simulados no estado atual do projeto incluem o [Mics-6814](https://datasheetspdf.com/product/1350171/SGX/MiCS-6814/index.html) e o [RXW-LIB-900](https://sigmasensors.com.br/produtos/).

Para viabilizar essa simulação, foi desenvolvida uma função especializada em gerar dados aleatórios, alinhados com as faixas de alcance dos sensores reais. A implementação dessa função pode ser encontrada no diretório `2024-T0002-EC02-G01/src/internal/common`.

```
type MaxMin struct {
	MaxValue float64 `json:"max_value"`
	MinValue float64 `json:"min_value"`
}

func RandomValues(rangeValues map[string]MaxMin, gas string) float64 {
	rand.Seed(time.Now().UnixNano()) 

	maxValue := rangeValues[gas].MaxValue
	minValue := rangeValues[gas].MinValue
	value := rand.Float64()*(maxValue-minValue) + minValue

	return math.Round(value*100) / 100
}
```

Para garantir um desempenho ideal, a função requer a entrada de um mapa contendo os valores máximos e mínimos capturados pelo sensor. Os intervalos para a geração de valores foram escolhidos com base nas especificações fornecidas nos datasheets dos respectivos sensores. Isso assegura que os dados gerados estejam alinhados com as características esperadas dos sensores em questão.

- Mics-6814

```
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
```

- RXW-LIB-900

```
var radiationRange = map[string]Common.MaxMin{
	"radiation": {1, 1280},
}
```



