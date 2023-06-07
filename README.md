# Singo

Singo: Simple Single Golang Web Service

Develop Web services using Singo: With the simplest architecture, implement a sufficient framework to serve a massive number of users.

https://github.com/Gourouting/singo

## Purpose

This project uses a series of popular components in Golang, allowing you to quickly build a Restful Web API based on this project

## Characteristics

This project has integrated many components necessary for developing APIs:


1. [Gin](https://github.com/gin-gonic/gin): A lightweight Web framework, self-proclaimed as having the fastest routing speed in Golang. 
2. [GORM](https://gorm.io/index.html): An ORM tool. This project needs to be used in conjunction with MySQL. 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Session operation tool provided by the Gin framework.
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis client.
5. [godotenv](https://github.com/joho/godotenv): An environment variable tool for the development environment, making it convenient to use environment variables.
6. [Gin-Cors](https://github.com/gin-contrib/cors): Cross-origin middleware provided by the Gin framework.
7. It has implemented some basic functionality of internationalization i18n.
8. This project uses a session based on cookies to save the login status, which can be modified to token verification if needed.

This project has pre-implemented some commonly used code for reference and reuse:


1. Created a user model.
2. Implemented the /api/v1/user/register user registration interface.
3. Implemented the /api/v1/user/login user login interface.
4. Implemented the /api/v1/user/me user profile interface (requires session after logging in).
5. Implemented the /api/v1/user/logout user logout interface (requires session after logging in).

This project has pre-created a series of folders to divide the following modules:

1. The api folder is the controller of the MVC framework, responsible for coordinating various components to complete tasks.
2. The model folder is responsible for storing database models and database operation related code.
3. The service is responsible for handling more complex business. Modeling business code can effectively improve the quality of business code (such as user registration, recharge, order, etc.).
4. The serializer stores common json models, converting the database models obtained from the model into the json objects required by the api.
5. The cache is responsible for the code related to Redis caching.
6. The auth is the authorization control folder.
7. The util includes some common small tools.
8. The conf folder holds some statically placed configuration files, with the locales inside holding translation-related configuration files.



## Go Mod

This project uses Go Mod to manage dependencies.(https://github.com/golang/go/wiki/Modules).

```shell
go mod init go-crud
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // Automatically install
```

## Run

```shell
go run main.go
```

After the project runs, it starts on port 3000 (which can be modified, see the gin documentation).


