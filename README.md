# Hello World API

A simple Go API that accepts a name as input and returns a greeting message.

## Project Overview

This project is a Go-based REST API with a simple endpoint that validates name inputs and returns personalized greetings.

## Features

- REST API endpoint that accepts name parameters
- Input validation for names (only allows if the first name start with A-M)
- JSON response formatting
- Comprehensive test coverage

## Architecture

The project follows a clean architecture pattern:

- **Controllers**: Handle HTTP requests and responses
- **Use Cases**: Implement business logic

## API Endpoints

### GET /hello-world?name={name}

Returns a greeting for the provided name.

**Parameters:**
- `name`: A string containing a valid name (first name start with A-M)

**Responses:**
- `200 OK`: Returns a JSON object with a greeting message
  ```json
  {"message": "Hello Name"}
  ```
- `400 Bad Request`: Returns when the name is invalid
  ```json
  {"error": "Invalid Input"}
  ```

## Examples

### Valid Names (A-M)
```bash
curl "http://localhost:8080/hello-world?name=Alice"
# Response: 200 OK - {"message":"Hello Alice"}

curl "http://localhost:8080/hello-world?name=Alice%20Nam"
# Response: 200 OK - {"message":"Hello Alice Nam"}

curl "http://localhost:8080/hello-world?name=John%20Doe%20Jr"
# Response: 200 OK - {"message":"Hello John Doe Jr"}

curl "http://localhost:8080/hello-world?name=mark"
# Response: 200 OK - {"message":"Hello mark"}
```

### Invalid Names (N-Z)
```bash
curl "http://localhost:8080/hello-world?name=Nick"
# Response: 400 Bad Request - {"error":"Invalid Input"}

curl "http://localhost:8080/hello-world?name=Nancy%20Smith"
# Response: 400 Bad Request - {"error":"Invalid Input"}

curl "http://localhost:8080/hello-world?name=zoe"
# Response: 400 Bad Request - {"error":"Invalid Input"}
```

### Error Cases
```bash
curl "http://localhost:8080/hello-world"
# Response: 400 Bad Request - {"error":"Invalid Input"}

curl "http://localhost:8080/hello-world?name="
# Response: 400 Bad Request - {"error":"Invalid Input"}

curl "http://localhost:8080/hello-world?name=123John"
# Response: 400 Bad Request - {"error":"Invalid Input"}
```

## Requirements

- Go 1.18 or higher

## Running the Application

### Method 1: Direct Go Execution

1. **Clone or download the code to your workspace**

2. **Navigate to the project directory:**
   ```bash
   cd /hello-world-api
   ```

3. **Run the application:**
   ```bash
   go run .
   ```

   The server will start on port 8080:

4. **Test the endpoint:**
   ```bash
   curl "http://localhost:8080/hello-world?name=Alice"
   ```

5. **Graceful shutdown:**
    - Press `Ctrl+C` or send `SIGTERM`/`SIGINT` signal
    - Server will complete active requests and shutdown gracefully
    - 30-second timeout for completion of active requests

## Running the Tests

Run all unit tests:
```bash
go test
```

Run tests with verbose output:
```bash
go test -v
```

Run tests with coverage:
```bash
go test -cover
```

Run tests with detailed coverage report:
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Design Decisions & Assumptions

1. **Alphabet Range**: The first half is defined as A-M (inclusive), second half as N-Z (inclusive)

2. **Case Sensitivity**: The validation is case-insensitive - both 'Alice' and 'alice' are treated the same

3. **Input Validation**:
    - Names must start with a letter (not numbers or special characters)
    - Empty strings and missing parameters are treated as invalid
    - Leading/trailing whitespace is preserved in the response but affects validation

4. **Response Format**: All responses are in JSON format with appropriate HTTP status codes

5. **Error Handling**: All error cases return the same generic "Invalid Input" message as specified in requirements

6. **Multi-word Names**: Names with spaces are allowed (e.g., "John Doe") and the validation applies only to the first character