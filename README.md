# API em Go usando Gin Gonic e GORM

Este projeto é uma API construída utilizando o padrão MVC (Model-View-Controller). Ele utiliza o framework Gin Gonic para o roteamento e manipulação de requisições HTTP, e o Gorm como um ORM (Object-Relational Mapper) para a interação com o banco de dados.

## Tecnologias Utilizadas

- **Gin Gonic**: Um framework web de alto desempenho escrito em Go (Golang). Ele foi desenvolvido para construir aplicações web robustas e eficientes.
- **Gorm**: Um fantástico ORM (Object-Relational Mapper) para Golang, que suporta as operações básicas de CRUD e muito mais.
- **Docker**: Utilizado para a criação e gerenciamento do banco de dados em um ambiente isolado.
- **DBeaver**: Uma ferramenta de banco de dados universal gratuita e de código aberto para desenvolvedores e administradores de banco de dados. Neste projeto, é utilizado para conectar e interagir com o banco de dados.
- **Postman**: Uma plataforma de colaboração para o desenvolvimento de API. Neste projeto, é utilizado para testar as rotas e funcionalidades da API.

## Como Rodar o Projeto

1. Clone este repositório.
```git
git@github.com:mfcbentes/go-gin-gorm.git
```
2. Instale as dependências.
```go
go mod tidy
```
3. Crie um container Docker para o banco de dados Postgres.
```docker
docker run --name db-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

4. Pode utilizar um gerenciador para conectar ao banco (opcional). Ex: [DBeaver](https://dbeaver.io/download/)

5. Inicie o servidor.
```go
go run main.go
```
6. Use o [Postman](https://www.postman.com/downloads/) para testar as rotas da API.



## Contribuição

Contribuições são sempre bem-vindas. Sinta-se à vontade para abrir uma issue ou enviar um pull request.
