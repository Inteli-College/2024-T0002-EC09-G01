---
label: Gráficos
---

# Planejamento da interface gráfica da solução 

## Prototipação da interface
Para a prototipação de como os dados serão apresentados, foi pensado em  duas telas principais para disponibilização da informação. A primeira tela irá apresentar um gráfico de pizza que mostra a porcentagem dos sensores que estão com uma medição negativa de qualidade de ar ou intensidade de luminosidades, junto a esse gráfico dois gráficos de barras serão apresentados mostrando a média da medida de sensores por região, um em gases e o outro em radiation. 

Acima desses gráficos estariam algumas regiões agrupadas por uma condicional, junto a informações da navbar:

## Versão antiga

<img src={require('/img/VersaoAntiga.png').default} width='100%'/>
<sub>Versão antiga</sub>

## Versão final da prototipação

<img src={require('/img/prototipoVersaoAtual.png').default} width='100%'/>
<sub>Versão Sprint 3</sub>

## Versão atualizada

<img src={require('/img/versaoAtual.png').default} width='100%'/>
<sub>Atualização </sub>

## Explicação da escolhas dos gráficos

### Gráfico de barras 

Para a visualização de valores de medição dos sesnores, decidimos agrupar em médias por região, a fim de mostrar um valor mais abrangente das informações coletadas, o gráfico de barras é interessante não só pela sua visualização comparativa de fácil entendimentos sobre a diferença de grandeza entre os valores medidos ele facilita a exibição de varias informações a serem comparadas.

### Gráfico de pizza

O gráfico de pizza é utilizado para comparar proporcionalmente a quantidade de valores com medições abaixo dos valores ideais esperados e medições que estão dentro do valor esperado, facilitando assim o entendimento rápido da situação geral de todas as áreas abrangidas.

### Mapa de calor

A utilização do mapa de calor é ideal para a visualização de areas que possuem uma caracteristica em comum que pode ser indicada pela cor apresentada em relação a uma legenda. Dessa forma, ele se torna uma opção ideal para apresentar uma visão geral com informações geograficas e regionais.

