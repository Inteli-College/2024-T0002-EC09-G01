---
label: "Instruções para execução do projeto"
---

# Instruções para execução do projeto


## Estutura de pastas

O repositório está organizado da seguinte forma:

```bash
.
├── docs
├── README.md
├── src
│   ├── build
│   ├── cmd
│   ├── config
│   ├── internal
│   ├── pkg
│   ├── tests
├── .github
```

- docs: Diretório que contem todos os arquivos referentes à documentação do projeto.
- src: Contem todos os arquivos relacionados com o código fonte da solução
    - build: Dentro da pasta 'build', encontram-se os arquivos necessários para executar a solução contêinerizada.
    - cmd: Na pasta cmd possuem os arquivos de ponta de entrada, isto é, arquivos que são rodados no terminal para iniciar a solução.
    - config: Em config são armazenados arquivos de configuração da solução.
    - internal: Na pasta internal são armazenados arquivos internos essenciais para a construção da solução.
    - pkg: Na pasta pkg são armazenados os arquivos que constroem a simulação dos sensores.
    - tests: São separados os testes feitos para garantir o funcionamento da solução.
- .github: Contem o fluxo de integração contínua do repositório

### Pré-requisitos
As seguintes ferramentas são necessárias para rodar o projeto sem problemas

- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Go](https://go.dev/doc/install)
- [Docker](https://www.docker.com/)

## Lançando a aplicação

Para utilizar os módulos:

**1. Clone o repositório, se ainda não o tiver feito:**

```bash
git clone https://github.com/Inteli-College/2024-T0002-EC09-G01.git
```

**2. Entre no diretório _src_:**

```bash
cd src
```

**3. Preparando o ambiente:**

Dentro da pasta config crie um arquivo `.env` e adicione informações na seguinte forma:

```bash
BROKER_ADDR = <endereço-do-broker>
HIVE_USER = <user>
HIVE_PSWD = <senha>

CONFLUENT_BOOTSTRAP_SERVER_SASL= <bootstrap-confluent>
CONFLUENT_API_KEY= <api-key-confluent>
CONFLUENT_API_SECRET= <api-secret-confluent>
KAFKA_TOPIC= <nome-do-topico-kafka>

MONGO_USER= <usuario-mongo-db>
MONGO_PASSWORD= <senha-mongo-db>
```
> Importante: Para que a solução funcione conforme o esperado é necesssário realizar uma integração com o Broker [HiveMQ](https://www.hivemq.com/) com a nuvem da [Confluent](https://www.confluent.io/lp/confluent-kafka/?utm_medium=sem&utm_source=google&utm_campaign=ch.sem_br.brand_tp.prs_tgt.confluent-brand_mt.xct_rgn.latam_lng.eng_dv.all_con.confluent-general&utm_term=confluent&creative=&device=c&placement=&gad_source=1&gclid=CjwKCAjw48-vBhBbEiwAzqrZVLLJ07a0teT8LEMGgOZONiDB_GAP2uVqed6ZIrZSOQyQZEjfB_kkmRoCRxgQAvD_BwE). Para isso, siga o seguinte [tutorial](https://www.hivemq.com/blog/harnessing-power-hivemq-cloud-confluent-cloud-mqtt-kafka-for-iot/).

**4. Rodando o simulador de sensores:**

```bash
cd src/cmd/simulation
go run simulation.go
```

**5. Rodando o consumer kafka que receber os dados e registra do banco de dados:**

```bash
cd src/cmd/app
go run app.go
```

**5. Rodando a dashboard Metabase:**

```bash
cd src/build
docker compose up
```

## Testes
Para garantir a funcionalidade da solução, foram criados testes automatizados que aferem a qualidade das funcionalidades. Cada pacote epropostasstá acompanhado de um arquivo de teste que indica o tempo de execução de algumas funções, bem como se os testes foram ou não bem sucedidos.

| Recurso | Testes Desenvolvidos | Status | Nome do Arquivo |
|:-------:|:--------------------:|:------:|:---------------:|
| Publisher | <ul><li>Criação de uma instância de Sensor</li><li>Criação correta de uma carga útil de dados em formato JSON</li></ul> | ✅ Sucesso | publisher_test.go |
| Subscriber | Teste de Subscrição em um tópico | ✅ Sucesso | subscriber_test.go |
| Client | Criação de um cliente com conexão com um Broker | ✅ Sucesso| client_test.go |



### Como rodar os testes

1. Entre no diretório do pacote que deseja testar
2. Rode o comando de teste do Go

Exemplo com o pacote _Publisher_:

```bash
cd src/tests
go test
```

### Observações
- Existem planos para dockerizar a solução como um todo no futuro. 🐋
