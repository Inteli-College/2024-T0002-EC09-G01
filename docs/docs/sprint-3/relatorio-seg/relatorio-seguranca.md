---
label: "Relatório de Segurança - Sprint 3"
---

# Relatório de Segurança do Sistema - Sprint 3
O relatório analisou a segurança do sistema desenvolvido para cidades inteligentes, com ênfase na integridade e autenticidade dos dados transmitidos pelo broker online em nuvem HIVEMQ. As principais preocupações identificadas incluem:

- **Autenticação Simples**: Recomenda-se a implementação de práticas de autenticação mais robustas, como senhas fortes e autenticação de dois fatores (2FA), para garantir uma camada adicional de segurança.
- **Vulnerabilidades Relacionadas à Disponibilidade**: É necessário implementar medidas preventivas para proteger o sistema contra ataques de negação de serviço (DDoS) e outras ameaças que possam comprometer a disponibilidade do sistema.
- **Intercepção de Dados**: Embora o protocolo TLS seja utilizado para garantir a confidencialidade dos dados durante a transmissão, é crucial considerar outras possíveis formas de invasão, como ataques Man-in-the-Middle (MitM) e vulnerabilidades na implementação do TLS.
- **Arquitetura da Solução**: A atual arquitetura implementada utiliza um simulador de sensores enviando dados para o broker HIVEMQ, que posteriormente são enviados para um banco de dados Postgres integrado ao Metabase. A principal vulnerabilidade está na passagem dos dados para o banco de dados, tornando-o suscetível a ataques.

Além disso, é ressaltado que na Sprint 3 não foram desenvolvidos novos testes e melhorias na segurança devido ao foco no desenvolvimento do sistema. Entretanto, na próxima sprint, a prioridade será dada à implementação de novos testes e aprimoramento nas medidas de segurança.!