# konoha-chat

Chat de conversa entre diversas linguagens de programação, com propósito educacional. Não se preocupa com nenhum tipo de autenticação ou segurança, apenas com o envio de mensagens.

## Tabela de conteúdos

1. **[Docker](#docker)**
   1. _[Estrutura do arquivo docker-compose.yml](#estrutura-do-arquivo-docker-composeyml)_
   1. _[Comandos úteis](#comandos-docker-úteis)_
1. **[Banco de dados](#banco-de-dados)**
   1. _[Arquivo .env banco](#arquivo-env-banco)_
   1. _[Estrutura do banco de dados](#estrutura-do-banco-de-dados)_
   1. _[Padrões de nomenclatura](#padrão-de-nomenclatura)_
   1. _[Tabela tb_users](#tabela-tb_users)_
   1. _[Tabela tb_messages](#tabela-tb_message)_
1. **[API](#api)**
   1. _[Arquivo .env api](#arquivo-env-api)_
   1. _[Rodando a api](#rodando-a-api)_
   1. _[Endpoints disponíveis](#endpoints-disponíveis)_

### Docker

É possível os códigos deste repositório sem ter instalado nada na máquina exceto o [Docker](https://docs.docker.com/get-docker/) e [Docker-Compose](https://docs.docker.com/compose/install/), após instalado, é necessário somente rodar o seguinte comando estando na pasta que contém o arquivo [`docker-compose.yml`](docker-compose.yml):

```sh
docker-compose up -d # a flag `-d` representa que os containers dockers serão executados sem 'prender' o terminal, deixando-o livre para outros comandos.
```

#### **Estrutura do arquivo docker-compose.yml**

Essa seção explica a estrutura do arquivo [`docker-compose.yml`](docker-compose.yml) que é utilizado para configurar o servidor e o banco de dados do chat.

```yml
version: '3'
services:
  # declara o container do banco de dados
  database: 
    # utiliza a imagem oficial postgres
    image: postgres
    # define o local do arquivo .env que injetará as variáveis de ambiente
    env_file:
      - database/.env
    # define o fuso horário que será utilizado para o container do banco de dados
    environment:
      - TZ=America/Sao_Paulo
    # expõe a porta para o host
    ports:
      - "5432:5432"
    # define os volumes que serão utilizados para o container do banco de dados, volumes não são removidos quando rodado docker-compose down, se o container for recriado, os volumes serão reutilizados, caso o volume for uma pasta no host, não será removido pelo docker volume prune
    volumes:
      # volume criado posteriormente no arquivo, é removido pelo docker volume prune
      - database:/var/lib/postgresql/data 
      # volume é uma pasta na raiz do host, não é removido pelo docker volume prune, insere um arquivo init.sql que cria a estrutura do banco de dados
      - ./database/init.sql:/docker-entrypoint-initdb.d/init.sql
  # declara o container da api
  api:
    # utiliza a imagem oficial python
    image: python
    # define os volumes que serão utilizados para o container da api, volumes não são removidos quando rodado docker-compose down, se o container for recriado, os volumes serão reutilizados, caso o volume for uma pasta no host, não será removido pelo docker volume prune
    volumes:
      # volume é uma pasta na raiz do host, não é removido pelo docker volume prune
      - ./api:/home/api
    # sobreescreve o arquivo que roda ao iniciar o container, com o existente no volume supracitado, esse script é utilizado pra iniciar a api
    entrypoint: /home/api/entrypoint.sh
    # define o diretório padrão de execução dos comandos
    working_dir: /home/api
    # define o local do arquivo .env que injetará as variáveis de ambiente
    env_file:
      - api/.env
    # define o fuso horário que será utilizado para o container da api
    environment:
      - TZ=America/Sao_Paulo
    # expõe a porta para o host
    ports:
      - "8080:8080"
# declara o volume virtual
volumes:
  database:
    driver: local
```

#### **Comandos docker úteis**

- ```sh
  docker-compose down # desliga e remove dados dos containers
  ```

- ```sh
  docker ps # lista todos os containers sendo executados na máquina
  ```

- ```sh
  docker volume prune # remove todos os volumes que não estão em uso, ou seja, após rodar o docker-compose down, o volume fica sem uso
  ```

### Banco de dados

Explicações sobre o banco de dados do konoha-chat.

#### **Arquivo .env banco**

Esse arquivo é utilizado para configurar o container docker do banco de dados. Caso você deseje rodar localmente, esse arquivo pode ser ignorado.

Para utilizá-lo é necessário renomear ou fazer uma cópia renomeando do arquivo [.env_example](database/.env_example) para [.env](database/.env), alterando adequadamente.

#### **Estrutura do banco de dados**

Composta por duas simples tabelas, segue alguns padrões de nomenclatura:

#### **Padrão de nomenclatura**

- Prefixo `tb_`: para tabelas;
- Prefixo `tx_`: quando o campo for um `TEXT`;
- Prefixo `nr_`: quando o campo for numérico (`NUMERIC(X,Y)`, `INTEGER`);
- Prefixo `dt_`: quando o campo for uma data (`DATE`, `DATETIME`, `TIMESTAMP`);
- Prefixo `fk_`: para chaves estrangeiras;
- Sufixo `_id`: para campos únicos utilizados em chaves estrangeiras, exceto a chave primária da tabela que é nomeada somente como `id`;

#### **Tabela `tb_users`**

```sql
CREATE TABLE tb_user (
    id SERIAL PRIMARY KEY, -- chave primária
    tx_name TEXT NOT NULL -- nome do usuário
);
```

#### **Tabela `tb_message`**

```sql
CREATE TABLE tb_message (
    id SERIAL PRIMARY KEY, -- chave primária
    sender_id INTEGER NOT NULL, -- id do usuário que enviou a mensagem
    recipient_id INTEGER NOT NULL, -- id do usuário que recebeu a mensagem
    tx_message TEXT NOT NULL, -- mensagem
    id_message_type TEXT NOT NULL DEFAULT 'plain/text', -- tipo da mensagem (plain/text, json, html, markdown, xml)
    dt_created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- data e hora de criação por padrão é a data e hora atual do sistema na inserção
    dt_readed_at TIMESTAMP NULL DEFAULT NULL, -- data e hora de leitura, nula se a mensagem ainda não foi lida

    CONSTRAINT fk_sender_id FOREIGN KEY (sender_id) REFERENCES tb_user (id), -- chave estrangeira para a tabela tb_user
    CONSTRAINT fk_recipient_id FOREIGN KEY (recipient_id) REFERENCES tb_user (id) -- chave estrangeira para a tabela tb_user
);
```

### Api

Api em python com FastApi para leitura e escrita simples de dados no banco. Não utiliza nenhum tipo de autenticação.

#### **Arquivo .env api**

Serve para configurar o container docker da api. Caso você deseje rodar localmente, pode ser utilizado o pacote [Python-DotEnv](https://pypi.org/project/python-dotenv/) para a injeção de variáveis de ambiente, ou definindo estas variáveis no seu sistema operacional.

Para utilizá-lo é necessário renomear ou fazer uma cópia renomeando do arquivo [.env_example](api/.env_example) para [.env](api/.env), alterando adequadamente.

#### **Rodando a api**

Caso você esteja utilizando docker a api tentará iniciar automaticamente, por padrão a api está rodando na porta 8080 e localhost, caso deseje alterar essas configurações, basta alterar o arquivo [entrypoint.sh](api/entrypoint.sh) e reinicializar o container. Para rodar localmente é necessário criar um ambiente virtual e instalar os pacotes contidos em [requirements.txt](api/requirements.txt) _([Como fazer?](https://realpython.com/python-virtual-environments-a-primer/#create-it))_, executando o comando existente no mesmo arquivo.

#### **Endpoints disponíveis**

Para consultar os endpoints disponíveis, após iniciar a api, acesse a url `http://host:port/docs`, caso rodado via docker com configurações padrão esse caminho será `http://localhost:8080/docs`.
