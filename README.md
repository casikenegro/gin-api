## Can I use docker? 😎

We got your back, the project already uses docker so you can build docker image 
1.  Clone this repo
2.  Copy the  `.env.example`  file over to your own  `.env`  file and update the variables
3.  Run  `docker-compose up -d`  to setup local environment with Docker


## Description

  

[GIN](https://gin-gonic.com/)  Framework.
[GORM](https://github.com/go-gorm/gorm) ORM starter Package.
  

## Running the app


```bash

# development

$ go run main.go


# production mode

$ go build main.go -o gin-api && ./gin-api

```

  
# Swagger Documentation

## Install Swagger Library

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

## Generate Swagger Documentation

```bash
swag init
```

## License

GO use 3-clause BSD + patent grant.