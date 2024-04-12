---
label: "Backend"
---

# Introdução

A fim de o usuário administrador pudesse criar nosso sensores e que os cidadãos pudessem informar de situações para que outros fiquem alerta, foi desenvolvida um aplicação web. Esta seção é dedicada a parte back-end da aplicação e foi organizada em arquivos de código, que respectivamente são e fazem:

## mongodb.go

### Funcionalidades:
1. **Conexão com o MongoDB:**
   - `ConectarAoMongo()`: Estabelece uma conexão com o MongoDB utilizando credenciais armazenadas em um arquivo `.env`.

2. **Operações de Banco de Dados:**
   - `ObterDoMongo(cliente *mongo.Client) []Sensor`: Realiza uma consulta na coleção "sensores" e retorna uma lista de sensores.
   - `InserirNoMongo(cliente *mongo.Client, dados interface{}, colecao string) string`: Insere um documento em uma coleção específica.

## items.go

### Estruturas de Dados:
1. `Sensor`: Estrutura que define os atributos de um sensor. Dividido em:
- Nome do sensor
- Tipo do sensor
- Latitude
- Longitude
2. `Alerta`: Estrutura que define os atributos de um alerta.  Dividido em:
- Nome do alerta
- Tipo do sensor
- Latitude
- Longitude

## api.go

### Funcionalidades:

1. **Configuração do Servidor HTTP:**
   - Configura o roteador Gin.
   - Habilita o CORS para permitir solicitações de origens múltiplas.

2. **Rotas da API:**
   - `GET /sensores`: Retorna todos os sensores.
   - `POST /sensores`: Cria um novo sensor.
   - `GET /alertas`: Retorna todos os alertas.
   - `POST /alertas`: Cria um novo alerta.

3. **Funções de Manipulação de Solicitações HTTP:**
   - `criarSensor(c *gin.Context)`: Manipula a criação de um novo sensor.
   - `obterSensores(c *gin.Context)`: Manipula a obtenção de todos os sensores.
   - `criarAlerta(c *gin.Context)`: Manipula a criação de um novo alerta.
   - `obterAlertas(c *gin.Context)`: Manipula a obtenção de todos os alertas.

Esta documentação fornece uma visão geral das funcionalidades e estruturas de dados presentes nos arquivos, sendo útil para compreensão e manutenção do código-fonte.


## Lançando a aplicação
Partindo a partir da pasta raiz:
```go
run src/backend/api.go
```