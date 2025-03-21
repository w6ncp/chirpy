# Chirpy
Chripy is a backend Go project used during the Boot.dev HTTP Server lessons. It is able to handle users and "chirps" the users create, but has no front end.

## Required

This project requires a Golang installation and postgres SQL server running.

## Installation

Inside a Go module:

```bash
go install github.com/w6ncp/chirpy
```

## Configuration

Create a `.env` file in your project root with the following structure:

```env
DB_URL="postgres://username:password@localhost:PORT/chirpy?sslmode=disable"
FILEPATH_ROOT ="./www"
TESTING_PORT = "8080"
PLATFORM = "dev or live"
TOKEN_SECRET = "secret_token"
POLKA_KEY = "polka_API_key"
```
## User Structure
Structure of a User Object returned by Chirpy
```json
{
    "id": "UUID",
    "created_at": "TIMESTAMP",
    "updated_at": "TIMESTAMP",
    "email": "STRING",
    "is_chirpy_red": "BOOL",
    "token": "STRING",
    "refresh_token": "STRING"
}
```
