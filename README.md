# Project Name

Ready-made code for creating fast apis with the gin framework

## Table of Contents

- [Install](#install-)
- [Run](#run)
  - [System](#system)
  - [Docker](#docker)

## Install 

To run this project locally, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/daniil-sargsyan/go-basic
    cd go-basic
   ```
## Run
    
Create an environment file (e.g., `.env`) with the following configurations:

   ```env
      DB_HOST=localhost
      DB_PORT=5432
      DB_USER=root
      DB_PASS=root
      DB_NAME=postgres
      BASE_PATH=/api/v1
      ADDRESS=localhost:80
      RUN_MODE=dev
   ``` 

### System

1. Build the project:

    ```bash
    go build -o api cmd/main.go
    ```

2. Run the application with the provided environment file:

    ```bash
    ./api path/to/env
    ```

### Docker

1. Create a Docker network:

    ```bash
    docker network create base
    ```

2. Run PostgreSQL container:

    ```bash
    docker run --name basic_postgres -e POSTGRES_PASSWORD=your_db_password -e POSTGRES_USER=your_db_user -e POSTGRES_DB=your_db_name -p 5432:5432 --network=your_network -d postgres
    ```
3. Build Docker image and run container:   
   ```bash
   docker build -f Dockerfile -t go-basic-image .
   docker run --name go-basic-container -p 80:80 --network my_network go-basic-image
   ```




