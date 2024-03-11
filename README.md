# go-basic

Ready-made code for creating fast apis with the gin framework

## Table of Contents

- [Run](#run)
- [Usage](#usage)

## Run

To run this project locally, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/daniil-sargsyan/go-basic
    cd go-basic
    ```

2. Build the project:

    ```bash
    go build -o api cmd/main.go
    ```

3. Run the application with the provided environment file:

    ```bash
    ./api path/to/env
    ```

## Usage

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
