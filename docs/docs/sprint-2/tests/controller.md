# Teste de Funcionamento do Controller - função responsável por coordenar a aplicação   

A fim de validar as mensagens recebidas, foram desenvolvidos dois testes: 

**1. Função `TestPublishFields`:**
   - Esta função testa se o payload da mensagem recebida se encaixa no padrão desejado.
   - Utiliza regex para comparar campos esperados e campos recebidos.
   - Utiliza o message handler para testar a função.

**2. Função `TestQos`:**
   - Testa se o Qos recebido corresponde o definido pelo cliente.
   - Utiliza o message handler para testar a função.

Ambas funções visam validar que a mensagem recebida se encaixa nos padrões pré estabelecidos. Deste modo garante-se a integridade de dados e evita o enviesamento de 

# Executando o teste
Para rodar o teste, basta entrar no diretório `src/tests`
```
go test -v controller_test.go
```



