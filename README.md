# Document Management System

A RESTful API service built with Go, Gin, and PostgreSQL for managing documents.

## Features

- Create, Read, Update, and Delete documents
- PostgreSQL database integration
- RESTful API endpoints
- Docker support

## Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Docker (optional)
- Docker Compose (optional)

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=document_system
DB_PORT=5432
PORT=8080
```

## API Endpoints

- `POST /api/documents` - Create a new document
- `GET /api/documents` - Get all documents
- `GET /api/documents/:id` - Get a document by ID
- `PUT /api/documents/:id` - Update a document
- `DELETE /api/documents/:id` - Delete a document

## Running Locally

1. Install dependencies:
```bash
go mod download
```

2. Run the application:
```bash
go run main.go
```

## Running with Docker

### Option 1: Using Docker Compose (Recommended)

1. Start all services:
```bash
docker-compose up -d
```

2. Stop all services:
```bash
docker-compose down
```

3. View logs:
```bash
docker-compose logs -f
```

### Option 2: Using Docker directly

1. Build the Docker image:
```bash
docker build -t document-system .
```

2. Run the container:
```bash
docker run -p 8080:8080 document-system
```

## Database Setup

<<<<<<< Updated upstream
- [@7lawa9111](https://github.com/7lawa9111) – Creator & Maintainer
- [@mariamkhaled99](https://github.com/mariamkhaled99) – Contributor
- [@RoadmannCoder](https://github.com/RoadmannCoder) – Contributor
- [@AbdallahAskar1](https://github.com/AbdallahAskar1) - Contributor
=======
1. Create a PostgreSQL database named `document_system`
2. The application will automatically create the required tables on startup
>>>>>>> Stashed changes

## API Request Examples

### Create Document
```bash
curl -X POST http://localhost:8080/api/documents \
  -H "Content-Type: application/json" \
  -d '{"title":"Sample Document","author":"John Doe","content":"This is a sample document"}'
```

### Get All Documents
```bash
curl http://localhost:8080/api/documents
```

### Get Document by ID
```bash
curl http://localhost:8080/api/documents/1
```

### Update Document
```bash
curl -X PUT http://localhost:8080/api/documents/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Document","author":"John Doe","content":"This is an updated document"}'
```

### Delete Document
```bash
curl -X DELETE http://localhost:8080/api/documents/1
``` 