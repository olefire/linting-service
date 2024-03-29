# Linting service
<hr>

## About the project

Linting service is a service that performs static code analysis on source code written in various programming languages. This helps identify potential problems, coding errors, and ensure compliance with coding standards.

<hr>

## Features
The project consists of two services. The main service is responsible for registration and authorization of users, as well as for saving and managing their data. The project uses MongoDB for data storage.


The second service provider static code analysis functionality. It accepts code analysis requests from the main service via the gRPC protocol. Users can upload the source code to the main service, which transmits the request for analysis to the second service. The second service performs the analysis and sends the results back to the main service.

The client part of the project was developed using the ReactJS library. It is a single-page application that interacts with the main service via REST API.

<hr>

## Usage

````
$ git clone https://github.com/olefire/linting-service.git
$ cd linting-service
````
#### Run with Docker:

<hr>

### Status

The project is currently in the alpha stage.

<hr>

### UI/UX

<p align="center">
    <img src="img/img.png" width="400">
</p>


<hr>

### Folder structure

```tree
linting-service/
├── backend/
│   ├── cmd/
│   │   └── app/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── controller/
│   │   │   └── http/
│   │   │       ├── controller.go
│   │   │       └── router.go
│   │   ├── models/
│   │   │   ├── paste.go
│   │   │   └── user.go
│   │   ├── repository/
│   │   │   ├── paste/
│   │   │   │   └── repository.go
│   │   │   └── user/
│   │   │       └── repository.go
│   │   └── services/
│   │       ├── paste/
│   │       │   └── service.go
│   │       └── user/
│   │           └── service.go
│   ├── pkg/
│   │   ├── middleware
│   │   └── logs
│   ├── .env
│   ├── docker-compose.yaml
│   ├── Dockerfile
│   ├── go.mod
│   └── Makefile
├── lint-service/
│   ├── api/
│   │   └── gen/
│   │       └── linting-service-api/
│   │           ├── lint.pb.go
│   │           ├── lint.pb.gw.go
│   │           └── lint_grpc.pb.go
│   ├── cmd/
│   │   └── main.go
│   ├── code/
│   │   ├── __init__.py
│   │   └── file.py
│   └── internal/
│       ├── controller/
│       │   ├── controller.go
│       │   └── router.go
│       ├── gapi/
│       │   └── services/
│       │       ├── linter/
│       │       │   ├── service.go
│       │       │   └── service__test.go
│       │       └── interfaces.go
│       ├── pkg/
│       │   └── executor.go
│       └── go.mod
└── frontend/
    └── src/
        ├── services/
        │   ├── lint.tsx
        │   ├── paste.tsx
        │   └── utils.tsx
        ├── styles/
        │   └── PastePage.css
        └── views/
            └── paste.tsx
```

