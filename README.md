# paper Test

This repository was created as a test that can solve the simple problem of e-wallet service.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)

## Installation

To get started with this project, follow these steps:

1. Extract from zip
2. This project are using go version 1.22.4
3. Install the required Go packages:
    ```sh
    go get github.com/gorilla/mux
    go get golang.org/x/crypto/bcrypt
    go get github.com/stretchr/testify
    ```

## Usage

To run the application, use the following command:
```sh
go run main.go
```

To build the application, use the following command:
```sh
# if using windows
go build -o paper_test.exe ./
## after build you can execute the .exe file

# if using linux or mac
go build -o paper_test ./
## after build you can execute the file using
./paper_test
```

## Project Structure

```sh
paper-test/
│
├── main.go        # The main entry point of the application
├── collection     # Contains Postman collection for testing purposes
├── constant       # Contains constants used in the repository, such as loan statuses or user types
├── handler        # Contains handler functions for REST API endpoints
├── helper         # Contains helper functions
├── storage        # contains function to access data, since no database is used, these functions are used to access data in memory
├── model          # Contains object structs and their associated methods
└── README.md      # Project documentation
```

## Dependencies

This project uses the following dependencies:
```sh
gorilla/mux: HTTP router for handling routing in Go applications.
```
