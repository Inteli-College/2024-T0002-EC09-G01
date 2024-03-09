---
id: Testes
title: Testes - Cliente
sidebar_position: 6
---


# Teste de Funcionamento do Client   

O teste de funcionamento do client visa essencialmente validar as atividades básicas que ele irá executar, como: 

**1. Função `Create a Client`:**
   - Este teste visa a criação de um cliente, utilizando a função CreateClient, que foi criada para facilitar a criação de um cliente mqtt.

**2. Função `Subscribe to topic`:**
   - Essta função testa tanto a inscrição à um tópico teste, quanto o cancelamento dela.

**. Função `Publish message`:**
   - Esta função testa a publicação de um cliente em um tópico teste

Apesar de teste simples, eles garantem que o básico da aplicação esteja sempre funcionando.

# Executando o teste
Para rodar o teste, basta entrar no diretório `src/tests`
```
go test -v client_test.go
```



