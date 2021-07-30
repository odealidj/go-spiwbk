# codeid-boiler
This project is code.id boilerplate to create API in Golang Echo Framework
- PORT : 3030
- PATH : /

## Installation

``` bash
# clone the repo
$ git clone 

# go into app's directory
$ cd my-project
```

## Build & Run

``` bash
# build 
$ go build

# run in development 
$ ENV=DEV go run main.go
$ ENV=DEV ./filego

# run in staging 
$ ENV=STAGING go run main.go
$ ENV=STAGING ./filego

# run in production 
$ ENV=PROD go run main.go
$ ENV=PROD ./filego

# run in docker
$ docker-compose up 
```

## Swagger Documentation

Install go swagger
``` bash
# get swagger package 
$ go get github.com/swaggo/swag

# move to swagger dir
$ cd $GOPATH/src/github.com/swaggo/swag

# install swagger cmd 
$ go install cmd/swag
```

Generate API documentation
``` bash
# generate swagger doc
$ swag init --propertyStrategy snakecase
```
to see the results, run app and access {{base_url}}/swagger/index.html

## Feature 
This project have default feature
- Http
- Middleware 
- Database
- Validation
- Auth 
- CRUD  
- Transaction
- Pagination
- Response
- Env
- Redis
- Elasticsearch
- Swagger
- Log

# Author
CodeID Backend Team