# GETNET

Uma api rest desenvolvida em **golang** para o processo seletivo de desenvolvedor backend da **GETNET**.

## Installation & Run

-  Baixe esse projeto

```bash  

-  git clone git@github.com:vitormcid/getnet-challange.git
-  Crie um arquivo .env na raiz do projeto

```
- Instale o docker em sua máquina:

**Mac**
```bash
	https://docs.docker.com/desktop/mac/install/
```

**Linux**
```bash
	https://docs.docker.com/engine/install/ubuntu/
```

Para executar o projeto, basta entrar na raiz e rodar os comandos:

- docker-compose build

- docker-compose up

  

## Migration

- É importante rodar a migration para criar a tabela de usuários:
 ```bash
migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path db/migrations up
```

```bash
Considera-se os seguintes parâmetros:
- user = postgres
- password = postgres
- bd = postgres
- porta = 5432
```
	

- Você também pode criar manualmente no banco se preferir

  
## API

```

# API Endpoint : http://127.0.0.1:8010

```

  

#### /users

* `GET` : retorna todos os usuários

  

#### /user/id

* `GET` : retorna um único usuário

* `POST` : cria um usuário

* `PUT` : altera um usuário

* `DELETE` : remove um usuário

  

#### body:

  ```bash  

{

	name: string

	password: string

	email: string

}
 ```
  

## Teste Unitários

* Os teste unitários são executados quando o projeto é buildado

* Também é possível executalos através do comando **go test -v**

*  Os testes foram desenvolvidos para cobertura de todos os endpoints

## testes de carga

* os testes de carga foram executados com a ferramenta K6 e podem

ser visualizados na pasta load_tests

## Stack

* Golang
* postgreSql
* Gorilla/Mux
* Docker
* Grafana K6


