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
├── .github
└── tests
```

- docs: Diretório que contem todos os arquivos referentes à documentação do projeto.
- src: Contem todos os arquivos relacionados com o código fonte da solução
- tests: Armazena testes automatizados da solução
- .github: Contem o fluxo de integração contínua do repositório

### Pré-requisitos
As seguintes ferramentas são necessárias para rodar o projeto sem problemas

- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Go](https://go.dev/doc/install)

## Lançando a aplicação

Para utilizar os módulos:

1. Clone o repositório, se ainda não o tiver feito:

```bash
    git clone https://github.com/Inteli-College/2024-T0002-EC09-G01.git
```

2. Entre no diretório _src/mqtt_

```bash
    cd src/mqtt
```

3. Entre no diretório do pacote que deseja rodar:

Exemplo com o pacote _Subscriber_:
   
```bash
cd subscriber
go run .
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
cd src/mqtt
go test
```

### Observações
- Existem planos para dockerizar a solução no futuro. 🐋
- Apenas a execução da solução em Go foi demonstrada, uma vez que pretendemos seguir com o desenvolvimento utilizando esta ferramenta, apesar de existirem códigos em Python.
