Aqui, você deve colocar os arquivos de configuração do projeto. Isso inclui:
***.env***: Arquivo para configurar as variáveis de ambiente
***mosquitto.conf***: Configuração do broker local, caso queira substituir o cluster
***Arquivos .pem***: Isto inclui certificados da Autoridade Certificadora (CA), bem como a do cliente e sua chave privada.

Para gerar os certificados e a chave privada, utilize o tutorial do HIVEMQ clicando [aqui!](https://www.hivemq.com/blog/securing-hivemq-broker-deployments-with-intermediate-ca-certificates/)

> Nota: Temos a conciência de que os arquivos não deveriam estar no repositório do projeto. Porém, por motivos de facilidade e demonstração, disponibilizamo-los aqui.