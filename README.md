# Konoha-Chat

![image](https://img.shields.io/badge/N%C3%ADvel-M%C3%A9dio-yellow)

Crie uma aplicação fullstack que permita, por meio de uma base de dados, uma comunicação em estilo de chat.

## Esperado

Coisas que são esperadas nesse desafio, porém não obrigatórias:

- [Banco de dados](#banco-de-dados)
- [Backend](#backend)
- [Frontend](#frontend)

---

### Banco de dados

Explicações sobre o banco de dados do konoha-chat, lembre-se de disponbilizar scripts para a criação com explicações detalhadas.

#### **Estrutura do banco de dados**

Deve ser composta por duas simples tabelas, seguem especificações:

1. **Tabela de usuários**: uma tabela que grave o nome e um identificador do usuário;
1. **Tabela de mensagens**: uma tabela que grave o identificador da mensagem, o identificador de quem envia e de quem recebe, conteúdo da mensagem, tipos suportados de mensagens (Ex: `plain/text, json, html, markdown, xml`), data de envio e data de leitura;

### Tecnologias sugeridas banco

Lista de tecnologias que podem ser mais adequadas para desenvolvimento:

- Banco relacional (Ex: `PostgreSQL, SQLite...`);
- Banco não relacional (Ex: `MariaDB, MongoDB...`);
- Docker;

---

### Backend

API que realize as inserções e leituras necessárias no banco de dados, com validações pertinentes. Lembre-se de disponibilizar scripts de deploy com explicações detalhadas.

#### Endpoints sugeridos

Para uso adequado da API, sugerimos que existam os seguintes endpoints:

- Cadastro de usuário;
- Edição de usuários;
- Listagem de usuários;
- Exclusão de usuários;
- Envio de mensagem;
- Leitura de mensagem;
- Listagem de mensagens;
- Exclusão de mensagens;

#### Tecnologias sugeridas API

Lista de tecnologias que podem ser mais adequadas para desenvolvimento:

- Java;
- C++;
- Python;
- Go;
- Docker;

---

### Frontend

Aplicação gráfica para utilização do chat, com validações pertinentes. Lembre-se de disponibilizar scripts de deploy com explicações detalhadas.

#### Telas sugeridas

Sugerimos que exista pelo menos três telas, um crud de usuários, envio de mensagem, listagem de mensagens pro usuário.

#### Tecnologias sugeridas frontend

Lista de tecnologias que podem ser mais adequadas para desenvolvimento:

- HTML/PHP/JavaScript;
- Flutter;
- React;
- Docker;
