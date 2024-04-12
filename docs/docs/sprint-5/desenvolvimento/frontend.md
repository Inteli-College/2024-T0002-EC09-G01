---
label: "Frontend"
---

# Documentação do Front-end em React

A fim de o usuário administrador pudesse criar nosso sensores e que os cidadãos pudessem informar de situações para que outros fiquem alerta, foi desenvolvida um aplicação web. Esta seção é dedicada a parte front-end. Foi desenvolvida em react, estando organizada na seguinte estrutura de pastas:
```
.
├── alertPage.tsx
├── App.css
├── App.test.tsx
├── App.tsx
├── components
│   ├── alertForm.tsx
│   └── form
│       ├── button.tsx
│       ├── form.tsx
│       └── input.tsx
├── index.css
├── index.tsx
├── logo.svg
├── react-app-env.d.ts
├── reportWebVitals.ts
└── setupTests.ts
```

## Funcionalidades Principais

- **Formulário de Sensores**:
  - Permite ao usuário inserir dados de sensores, incluindo sensor, tipo, longitude e latitude.

<img title="Table Schema" alt="Imagem representando as tabelas contidas no banco de dados" src={require('/img/sensorPage.jpeg').default} />

- **Formulário de Alertas**:
  - Semelhante ao formulário de sensores, mas usado para inserir dados de alertas, incluindo alerta, tipo, longitude e latitude.

<img title="Table Schema" alt="Imagem representando as tabelas contidas no banco de dados" src={require('/img/alertPage.png').default} />

- **Integração com o Servidor**:
  - Os formulários usam requisições HTTP POST para enviar os dados para o servidor em Go.
  - As respostas do servidor são tratadas de acordo com o sucesso ou falha do envio dos dados.

## Lançando a aplicação
Para lançar a aplicação é necessário ter instalado node.js em seu computador, caso não tenha é possível baixá-lo neste [link](https://nodejs.org/en). Assim que tiver baixado, basta executar o comando:

```
npm run build
```