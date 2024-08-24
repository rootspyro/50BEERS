![banner](https://github.com/user-attachments/assets/e74a2890-b9e9-471b-9a08-3bb0b3749dc0)

# 50 BEERS: The Challenge

In Onctober 2023 I traveled to Spain for 3 months to visit my sister, during this trip a good friend of mine challenged me to taste 50 different beers.

Over the course of those three months I was tasting and documenting more than 50 beers as well as different alcoholic beverages in a diary. Since then I have continued with this hobby and have decided to digitize the content and essence of this diary in the format of a blog.

Likewise in this repository I document how I build this blog.


## Tech Stack

__FRONTEND (BLOG)__

![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![Astro](https://img.shields.io/badge/astro-%232C2052.svg?style=for-the-badge&logo=astro&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
![Vercel](https://img.shields.io/badge/vercel-%23000000.svg?style=for-the-badge&logo=vercel&logoColor=white)

__FRONTEND (CPANEL)__

![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![Next JS](https://img.shields.io/badge/Next-black?style=for-the-badge&logo=next.js&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
![Vercel](https://img.shields.io/badge/vercel-%23000000.svg?style=for-the-badge&logo=vercel&logoColor=white)

__BACKEND__

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

## Setup Project

This project is builded over Ubuntu, so some instrucctions will use linux commands.

### Requirements

To run this project is required to install:

- Node v22.3.0
- Golang v1.22
- GNU Make v4.3
- MongoDB Community Server

### 1. Clone the repository

```shell
$ git clone https://github.com/rootspyro/50BEERS.git
$ cd 50BEERS
```

### 2. Build the backend

2.1 - Enter into the [backend directory](./backend) and install the required golang modules.
```shell
$ cd backend
$ go mod tidy
```

2.2 - Copy the content of the [.env.copy](./backend/.env.copy) file into a new `.env` file.
```shell
$ cp .env.copy 
```

2.3 - Insert the required data on the new .env file.
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

2.4 - Create the mongodb database and also the required collections:

 - country
 - drink
 - location
 - tag

2.5 - Feed the database with the json data in the [database](./backend/db/data) directory.

2.6 - Build and Start the API Server
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

2.7 - Check the server status
```shell
$ curl http://[host]:[port]/api/v1/health

# RESPONSE:
# HTTP 200 - OK
# JSON: {"status":"success","statusCode":200,"data":"Server is up!"}
```
