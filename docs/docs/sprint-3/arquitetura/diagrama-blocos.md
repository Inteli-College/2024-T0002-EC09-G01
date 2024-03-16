---
label: "Diagrama de blocos - v2"
---

# Diagrama de blocos


O diagrama de blocos é uma representação visual de como planeja-se que seja o fluxo de comunicação entre as partes do projeto. Abaixo é mostrado a arquitetura esperada do sistema:

<img src={require('/img/diagrama-de-blocos-v3.png').default} width='100%'/>
<sub>Diagrama de blocos da solução - Autoria própria.</sub>

## Descrição do diagrama solução

### Sensores simulados
Para a exucução do projeto serão usados dados simulados de dois sensores diferentes. Sendo um deles o MICS-6814, que mede a concentração de CO e NO2 no ar. Além deles será também utilizado um sensor que mede radiação e luminosidade, o RXW-LIB-900, que coleta dados de evapotranspiração e radiação solar. Os dados serão enviados através do protocolo MQTT, por meio de um publisher de informações construído com a linguagem de programação Golang e a biblioteca Paho.

### Broker
O broker é parte essencial para a transmissão de dados via MQTT, protocolo amplamente utilizado para envio de informações coletadas por sensores. Essa parte da solução serve como uma ponte que coordena diversos dados publicas em tópicos. Quando uma mensagem é publicada o broker é quem fará a distribuição dos dados aos inscritos no canal. Para sua solução será utilizado o HiveMQ pois essa ferramenta permite implementar um cluster próprio para escalabilidade, alta disponibilidade e segurança em sistemas de mensageria.

### Serviço de mensageria
Como serviço de mensageria foi escolhido o Kafka hospedado na Confluent. Essa escolha foi feita baseando-se na integração nativa entre o HiveMQ e a Confluent. O uso do sistema de mensageria foi usado para que o sistema fosse escalável para o recebimento de um alto fluxo de informações.

### Consumer
O papel do consumer é acessar as mensagens registradas no serviço de mensageria Kafka e registrar no banco de dados. 

### Banco de dados
O banco de dados escolhido para ser usado na solução foi o MongoDB. Sua escolha foi feita para garantir a escabalabilidade da adição de informações desestruturadas e simplicidade de acesso.

### Dashboard
O frontend foi contruido utilizando uma ferramenta de Bussiness Inteligence (BI), o Metabase, para facilitar o acesso a informações do bando de dados e construção de gráficos.



