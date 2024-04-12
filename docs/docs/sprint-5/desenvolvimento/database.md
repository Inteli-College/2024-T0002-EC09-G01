---
label: "Database"
---

# Database 

## Banco de dados relacional - SQLite

Um banco de dados relacional é um tipo de banco de dados que armazena e fornece acesso a pontos de dados relacionados entre si. Bancos de dados relacionais são baseados no modelo relacional, uma maneira intuitiva e direta de representar dados em tabelas. Em um banco de dados relacional, cada linha na tabela é um registro com uma ID exclusiva chamada chave. As colunas da tabela contêm atributos dos dados e cada registro geralmente tem um valor para cada atributo, facilitando o estabelecimento das relações entre os pontos de dados.[^1]

Pensando no modelo de desenvolvimento em espiral, optamos pela implementação de uma database SQLite pela facilidade de implementação. Em um primeiro momento, tivemos a projeção de três tabelas que indicam:

- **sensors**: Os sensores cadastrados.
- **radiation**: Leituras dos sensores de radiação.
- **gases**: Leituras dos sensores de gases.


<img title="Table Schema" alt="Imagem representando as tabelas contidas no banco de dados" src={require('/img/Table-Schema.png').default} />

## Tabelas

A descrição das tabelas pode ser definida da seguinte forma:

### sensors
- **id**: Número de registro do sensor. Serve como chava estrangeira para outras tabelas.
- **latitude**: Latitude do sensor.
- **longitude**: Longitude do sensor.
- **sensor_name**: Nome do sensor.
- **sensor_code**: Código serial do sensor.
- **manufacturer**: Nome do fabricante do sensor.
- **registered_by**: Pessoa responsável pelo registro.
- **registration_date**: Data de registro do sensor na base.

### radiation
- **id**: Número de registro da leitura do sensor.
- **sensorId**: Número de registro do sensor que fez a leitura.
- **sensorName**: Nome do sensor
- **unit**: Unidade de medida do sensor
- **time**: Momento em que a leitura foi registrada
- **radiation**: Leitura do valor da radiação na unidade registrada

### gases
- **id**: Número de registro da leitura do sensor.
- **sensorId**: Número de registro do sensor que fez a leitura.
- **sensorName**: Nome do sensor
- **unit**: Unidade de medida do sensor
- **time**: Momento em que a leitura foi registrada
- **NH3**: Leitura do valor da concentração de amônia na unidade registrada
- **CO**: Leitura do valor da concentração de monóxido de carbono (CO) na unidade registrada
- **C2H5OH**: Leitura do valor da concentração de etanol (C2H5OH) na unidade registrada
- **H2**: Leitura do valor da concentração de hidrogênio (H2) na unidade registrada
- **iC4H10**: Leitura do valor da concentração de iso-butano (iC4H10) na unidade registrada
- **CH4**: Leitura do valor da concentração de metano (iC4H10) na unidade registrada
- **NO2**: Leitura do valor da concentração de dióxido de nitrogênio (NO2) na unidade registrada
- **C3H8**: Leitura do valor da concentração de propano (C3H8) na unidade registrada

## Exemplo de representação no metabase

Utilizando uma ferramenta de Bussiness Inteligence (BI), podemos visualizar dados de uma forma que facilite a obtenção de insights para melhorar tomada de decisões no contexto de Big Data. Podemos exemplificar a visualização dos dados nas imagens a seguir:

<img title="Representação de dados no Metabase" alt="Imagem representando os dados da tabela gases no metabase" src={require('/img/sensors-metabase.png').default} />

<img title="Representação de dados no Metabase" alt="Imagem representando os dados da tabela gases no metabase" src={require('/img/gases-values-metabase.png').default} />

## Próxima Sprint - Database NoSQL

No âmbito das cidades inteligentes e suas aplicações em larga escala, a arquitetura com banco de dados relacional enfrenta limitações de escalabilidade devido ao volume massivo de dados gerados. Considerando a extensão do projeto abrangendo a cidade de São Paulo, a quantidade exponencial de dados provenientes dos sensores requer uma abordagem não relacional, que possibilite a escalabilidade horizontal necessária para lidar com essa imensa carga de informações. Nesse contexto, soluções como MongoDB ou Apache Cassandra despontam como alternativas ideais, dada sua capacidade de escalar horizontalmente e lidar eficientemente com volumes de dados em constante crescimento.


## Referências

[^1]: [O que é um banco de dados relacional (RDBMS)?](https://www.oracle.com/br/database/what-is-a-relational-database/)