# Konoha Chat API
* API para possibilitar a troca de mensagens entre as linguagens de programação
* Utilizado a linguagem GO
* Guia básico sobre a linguagem GO (https://learnxinyminutes.com/docs/go/)

## Executar
```go
go get
go run main.go
```

## Compilar
```go
go build
```

## Consumir a API
#### Verificar se está online
- http://localhost:5000/konoha-chat/api/ping

#### Criar usuário
- http://localhost:5000/konoha/chat/createUser?name=NOME_USUARIO

#### Buscar dados do usuário
- http://localhost:5000/konoha/chat/getUser?name=NOME_USUARIO

#### Enviar mensagem
- http://localhost:5000/konoha/chat/sendMessage?from=ID_USUARIO&to=ID_USUARIO&message=MENSAGEM

#### Buscar mensagens
- http://localhost:5000/konoha/chat/getMessages?from=ID_USUARIO&to=ID_USUARIO