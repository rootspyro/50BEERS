# 50 BEERS: Backend

For this project I builded a Restful API to consume the database from the [Blog](../blog) and manage the content on the [Admin Dashboard]().

## Tech Stack

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

## Project Setup

This project is builded over Ubuntu, so some instrucctions will use linux commands.

### Requirements

To run this project is required to install:

- Golang v1.22
- GNU Make v4.3
- MongoDB Community Server

### Steps

1 - Enter into the [backend directory](./backend) and install the required golang modules.
```shell
$ go mod tidy
```

2 - Copy the content of the [.env.copy](./backend/.env.copy) file into a new `.env` file.
```shell
$ cp .env.copy 
```

3 - Insert the required data on the new .env file. MongoDB credentials are required for step 2.4 and 2.5.
```shell
# AUTHOR
AUTHOR_NAME=rootpsyro

# SERVER
PORT=3000
HOST=localhost
SECRET=Your-secret-key-here

# MONGODB
DB_HOST=localhost
DB_PORT=27017
DB_USER=mongodb
DB_PASSWORD=p4ssw0rd
DB_NAME=50BEERS
```

__Important__: You can generate a new Secret key running the command `openssl rand -base64 32`

4 - Run the migrations to create the required collection on the database. 

```shell
$ make migrate

# go run cmd/main.go -migrate
# 2024-08-24 01:23:48.373082598 -0400 -04 | INFO | Running Migrations...
# 2024-08-24 01:23:48.373150584 -0400 -04 | INFO | Creating country collection...
# 2024-08-24 01:23:48.373456774 -0400 -04 | INFO | The country collection was successfully created!
# 2024-08-24 01:23:48.373467031 -0400 -04 | INFO | Creating drink collection...
# 2024-08-24 01:23:48.373683265 -0400 -04 | INFO | The drink collection was successfully created!
# 2024-08-24 01:23:48.373696216 -0400 -04 | INFO | Creating location collection...
# 2024-08-24 01:23:48.373876648 -0400 -04 | INFO | The location collection was successfully created!
# 2024-08-24 01:23:48.373883297 -0400 -04 | INFO | Creating tag collection...
# 2024-08-24 01:23:48.374058152 -0400 -04 | INFO | The tag collection was successfully created!
# 2024-08-24 01:23:48.374064311 -0400 -04 | INFO | All migrations where successfully executed!
```

5 - Seed the database with the default data (Tags, Countries, Locations).
```shell
$ make seed
```

6 - Build and Start the API Server
```shell 
$ make build
# go build -o dist/app cmd/main.go

$ make start
./dist/app
#
#   ...
#   | |
#   | |     
#  /   \    50 BEERS API
# |     |   Listening on localhost:3000...
# |     |   
# |     |
# |_____|   By  rootspyro 

```

7 - Check the server status
```shell
$ curl http://[host]:[port]/api/v1/health

# RESPONSE:
# HTTP 200 - OK
# JSON: {"status":"success","statusCode":200,"data":"Server is up!"}
```
