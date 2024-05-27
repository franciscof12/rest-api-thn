
# THN REST API with simple audit features

## Project Structure
```
/thn-rest-api
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
