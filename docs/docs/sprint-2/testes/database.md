# Teste de Confiabilidade da API com o Banco de Dados

Para garantir a confiabilidade do sistema, foi desenvolvido um arquivo de teste no projeto, denominado `api_db_test.go`. Este arquivo contém funções específicas para testar a integridade e a funcionalidade dos endpoints responsáveis por armazenar dados no banco de dados.

**1. Função `testPostsensor`:**
   - Esta função testa o endpoint responsável por receber dados de sensores.
   - Cria um payload com informações simuladas de um sensor.
   - Realiza uma solicitação POST ao endpoint correspondente.
   - Verifica se a solicitação é bem-sucedida e loga o corpo da resposta para análise.

**2. Função `testPostgas`:**
   - Testa o endpoint encarregado de receber dados sobre gases.
   - Gera um payload com informações fictícias de gases.
   - Envia uma solicitação POST ao endpoint específico.
   - Avalia se a operação é concluída com sucesso e registra o corpo da resposta para avaliação.

**3. Função `testPostradiation`:**
   - Testa o endpoint destinado a dados de radiação.
   - Constrói um payload com informações simuladas sobre radiação.
   - Envia uma solicitação POST ao endpoint correspondente.
   - Verifica se a solicitação é realizada com sucesso e registra o corpo da resposta para inspeção.

**4. Função `testGet`:**
   - Testa a funcionalidade dos endpoints de leitura (GET) para sensores, gases e radiações.
   - Itera sobre os endpoints de leitura e verifica se é possível obter dados do servidor.
   - Registra o corpo da resposta para análise.

**5. Função `TestApi`:**
   - Agrega todas as funções de teste em um único teste global.
   - Chama cada uma das funções de teste anteriormente mencionadas.

Ao executar o teste global `TestApi`, todas as funções individuais são invocadas, garantindo que tanto os endpoints de envio quanto os de leitura funcionem conforme o esperado. Este conjunto de testes proporciona uma verificação abrangente da integridade e confiabilidade do sistema, assegurando que os dados possam ser armazenados e recuperados de forma eficaz.

# Executando o teste
Antes de tudo, garanta que a api responsável pelo controle do banco de dados esteja ativada executando o seguinte comoando no diretório `/database-api`:
```
go run .
```

Agora, rode o seguinte comando nos diretório `/tests`:
```
go test -v api_db_test.go
```



