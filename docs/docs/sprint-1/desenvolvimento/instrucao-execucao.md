---
label: "Instruções para execução do projeto"
---

# Instruções para execução do projeto

## Estutura de pastas

O repositório está organizado da seguinte forma:

```bash
.
├── docs
├── README.md
├── src
├── .github
└── tests
```

- docs: Diretório que contem todos os arquivos referentes à documentação do projeto.
- src: Contem todos os arquivos relacionados com o código fonte da solução
- tests: Armazena testes automatizados da solução
- .github: Contem o fluxo de integração contínua do repositório

### Pré-requisitos
As seguintes ferramentas são necessárias para rodar o projeto sem problemas

- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Go](https://go.dev/doc/install)

## Lançando a aplicação

Para utilizar os módulos, clone o repositório, se ainda não o tiver feito e entre no diretório _src/mqtt_:

```bash
    git clone https://github.com/Inteli-College/2024-T0002-EC09-G01.git
    cd src/mqtt
```


### Publisher:
1. Entre no diretório _publisher_:
   
```bash
cd publisher 
```

2. Rode o comando:
   
```bash
go run .
```

### Subscriber:
1. Entre no diretório _subscriber_:
```bash
cd subscriber
```
2. Rode o comando:
```bash
go run .
```

### Observações
- Existem planos para dockerizar a solução no futuro. 🐋
- Apenas a execução da solução em Go foi demonstrada, uma vez que pretendemos seguir com o desenvolvimento utilizando esta ferramenta, apesar de existirem códigos em Python.
