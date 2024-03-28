# Backend para Recebimento de Dados de Sensores e Alertas

Este é um backend construído para receber requisições de dados de sensores e alertas e adicioná-los no banco de dados MongoDB. Ele fornece endpoints para enviar dados de sensores e alertas, bem como recuperar esses dados do banco de dados.

## Tecnologias Utilizadas

    Gin: Um framework web para Golang que ajuda a criar APIs de forma rápida e simples.
    MongoDB: Um banco de dados NoSQL altamente escalável e flexível.
    MongoDB Go Driver: Pacote oficial do MongoDB para Go, utilizado para interagir com o banco de dados MongoDB.

## Instalação e Configuração

Antes de executar o backend, você precisa configurar o arquivo `.env` no diretório `/config` com as seguintes variáveis:

    MONGO_USER=seu_usuario_mongodb
    MONGO_PASSWORD=sua_senha_mongodb

Além disso, é necessário ter o MongoDB instalado em sua máquina ou utilizar um serviço de hospedagem do MongoDB, como o MongoDB Atlas.

## Endpoints

1. Adicionar Dados de Sensor

        Método: POST
        Endpoint: /sensors
        Corpo da Requisição: JSON contendo os dados do sensor

        json

        {
        "sensor": "nome_do_sensor",
        "tipo": "tipo_do_sensor",
        "longitude": "longitude_do_sensor",
        "latitude": "latitude_do_sensor"
        }

    Resposta de Sucesso: Retorna os dados do sensor adicionados no formato JSON com status HTTP 201 (Created).

2. Obter Dados de Sensores

        Método: GET
        Endpoint: /sensors
        
    Resposta de Sucesso: Retorna uma lista de todos os sensores cadastrados no formato JSON com status HTTP 200 (OK).

3. Adicionar Alerta

        Método: POST
        Endpoint: /alerts
        Corpo da Requisição: JSON contendo os dados do alerta

        json

        {
        "alert": "nome_do_alerta",
        "tipo": "tipo_do_alerta",
        "longitude": "longitude_do_alerta",
        "latitude": "latitude_do_alerta"
        }

    Resposta de Sucesso: Retorna os dados do alerta adicionados no formato JSON com status HTTP 201 (Created).

4. Obter Alertas

        Método: GET
        Endpoint: /alerts

    Resposta de Sucesso: Retorna uma lista de todos os alertas cadastrados no formato JSON com status HTTP 200 (OK).

## Executando o Servidor

Após configurar o arquivo .env, execute o seguinte comando no diretório `/src`:

    source .bashrc
    cd backend
    go run .

O servidor será iniciado na porta 8000 por padrão.


## Observações

Certifique-se de que o serviço do MongoDB está em execução e acessível.