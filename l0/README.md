# Order Demo Service (wb-tech-l0)

## Description

This is a demo microservice written in Go. It shows how to work with Kafka, PostgreSQL, and caching. The service gets order data from a message queue (Kafka), saves it to a database (PostgreSQL), caches it in memory for fast access, and provides an HTTP API and a simple web interface to view order info by ID.

---

## Key Features

- **Gets orders from Kafka**: Listens to a Kafka topic and processes incoming order messages.
- **Validates data**: Checks if messages are valid. Invalid ones are logged and ignored.
- **Saves orders to PostgreSQL**: Stores valid orders in the database.
- **Caches orders**: Recently viewed orders are kept in cache for faster access.
- **HTTP API**: Get order details by UID (`GET /api/order/<order_uid>`), returns JSON.
- **Web interface**: Simple page where you can enter an order ID and see its info.
- **Flexible setup**: Easy to add new brokers, storage, or cache types using the registry.
- **Graceful shutdown**: Closes all connections properly when stopping.
- **Logging**: Detailed logs for debugging and monitoring.

---

## Project Structure

```
.
├── cmd/                # Main entry point (main.go)
├── internal/
│   ├── app/            # App startup and lifecycle management
│   ├── broker/         # Message brokers
│   ├── cache/          # Caches
│   ├── config/         # Loads and validates config from env
│   ├── logger/         # Logger interface
│   ├── models/         # Data models
│   ├── registry/       # Service registry
│   ├── server/         # HTTP server, router, handlers
│   └── storage/        # Databases
├── migrations/         # SQL migrations for tables
├── frontend/           # Web interface (HTML, nginx)
├── docker-compose.yml  # Docker Compose for local setup
├── Dockerfile          # Backend Dockerfile
└── README.md           # This file
```

---

## How to Run

### 1. Clone the repo

```sh
git clone https://github.com/w3hhh-m/wb-tech
cd l0
```

### 2. Create `.env` file

Use `.env.example` as a template.

### 3. Start services with Docker Compose

```sh
docker-compose up --build
```

- Backend: `http://localhost:8080`
- Frontend (web interface): `http://localhost:8081`

### 4. Test it

- To check the API:  
  `GET http://localhost:8080/api/order/<order_uid>`
- To check the web interface:  
  Open `http://localhost:8081`, enter an order_uid and see the data.

---

## How It Works (Order Flow)

1. **Get order**:  
   The service listens to Kafka for new order messages.

2. **Validate**:  
   Messages are parsed and checked. Invalid ones are logged and skipped.

3. **Save**:  
   Valid orders are saved to PostgreSQL (uses transactions to avoid data loss).

4. **HTTP API**:  
   When you call `/api/order/<order_uid>`:
    - First checks the cache.
    - If not found, gets from DB, adds to cache, and returns.
    - If still not found, returns 404.

---

## Registry

The project uses a registry pattern for services (broker, storage, cache). This makes it easy to add new implementations (like a different cache or broker) – just register them and set the type in environment variables.

---

## Graceful Shutdown

The service stops properly:
- Catches shutdown signals (SIGINT, SIGTERM).
- Closes HTTP server, DB, broker, and cache connections.
- Waits for operations to finish (with timeout `SHUTDOWN_TIMEOUT`).

---

## Migrations

Migrations for creating tables are in `migrations/`.  
They run automatically when starting with docker-compose.

---

## Web Interface

- Located in `frontend/`.
- Runs on nginx, calls backend API.
- Lets you enter an order_uid and see its details.

---

## Example Request

```
GET http://localhost:8080/api/order/b563feb7b2b84b6test
```

Response – JSON with full order details.

---

## Future Improvements

- Easy to add new brokers, caches, or storage (via registry).
- Could add log collection, metrics, and monitoring.