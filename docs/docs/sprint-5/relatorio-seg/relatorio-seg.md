# Relatório de Segurança do Sistema - Projeto
## Introdução

- O presente relatório tem como objetivo fornecer uma análise detalhada da segurança do sistema desenvolvido para cidades inteligentes, com foco na integridade e autenticidade dos dados transmitidos através do broker online em nuvem HIVEMQ.
Avaliação de Segurança

## 1. Autenticação Simples

O broker HIVEMQ utiliza um sistema de autenticação simples, onde senhas podem ser consideradas uma potencial vulnerabilidade. A autenticação é uma parte crucial da segurança do sistema e deve ser abordada com cuidado. Recomenda-se a implementação de práticas de autenticação mais robustas, como o uso de senhas fortes e autenticação de dois fatores (2FA), para garantir uma camada adicional de segurança.

## 2. Vulnerabilidades Relacionadas à Disponibilidade

Durante a avaliação, identificou-se a ausência de tratativas para ataques relacionados à disponibilidade do sistema. Recomenda-se a implementação de medidas preventivas, como controle de acesso adequado, limitação de tentativas de conexão e monitoramento proativo para detectar e responder a possíveis ataques de negação de serviço (DDoS) ou outras ameaças que possam impactar a disponibilidade do sistema.

## 3. Intercepção de Dados

Apesar do uso do protocolo TLS para assegurar a confidencialidade dos dados durante a transmissão, é essencial considerar outras possíveis formas de invasão que podem comprometer a integridade ou a disponibilidade dos dados:

- 1.  Ataques Man-in-the-Middle (MitM): Embora o TLS seja eficaz contra MitM, é crucial manter as bibliotecas e configurações TLS atualizadas para evitar vulnerabilidades conhecidas. Além disso, monitorar ativamente a presença de certificados SSL/TLS inválidos ou não confiáveis pode ajudar a detectar tentativas de MitM.

- 2. Vulnerabilidades no Implementação TLS: Certificar-se de que a implementação do TLS no sistema esteja configurada corretamente e não seja suscetível a ataques como o Protocol Downgrade ou o Logjam. Manter-se informado sobre as atualizações de segurança e aplicá-las regularmente é fundamental.

- 3. Segurança da Chave Privada: Garantir que as chaves privadas utilizadas para estabelecer a comunicação segura estejam armazenadas de maneira segura e que apenas as entidades autorizadas tenham acesso. Uma violação da segurança da chave privada poderia comprometer a eficácia do TLS.

## 4. Arquitetura da solução 

A atual arqutitetura implementada utiliza da criação de um simulador de sensores que são eviados para o Broker HIVEMQ que possui um sistema de auteticação supracitado, posteriomente esses dados são enviados para um subscriber que realiza alguns testes de integração para garantir a funcionalidade do sitema, assim ele envia esses dados diretamente para o banco de dados Postgress que está integrado com o metabase que realiza a apresentação dos dados.

A principal vulnerabilidade do siste a está contida na passagem dos dados feitas para o banco de dados que está vulnerável a ataques podendo compromenter o funcionameto da dash de apresentação e até realizar o sequestro de dados que seriam gravado no banco de dados.

Logo, a contramedida pensada é na criação de uma API que faça o controle dos dados que serão enviados para o banco garantindo uma camada de segurança antes de gravar os dados no banco. Posteriormente é visto a necessidade de criar validações de sessões e permissionamento no sistema contudo, essa feature não foi implementada ainda

## Conclusão

Este relatório destaca as principais preocupações de segurança identificadas no sistema do projeto em desenvolvimento. 
A implementação das recomendações propostas fortalecerá significativamente a segurança do sistema, proporcionando uma base sólida para a integridade, autenticidade e disponibilidade dos dados. 