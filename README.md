---
title: ‌Introduction To GraphQL Server With Golang
description: Introduction to GraphQL Server with Golang and Gqlgen.
tags: graphql, go, api, gqlgen
---
## Table Of Contents
- [Table Of Contents](#table-of-contents)
    - [How to Run The Porject](#how-to-run-project)
    - [Getting started ](#getting-started)
        - [Project Setup](#project-setup)
        - [Defining Our Schema](#defining-out-schema)
    - [Queries](#queries)
        - [Simple Query](#simple-query)
    - [Mutations](#mutations)
        - [A Simple Mutation](#a-simple-mutation)
    - [Database](#database)
        - [Setup MySQL](#setup-mysql)
        - [Models and migrations](#models-and-migrations)
    - [Create and Retrieve Links](#create-and-retrieve-links)
        - [CreateLinks](#createlinks)
        - [links Query](#links-query)
    - [Authentication](#authentication)
        - [JWT](#jwt)
        - [Setup](#setup)
            - [Generating and Parsing JWT Tokens](#generating-and-parsing-jwt-tokens)
            - [User Signup and Login Functionality](#user-signup-and-login-functionality)
            - [Authentication Middleware](#authentication-middleware)
    - [Continue Implementing schema](#continue-implementing-schema)
        - [CreateUser](#createuser)
        - [Login](#login)
        - [RefreshToken](#refresh-token)
        - [Completing Our App](#completing-our-app)


### How to Run The Project <a name="how-to-run-project"></a>
First start mysql server with docker:
```bash
docker run --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -d mysql:latest
```
Then create a Table names hackernews for our app:
```sql
docker exec -it mysql bash
mysql -u root -p
CREATE DATABASE hackernews;
```
finally run the server:
```bash
go run server/server.go
```
Now navigate to https://localhost:8080 you can see graphiql playground and query the graphql server.


### Tutorial
to see the latest version of tutorial visit https://www.howtographql.com/graphql-go/0-introduction/
