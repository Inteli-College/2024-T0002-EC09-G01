---
label: "Integração do Kafka e Hive MQ"
---

# Introdução
Este tópico aborda a integração entre Apache Kafka e Hive MQ na nossa solução, principalmente em relação a coleta, processamento e armazenamento dos dados provenientes de sensores. 

# Kafka e Hive MQ
## Hive MQ
Como já tratado na Sprint 2, o Hive MQ é um broker de mensageria MQTT altamente escalável e de baixa latência, facilitando a comunicação entre dispositivos conectados.

## Kafka 
O Apache Kafka é uma plataforma de streaming distribuída, utilizada para construir sistemas em tempo real e processamento de eventos. Tal tecnologia permite a publicação e subscrição de streams de registros, semelhante a um sistema de mensagens tradicional, mas com recursos adicionais, como a capacidade de armazenamento de dados por um período configurável.

O Kafka, em geral, oferece uma arquitetura escalável e de alta disponibilidade, abaixo explicamos de maneira resumida os principais conceitos:

- **Tópicos**: canais de comunicação onde os dados são publicados e consumidos.
- **Produtores**: responsáveis por enviar dados para os tópicos.
- **Consumidores**: aplicativos ou sistemas que recebem e processam os dados dos tópicos.
- **Partições**: os tópicos são divididos em partições para permitir a distribuição e o paralelismo no processamento.

No contexto do projeto, preferimos utilizar o Confluent Cloud que desempenha o papel fundamental na ingestão, armazenamento e processamento dos dados provenientes dos sensores.

Dessa maneira, os dados coletados pelos sensores são enviados para o Confluent Cloud por meio de produtores Kafka, garantindo com que os dados sejam armazenados de forma durável e distribuída em tópicos Kafka. 

Além disso, o Confluent fornece conectores Kafka para integração fácil com o Metabase ou outras ferramentas de visualização de dados, fazendo com que os dados processados sejam enviados para o Metabase criando o nosso dashboard.

# Integração Kafka e Hive MQ
A integração entre Apache Kafka e Hive MQ desempenha um papel crucial no fluxo de dados provenientes de sensores em um ambiente de IoT. Assim, essa integração possibilita a coleta eficiente de dados em tempo real, sua entrega confiável ao Kafka para processamento e posterior análise.

## Configuração do Consumidor Kafka 
No script '', é possível observar a configuração de um consumidor Kafka para receber mensagens do Hive MQ e encaminhá-las para o Kafka.

```
configMap := &ckafka.ConfigMap{
    "bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
    "sasl.mechanisms":    "PLAIN",
    "security.protocol":  "SASL_SSL",
    "sasl.username":      os.Getenv("CONFLUENT_API_KEY"),
    "sasl.password":      os.Getenv("CONFLUENT_API_SECRET"),
    "session.timeout.ms": 6000,
    "group.id":           "manu",
    "auto.offset.reset":  "latest",
}
```

## Consumo de Mensagens
Após a configuração, o consumidor Kafka é utilizado para receber mensagens do tópico e processá-las conforme necessário.

```
for msg := range msgChan {
    log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

    var result map[string]interface{}
    err := json.Unmarshal(msg.Value, &result)
    if err != nil {
        log.Fatal(err)
    }

    // Aqui você pode processar os dados conforme necessário.
}
```

## Fluxo de dados
O fluxo de dados inicia-se com a coleta de informações pelos sensores, os quais podem capturar uma ampla gama de dados, como temperatura, umidade, pressão, entre outros. 

Dessa maneira, o Hive MQ atua como o broker MQTT responsável por receber os dados dos sensores, permitindo com que os dispositivos IoT publiquem dados em tópicos específicos, simplificando a comunicação entre os dispositivos e o broker, ou seja, os sensores coletam dados e os enviam para o Hive MQ.

Uma vez recebidos pelo Hive MQ, os dados são encaminhados para o Kafka por meio de um consumidor Kafka configurado para consumir mensagens do Hive MQ, tal consumidor utiliza a biblioteca `confluent-kafka-go` para se conectar ao Hive MQ e receber as mensagens.

Por fim, os dados são armazenados temporariamente e processados, para poderem ser guardados no banco de dados e exibidos no dashboard.