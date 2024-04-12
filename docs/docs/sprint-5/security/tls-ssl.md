# Segurança

## Implementação do TLS & SSL

<img title="TLS SSL Handshake" alt="TLS SSL Handshake" src={require('/img/tls-ssl-handshake.png').default} width='100%' />

TLS (Transport Layer Security) é um protocolo de segurança que garante a comunicação segura pela internet. Anteriormente conhecido como SSL (Secure Sockets Layer), TLS protege os dados transmitidos entre clientes e servidores, criptografando-os para evitar interceptações e manipulações por terceiros mal-intencionados. Funciona através de criptografia assimétrica e simétrica, estabelecendo um canal seguro de comunicação entre as partes envolvidas.[^1]

## Autoridades Certificadoras (CAs)

Para garantir a segurança na troca de mensagens entre o cliente e o Broker, optamos por utilizar criptografia das mensagens de forma a garantir que não haja interceptação ou modificação de seu conteúdo. Para isso, decidimos criar Autoridades Certificadoras, uma autoridade Raiz e uma Intermediária, que assinam certificados. Em uma prática de segurança padrão, a Autoridade de Certificação Raiz (CA) não é utilizada diretamente para assinar certificados de servidor ou cliente. Em vez disso, a função principal da CA raiz é estabelecer uma ou mais autoridades de certificação intermediárias. Essas CAs intermediárias são entidades confiáveis, designadas pela CA raiz, para lidar com a tarefa de assinatura de certificados. A adoção dessa abordagem hierárquica atende a um propósito crítico: permite que a chave raiz permaneça off-line e em grande parte inativa, aumentando assim a segurança geral. Se a chave intermediária for comprometida, a CA raiz poderá responder efetivamente revogando o certificado intermediário comprometido.

Para utilizá-lo no contexto do projeto, criamos uma configuração para comunicação TLS/SSL em um cliente. Primeiro, ele estabelece um pool de certificados vazio. Em seguida, lê e adiciona um certificado intermediário ao pool, se disponível. Em seguida, carrega um par de chave/certificado do sistema de arquivos, que será usado pelo cliente para autenticação. Depois, cria e retorna um objeto tls.Config com as configurações especificadas, incluindo o pool de certificados raiz, configuração de autenticação do cliente para não exigir um certificado do cliente, e a opção de ignorar a verificação de certificado do servidor, que é uma prática insegura em ambientes de produção.

```go
func NewTLSConfig() *tls.Config {
	certpool := x509.NewCertPool()

	pemCerts, err := os.ReadFile("../../config/intermediate.cert.pem")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	cert, err := tls.LoadX509KeyPair("../../config/mqtt-client.cert.pem", "../../config/mqtt-client.key.pem")

	return &tls.Config{
		RootCAs:    certpool,
		ClientAuth: tls.NoClientCert,
		ClientCAs:  nil,
		InsecureSkipVerify: true,
		Certificates: []tls.Certificate{cert},
	}
}
```

<img title="Autoridades Certificadoars" alt="Autoridades certificadoras" src={require('/img/CAs.png').default} width='100%' />

## Geração de certificados

A geração desses certificados envolve a criação de autoridades certificadoras e o uso do terminal de comando. Um tutorial completo para criação das chaves utilizadas no projeto pode ser conferido [aqui!](https://www.hivemq.com/blog/securing-hivemq-broker-deployments-with-intermediate-ca-certificates/)[^2]

## Referências

[^1]: [Como funcionam os certificados TLS/SSL | DigiCert](https://www.digicert.com/pt/how-tls-ssl-certificates-work)
[^2]: [Securing HiveMQ Broker Deployments With Intermediate CA Certificates](https://www.hivemq.com/blog/securing-hivemq-broker-deployments-with-intermediate-ca-certificates/)