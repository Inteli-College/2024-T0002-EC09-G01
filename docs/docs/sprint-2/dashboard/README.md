---
id: Dashboard
title: Dashboard
sidebar_position: 5
---


# Dashboard

O dashboard foi fei no metabase em um container no docker. O metabase está conectado diretamente no banco de dados para pegar todas as informações relevantes diretamente e ser atulizado em tempo real. Atualmente o dashboard tem um mapa de São Paulo com as localizações dos sensores mocados, um gráfico de barras para as médias medidas dos gases, outro gráfico de barras com a média das medidas da radiação e o numero explicito da quantidade de sensores registrado no sistema. 

# Execução do dashboard

Para executar o container com o dashboard basta rodar o seguinte comando

```
yarn start <container_id>
```
