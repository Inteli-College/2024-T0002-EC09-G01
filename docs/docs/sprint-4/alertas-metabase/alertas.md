# Criação de alertas no metabase 

- Para a criação de alertas no metabse é possivel criar três tipos diferentes de alertas sendo eles: alerta de ultrapassagem de linha de meta, alerta de barra de progresso e alera de resultados. Optamos por utilizar o alerta de resultado no primeiro teste devido a sua praticidade de uso, entretanto poderemos utilizar também o alerta de barra de progresso ou de linha de meta para complementar algumas informações. 

- A principal diferença entre o alerta de resultado dos outros tipos de alertas é que para ser emitido um alerta basta que uma pergunta (question) retorne qualquer tipo de resutado, assim ele se torna ideal para o alerta de valores outlier ou condições que apresentam algum risco. Já o alerta de linha de meta, ele emite um alerta toda vez que um grafico de linhas com base em uma séria temporal atinja o valor setado como meta, sendo util para mostrar alguns picos em variações de informações ao longo do tempo. O alerta de barra de progresso emite um alerta após o grafico de barra de progresso atingir um valor numérico espefico, esse tipo de alerta é interessante para utilizar na somatória de valores que podem gerar uma situação de risco tendo que ser alertado sobre o progresso dessa situação.

## Como os emails são enviados 

- Dentro da página de admin do metabase é possivel configurar um email atraves do protocolo smtp para o envio de mensagems, assim após criar uma questio é possivel definir o alerta podendo escolher se ele será enviado de hora em hora, diário, semanalmente ou mensal. Apesar de não ser possivel definir um horário especifico, é possivel definir um horário cheio (Ex. 6:00 ou 7:00) para o envio de alertas.

## Fotos do email

<img src={require('/img/foto-email1.png').default} width='100%'/>
<sub> Caixa de entrada</sub>

<img src={require('/img/foto-email2.png').default} width='100%'/>
<sub> Email aberto</sub>

## Integração com slack

- Além de ser possivel enviar alertas por email o metabase possibilita a integração com slack, possibilitando a escolha de um canal para criar um bot que envia mensagems no periodo definido pelo alera enviando fotos da question

## Foto slack

<img src={require('/img/foto-slack.png').default} width='100%'/>
