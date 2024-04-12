---
label: "Diagrama de blocos - v1"
---

# Diagrama de blocos


O diagrama de blocos é uma representação visual de como planeja-se que seja o fluxo de comunicação entre as partes do projeto. Abaixo é mostrado a arquitetura esperada do sistema:

<img src={require('/img/diagrama-de-blocos.png').default} width='100%'/>
<sub>Diagrama de blocos da solução - Autoria própria.</sub>

## Descrição do diagrama solução

### Sensores simulados
Para a exucução do projeto serão usados dados simulados de três sensores diferentes. Dois desses sensores captam dados de qualidade do ar, sendo eles o SPS30, sensor que mede partículas inaláveis (CO2, CO, NO2, MP10, MP2,5) e o MICS-6814, que mede a concentração de CO e NO2 no ar. Além deles será também utilizado um sensor que mede radiação e luminosidade, o RXW-LIB-900, que coleta dados de evapotranspiração e radiação solar. Os dados serão enviados através do protocolo MQTT, por meio de um publisher de informações.

### Broker
O broker é parte essencial para a transmissão de dados via MQTT, protocolo amplamente utilizado para envio de informações coletadas por sensores. Essa parte da solução serve como uma ponte que coordena diversos dados publicas em tópicos. Quando uma mensagem é publicada o broker é quem fará a distribuição dos dados aos inscritos no canal.

### Bridge
A bridge é a parte responsável por fazer a transmissão de dados recebidos pelos sensores para o lado web da solução. Ela é composta por duas partes, sendo a primeira delas um bloco inscrito nos tópicos que os sensores publicam as informações e a segunda um serviço de mensageria que facilita a comunicação assíncrona entre as partes do sistema.

### Banco de dados
O banco de dados será onde as infomações de toda a aplicação serão armazenadas, desde dados capturados pelos sensores à possíveis dados de login de usuários.

### Backend
O banckend da solução é a parte que tem responsabilidade de se comunicar diretamente com o banco de dados e servir informações solicitadas pelos usuários.

### Frontend
O frontend é onde todas infomações são mostardas ao usuário através de uma dashboard.




