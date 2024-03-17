---
label: "MongoDB"
---

# Datababe - V2

## Banco de Dados não relacional - MongoDB

Devido à natureza do projeto, no qual vários sensores produzem uma quantidade massiva de dados não estruturados, a utilização de um banco não relacional faz sentido. 

Um banco de dados não relacional armazena dados de uma forma não tabular, tendendo a ser mais flexível que bancos orientados à SQL com estruturas relacionais. No caso do MongoDB, os dados são armazenados em documentos. Essa capacidade de digerir e organizar vários tipos de informações lado a lado torna os bancos de dados não relacionais muito mais flexíveis do que os bancos de dados relacionais.[^1]

<img src={require('/img/sql-vs-nosql.jpg').default} width='100%'/>
<sub>Estrutura SQL vs NoSQL[^2]</sub>

Escolhemos o MongoDB por ser uma ferramenta simples e versátil. Além disso, possui uma infraestrutura em cloud (MongoDB Atlas) que facilita e lida com toda a a complexidade do deploy, possibilitando um maior foco no desenvolvimento de features.


## Criação e setup do MongoDB Atlas

Para configurar uma base de dados no mongoDB Atlas, sugerimos o vídeo explicativo a seguir:

<iframe width="560" height="315" style={{
            display: 'block',
            margin: 'auto',
            width: '100%',
            height: '50vh',
        }} src="https://www.youtube.com/embed/bBA9rUdqmgY?si=Iov8zcozfHgwRf61" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

## Organização dos dados

<img src={require('/img/MongoDB-Schema.png').default} width='100%'/>

Nosso banco de dados está estruturado em quatro coleções principais:



- ***alerts***: Armazena os alertas relacionados à desastres em determinadas áreas:
    
    ```bson
        _id: objectId('65f444d56bc7ac0d06e4210f')
        latitude: -23.5872976
        longitude: -46.6825611
        type: "Enchente"
    ```

- ***gases***: Armazena as leituras dos sensores de gases no formato abaixo:
  
    ```bson
        _id: objectId('65f1a7ec4d7ef06d11aa2eb0')
        unit: "ppm"
        time: "2024-03-13T10:14:56.471628161-03:00"
        methane: 5386.66
        carbon_monoxide:861.17
        nitrogen_dioxide: 3.42
        ethanol: 137.45
        ammonia: 375.11
        sensor: "MiCS-6814"
        id: 3
        propane: 1261.57
        iso_butane: 9664.24
        hydrogen: 884.05
    ```
    

- ***radiation***: Armazena as leituras dos sensores de gases no formato abaixo:

    ```bson
        _id: objectId('65f1a7eb4d7ef06d11aa2eaf')
        id: 4
        time: "2024-03-13T10:14:56.730273086-03:00"
        radiation: 794.12
        sensor: "RXWLIB900"
        unit: "W/m2"
    ```


- ***sensors***: Armazena os sensores cadastrados no formato abaixo:

    ```bson
        _id: objectId('65f444176bc7ac0d06e4210b')
        latitude: -23.5762
        longitude: -46.72482
        name: "Station 1"
    ```

## Fluxo

1. **Conexão**: A conexão com o banco se dá pela autenticação utilizando as credenciais geradas ao criar um usuário no mongoDB. Note que, na função criada, existem variáveis de ambiente para realizar a autenticação correta do usuário. São elas:

- ***MONGO_USER***: Nome do usuário configurado no banco de dados.
- ***MONGO_PASSWORD***: Senha configurada no banco de dados. Na sessão "Instruções para execução do projeto", siga os passos apresentados para configurar corretamente o arquivo que contém as variáveis de ambiente.


```go
func ConnectToMongo() *mongo.Client{
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load("../../config/.env")

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Recuperar usuário e senha do arquivo .env
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@sensors.zyzjabc.mongodb.net/?retryWrites=true&w=majority&appName=sensors", mongoUser, mongoPassword)).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client
}
```

2. **Inserção**: Como ilustrado na nova arquitetura da solução, os dados são produzidos pelos sensores (para testes, utilizamos o simulador), para então serem manuseados pelo serviço de fila Kafka. Este aciona uma função que trata os dados e então os insere no banco de dados, dependendo da coleção.

```go

    func InsertIntoMongo(client *mongo.Client, data map[string]interface{}) {
        db := client.Database("SmarTopia")
	
	var coll *mongo.Collection
	// fmt.Println(data["payload"])

	payloadData := data["payload"].(map[string]interface{})

	newData := make(map[string]interface{})

	if payloadData["gases-values"] != nil {
        
        gasesValues := payloadData["gases-values"].(map[string]interface{})

		newData["id"] = data["packet-id"]

		newData["time"] = payloadData["current_time"]

		for key, value := range gasesValues["gases-values"].(map[string]interface{}) {
			newData[key] = value
		}

		newData["sensor"] = gasesValues["sensor"]
		newData["unit"] = gasesValues["unit"]

		coll = db.Collection("gases") 

	} else {
        
        radiationValues := payloadData["radiation-values"].(map[string]interface{})

		newData["id"] = data["packet-id"]

		newData["time"] = payloadData["current_time"]

		for key, value := range radiationValues["radiation-values"].(map[string]interface{}) {
            newData[key] = value
		}

		newData["sensor"] = radiationValues["sensor"]
		newData["unit"] = radiationValues["unit"]

		coll = db.Collection("radiation") 
	}

	bsonData, err := bson.Marshal(newData)

	result, err := coll.InsertOne(context.TODO(), bsonData)

	if err != nil {
        log.Fatal(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

```
  3. **Consumo**: Após a inserção dos dados, estes podem ser acessados pela plataforma de BI utilizada (nesse caso, o metabase). A conexão entre estes serviços pode ser feita a partir da ***string*** de conexão gerada automaticamente pelo mongoDB Atlas. Nas configurações de administrador, ao adicionar uma nova base de dados, selecione MongoDB e utilize a opção ***Paste a connection string***.
   
<img src={require('/img/Metabase-mongo-config.png').default} width='100%'/>


   


## Referências

[^1]: [What Is a Non-Relational Database?](https://www.mongodb.com/databases/non-relational)
[^2]: [Relational and non relational databases](https://www.pragimtech.com/blog/mongodb-tutorial/relational-and-non-relational-databases/)

