# Arquiteura de Solução em diagrama UML

<img title="Table Schema" alt="Imagem representando as tabelas contidas no banco de dados" src={require('/img/arquiteturaUML.jpeg').default} />

## Publisher
O publisher, feito com a biblioteca paho e na linguagem golang, é responsável por simular o envio de dados de sensores, de modo que seja possível validar o fluxo e o funcionamento da solução. Ele publica em tópicos, respectivos a cada tipo de sensor.

## Kafka 
O Kafka é um gerenciador de eventos complexos, que, no caso do projeto, visa organizar o volume de dados gerados nos sensores para a inserção garantida no banco de dados. O utilizado, neste caso, é hosteado pela Confluent e possui integrações com HiveMQ. 

## NoSQL (MongoDB)
O Banco de dados NoSQL foi escolhido pela flexibilidade e escalabilidade. Neste caso, possui as coleções "alerts","gases","radiation" e "sensors", que fazem referência à: alertas que cidadãos podem colocar para avisar de situações perigosa ou delicadas na cidade, informações sobre qualidade do ar e gases atmosféricos, níveis de radiação e informações sobre os sensores espalhados pela cidade.

## Interface Web
A interface web é composta por um Front-end em React e um Back-end em Golang. Juntos, são resposáveis por viabilizar que o usuário possa acessar o site, inserir alertas para avisar a comunidade, ou que o administrador possa adicionar mais um sensor que ele deseja monitorar.

## Dashboard (Metabase)
A interface gráfica escolhida foi o Metabase, pois permite a fácil conexão com diferentes banco de dados, além de poder integrar com diversas outras plataformas, como Gmail e Slack, utilizados neste caso. O Metabase, para conseguir manter as configurações dos dashboards, armazena os dados em um banco SQL, que no projeto foi escolhido o PostGreSQL.
