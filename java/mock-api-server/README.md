# Mock API Server

A lightweight, standard-compliant mock REST server designed to simulate production environments for testing.

## Features
* **REST Compliant**: Follows standard HTTP methods and status codes.
* **JSON Payloads**: Utilizes uniform JSON formatting for requests and responses.
* **Resource-Oriented**: Logical URI structures representing distinct data entities.
* **State Simulation**: Supports full CRUD (Create, Read, Update, Delete) operations.

## Prerequisites
Before running the server, ensure you have the following installed:
* Java (21 or higher)

## Getting Started

### 1. Installation
Clone the repository and install the dependencies:
```bash
git clone https://github.com
cd mock-api-server
npm install
```

### 2. Configuration
Create a `.env` file in the root directory to customize your environment:
```env
PORT=3000
API_PREFIX=/api/v1
DELAY_MS=500
```

### 3. Running the Server
Start the mock server locally:
```bash
# Development mode with auto-reload
npm run dev

# Production mode
npm start
```
The server will be available at `http://localhost:3000/api/v1`.

## API Reference

### Base URL
```text
http://localhost:3000/api/v1
```

### Resources

#### Users (`/users`)

| Method | Endpoint | Description | Success Code |
| :--- | :--- | :--- | :--- |
| **GET** | `/users` | Retrieve a paginated list of users | `200 OK` |
| **GET** | `/users/:id` | Retrieve a single user by ID | `200 OK` |
| **POST** | `/users` | Create a new user record | `201 Created` |
| **PUT** | `/users/:id` | Completely replace an existing user | `200 OK` |
| **PATCH** | `/users/:id` | Partially update an existing user | `200 OK` |
| **DELETE** | `/users/:id` | Remove a user record permanently | `204 No Content` |

### Query Parameters (GET Requests)
* `_page`: Page number for pagination (e.g., `_page=1`)
* `_limit`: Number of items per page (e.g., `_limit=10`)
* `_sort`: Field name to sort by (e.g., `_sort=createdAt`)
* `_order`: Sorting order (`asc` or `desc`)

### Example Request & Response

#### Create a User
* **HTTP Method**: `POST`
* **Path**: `/api/v1/users`
* **Headers**: `Content-Type: application/json`

**Request Body:**
```json
{
  "name": "Jane Doe",
  "email": "jane.doe@example.com",
  "role": "admin"
}
```

**Response Body (`201 Created`):**
```json
{
  "id": "usr_9x8y7z6w",
  "name": "Jane Doe",
  "email": "jane.doe@example.com",
  "role": "admin",
  "createdAt": "2026-06-14T21:00:00Z"
}
```

## Error Handling
The server returns standard JSON error objects when requests fail.

### Standard Error Schema
```json
{
  "error": {
    "code": "RESOURCE_NOT_FOUND",
    "message": "The requested user with ID 'usr_123' does not exist.",
    "timestamp": "2026-06-14T21:05:00Z"
  }
}
```

### Common Status Codes Simulated
* `400 Bad Request`: Missing required fields or malformed JSON syntax.
* `401 Unauthorized`: Missing or invalid authentication token.
* `404 Not Found`: The resource or endpoint does not exist.
* `500 Internal Server Error`: Simulates unexpected server failures.
