<p align="center" >
    <img src="[assets/logo.png](https://storage.googleapis.com/gopherizeme.appspot.com/gophers/1e5bdb9a40326f275582821d1e01ac365d24498f.png)" alt="logo" width="250"/>
<h3 align="center">THN REST API</h3>
<p align="center">REST API with simple audit features</p>
<p align="center">Build with ❤ in Golang</p>
</p>

<p align="center" >
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/franciscof12/rest-api-thn">
    <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/franciscof12/rest-api-thn">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/franciscof12/rest-api-thn">
</p>

# THN REST API with simple audit features
[![Go Report Card](https://goreportcard.com/badge/github.com/franciscof12/rest-api-thn)](https://goreportcard.com/report/github.com/franciscof12/rest-api-thn)

## Project Structure
```
/rest-api_thn
|-- /cmd
|   `-- /api
|       `-- main.go
|-- /internal
|   |-- /handlers
|   |   |-- feature.go
|   |   |-- feature_test.go
|   |   |-- metrics.go
|   |   `-- metrics_test.go
|   `-- /models
|       `-- request_info.go
|-- /pkg
|   `-- /common
|       `-- json.go
|       `-- utils.go
|-- /test
|       `-- integration_test.go
|-- go.mod
|-- go.sum
|-- Makefile
|-- README.md
```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/franciscof12/simple-thn-api
cd thn-rest-api
```

2. Install dependencies:

```bash
make deps
```

## Usage

### Running the Application

To compile and run the application:

```bash
make run
```

### Running Tests

To run all tests:

```bash
make test
```

To run only the unit tests:

```bash
make test-unit
```

To run only the integration tests:

```bash
make test-integration
```

### Cleaning Up

To remove generated files:

```bash
make clean
```

## Endpoints

### `/feature`

- **Method**: GET
- **Description**: Stores information about the request in memory.
- **Response**: "Hello THN!"

### `/metrics`

- **Method**: GET
- **Description**: Provides metrics on how many requests a specific IP has made.
- **Query Parameters**:
  - `ip`: The IP for which you want to know the number of requests made.
- **Response**: `{
  "count": Number of Requests,
  "ip": IP to Query
  }`
