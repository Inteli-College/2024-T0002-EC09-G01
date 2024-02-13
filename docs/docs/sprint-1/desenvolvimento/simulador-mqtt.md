---
label: "Simulador MQTT"
---

# Simulador MQTT

<img title="MQTT Example" alt="Exemplo de arquitetura MQTT, apresentando os componentes principais" src={require('/img/simulador-mqtt').default} >

O MQTT (Message Queuing Telemetry Transport) é um protocolo leve de mensagens projetado para dispositivos com recursos limitados e redes instáveis. Ele usa o modelo de publicação/assinatura, onde os clientes se comunicam através de um intermediário (broker). Mensagens são enviadas com tópicos, permitindo que os clientes assinem apenas as mensagens de interesse. Isso o torna ideal para IoT, onde a eficiência e a escalabilidade são essenciais.

O simulador MQTT tem como finalidade reproduzir de forma inicial a dinâmica da comunicação entre sensores e dispositivos do parceiro. Para isso, foram estabelecidas estruturas de dados condizentes com as especificações dos sensores de forma a de fato simular um caso de uso da vida real. Utilizou-se o a biblioteca _paho-mqtt_ para constituir a principal ferramenta de comunicação MQTT.

O código fonte da solução pode ser encontrado no [repositório do grupo no Github](https://github.com/Inteli-College/2024-T0002-EC09-G01/tree/main)!

## Publisher

Com a finalidade de facilitar o teste com diversos sensores simulados, abstraímos uma estrutura de dados que comportasse as principais informações que compõem um Sensor, que neste caso, incluem _nome_, _latitude_, _longitude_, _medição_ e _frequência_, bem como o atributo para permtir a conexão com o broker. Dessa forma, a instancia de novas unidades para finalidade de testes é facilitada.

```go
type Sensor struct {
	name        string
	latitude    float64
	longitude   float64
	measurement float64
	rate        int
	mqtt.Client
}
```

A comunicação é então iniciada com um broker público (nesse exemplo, optamos por utilizar o hivemq), e as leituras do sensor são enviadas. Note que são gerados valores aleatórios entre 1 e 5 para simular uma amplitude de leitura do sensor para uma determinada grandeza:

```go
topic := "sensors/" + sensor.name

for {
    sensor.measurement = rand.Float64()*5
    payload := strconv.FormatFloat(sensor.measurement, 'f', 2, 64)
    token := sensor.Publish(topic, 0, false, payload)
    token.Wait()
    fmt.Printf("Published message: %s\n", payload)
    time.Sleep(time.Duration(sensor.rate) * time.Second)
}
```

## Subscriber

Dispositivos Subscribers recebem informações ao se inscrever em tópicos. Pensando que, num primeiro momento, o subscriber será uma inerface qualquer, optamos por manter um exemplo simples e conciso, para depois aplicar as devidas abstrações.

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