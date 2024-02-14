---
label: "Simulador MQTT"
---

# Simulador MQTT

<img title="MQTT Example" alt="Exemplo de arquitetura MQTT, apresentando os componentes principais" src={require('/img/mqtt-example.png').default} />

O MQTT (Message Queuing Telemetry Transport) é um protocolo leve de mensagens projetado para dispositivos com recursos limitados e redes instáveis. Ele usa o modelo de publicação/assinatura, onde os clientes se comunicam através de um intermediário (broker). Mensagens são enviadas com tópicos, permitindo que os clientes assinem apenas as mensagens de interesse. Isso o torna ideal para IoT, onde a eficiência e a escalabilidade são essenciais.

O simulador MQTT tem como propósito reproduzir a dinâmica da comunicação entre sensores e dispositivos do parceiro. Para isso, foram estabelecidas estruturas de dados condizentes com as especificações dos sensores, visando simular casos de uso reais.

O código fonte da solução pode ser encontrado no [repositório do grupo no Github](https://github.com/Inteli-College/2024-T0002-EC09-G01/tree/main)!

## Publisher

Para facilitar o teste com vários sensores simulados, abstraímos uma estrutura de dados que contém as principais informações de um sensor, como nome, latitude, longitude, medição, frequência, e unidade de medida. Isso simplifica a criação de novas instâncias para testes.

```go
type Sensor struct {
	Name        string
	Latitude    float64
	Longitude   float64
	Measurement float64
	Rate        int
	Unit        string
}
```

A comunicação é então iniciada com um broker público (nesse exemplo, optamos por utilizar o hivemq), e as leituras do sensor são enviadas. Note que são gerados valores aleatórios entre 0.03 e 1 para simular uma amplitude de leitura do sensor para uma determinada grandeza:

```go
for {
		for _, sensor := range sensors {
			
			topic := "sensors/" + sensor.Name

			sensor.Measurement = (rand.Float64() * (maxSensorRange - minSensorRange)) + minSensorRange

			payload, _ := sensor.ToJSON()

			token := client.Publish(topic, 0, false, payload)

			token.Wait()

			fmt.Printf("Published message: %s\n", payload)

			time.Sleep(time.Duration(sensor.Rate) * time.Second)

		}
	}
```

## Subscriber

Dispositivos Subscribers recebem informações ao se inscrever em tópicos. Optamos por manter um exemplo simples e conciso neste primeiro momento, considerando que o subscriber pode ser uma interface genérica. Abstrações adicionais podem ser aplicadas posteriormente.

```go
opts := MQTT.NewClientOptions().AddBroker(broker)
opts.SetClientID("go_subscriber")
opts.SetDefaultPublishHandler(messagePubHandler)

client := MQTT.NewClient(opts)
if token := client.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
}

if token := client.Subscribe("sensors/Sensor1", 1, nil); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    return
}

fmt.Println("Subscriber running...")
select {}
```