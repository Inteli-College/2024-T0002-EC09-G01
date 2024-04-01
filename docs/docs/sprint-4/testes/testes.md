---
label: "Testes"
---

# Testes do sistema 

## api_db_test.go
Nesse código temos um conjunto de testes para um sistema que lida com sensores de diferentes tipos, como sensores de temperatura, sensores de gases e sensores de radiação, utilizando a biblioteca padrão de teste (testing) para verificar se as rotas da API do sistema estão funcionando corretamente.

### Teste de envio de dados do sensor (**Função `testPostsensor`**):
Este teste envia um payload JSON contendo informações sobre um sensor, como latitude, longitude, código, fabricante, etc., para a rota /sensors do servidor. Após enviar os dados, o teste verifica se a resposta do servidor foi bem-sucedida.

### Teste de envio de dados de gás (**Função `testPostgas`**):
Similar ao teste anterior, este teste envia informações sobre a concentração de diferentes gases para a rota /gases do servidor. O payload contém informações como o ID do sensor, nome do sensor, unidade de medida e concentrações de vários gases. Após enviar os dados, o teste verifica se a resposta do servidor foi bem-sucedida.

### Teste de envio de dados de radiação (**Função `testPostradiation`**):
Novamente, este teste envia informações sobre a radiação detectada por um sensor para a rota /radiations do servidor, o payload inclui o ID do sensor, nome do sensor, unidade de medida e nível de radiação detectada. Por fim, o teste então verifica se a resposta do servidor foi bem-sucedida.

### Teste de requisição GET (**Função `testGet`**):
Este teste realiza uma requisição GET para as três rotas principais do servidor (/sensors, /gases e /radiations), verifica se o servidor responde corretamente às requisições GET para cada uma dessas rotas.

### Teste principal (**Função `TestApi`**):
Este é o teste principal que chama os testes individuais acima, ele executa todos os testes em sequência para verificar se todas as funcionalidades principais da API estão funcionando corretamente.

## aws_test.go
Tal teste verifica se uma aplicação está em execução em uma URL específica na nuvem, realizando uma série de tentativas de conexão, fazendo uma requisição HTTP GET para a URL da aplicação a cada tentativa. 

Se a conexão for bem-sucedida e a aplicação responder com o código de status 200 (OK), o programa imprime "A aplicação está em execução na nuvem!" e encerra com um código de saída 0, indicando sucesso. Caso contrário, ele aguarda um segundo e tenta novamente; após 30 tentativas sem sucesso, o programa imprime "Falha ao conectar à aplicação na nuvem." e encerra com um código de saída 1, indicando falha na conexão.

## client_test.go
Este código contém testes para um cliente MQTT, utilizando a biblioteca "github.com/eclipse/paho.mqtt.golang". 

### Teste de Criação de Cliente (**Função `CreateClient`**):
Teste que cria um cliente MQTT utilizando uma função de criação fornecida pelo pacote comum (common), verificando se a conexão do cliente é bem-sucedida.

### Teste de Inscrição em Tópico (**Função `SubscribeTopic`**):
Neste teste, o cliente MQTT é criado e conectado. Em seguida, ele se inscreve em um tópico específico e espera receber mensagens nesse tópico; após dois segundos, o cliente cancela a inscrição no tópico.

### Teste de Publicação de Mensagem (**Função `PublishMessage`**):
Cria um cliente MQTT, o conecta e publica uma mensagem em um tópico específico, verificando se a publicação da mensagem é bem-sucedida.

## controller_test.go
Testes para validar o comportamento do controlador de um sistema que lida com mensagens MQTT. 

### Teste de Publicação de Campos (**Função `TestPublishFields`**):
Verifica se os campos das mensagens MQTT publicadas estão de acordo com um padrão específico. Sendo assim, o controlador publica mensagens nos tópicos "sensor/gases" e "sensor/radiation". Para cada mensagem publicada, o teste verifica se o payload da mensagem corresponde ao padrão esperado para cada tópico. Se o payload não corresponder ao padrão esperado, o teste falhará, indicando uma inconsistência nos campos da mensagem.

### Teste de Qualidade de Serviço (QoS) (**Função `TestQos`**):
Tal teste verifica se o serviço de qualidade de serviço (QoS) está configurado corretamente. O controlador se inscreve em todos os tópicos sob "sensor/#", após uma pausa de dois segundos, o controlador verifica se todas as mensagens recebidas têm QoS igual a 1. Se alguma mensagem não tiver QoS igual a 1, o teste falhará, indicando uma configuração incorreta do QoS.

## database_test.go
Teste de unidade para verificar a funcionalidade de inserção de dados no MongoDB.

### Teste de Inserção no MongoDB (**Função `TestInsertIntoMongo`**):
Verifica se os dados são corretamente inseridos no MongoDB, criando um cliente MongoDB para conectar-se ao banco de dados utilizando a função setupMongoClient. Em seguida, define um conjunto de dados simulados para serem inseridos no banco de dados, com o auxílio da função mongo.InsertIntoMongo. Se ocorrer algum erro durante a inserção, o teste falhará, exibindo uma mensagem de erro.

### Configuração do Cliente MongoDB (**Função `setupMongoClient`**):
Tal função auxilia configuração e retorna um cliente MongoDB para ser usado nos testes, utilizando as opções do cliente para especificar a versão da API do servidor. Se ocorrer algum erro durante a conexão com o MongoDB, o teste falhará, exibindo uma mensagem de erro.

## kafka_test.go
O teste unitário é reponsável por verificar a funcionalidade de consumo de mensagens do Kafka.

No código, um consumidor Kafka é configurado para se inscrever em um tópico específico e receber mensagens desse tópico, utilizando um canal para receber as mensagens, uma goroutine é iniciada para processar os eventos do Kafka e enviar as mensagens recebidas para o canal. O teste então envia três mensagens de teste para o canal e verifica se o consumidor as recebe corretamente.

# Executando os testes
Antes de tudo, garanta que os requisitos e aplicações externas estão em funcionamento e acesse a pastas `/src/testes/`:

```
go test -v
```

Caso seja um teste específico:
```
go test -v nome-do-teste.go
```



