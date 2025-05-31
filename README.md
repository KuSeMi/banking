# Go Banking Project

## Project Description

This project is a Go-based banking application implementing the Hexagonal Architecture pattern (also known as Ports and Adapters). It provides core banking functionalities including customer management, account management, and transaction processing. The architecture ensures clear separation of concerns, making the application testable, maintainable, and adaptable to changing requirements.

## Project Structure

```
banking/
├── app/                   # Application entry point and HTTP handlers
├── domain/                # Domain models and repository interfaces
├── service/               # Business logic services
├── http-requests/         # HTTP test files for API endpoints
└── compose.yml            # Docker configuration
```

## Hexagonal Architecture Implementation

This project follows a strict hexagonal architecture pattern with the following components:

```
+------------------+     +-----------------+     +-------------------+
|                  |     |                 |     |                   |
|  Presentation    |     |    Domain       |     |  Infrastructure   |
|     Layer        |     |     Layer       |     |      Layer        |
|                  |     |                 |     |                   |
| (app/handlers.go)|     | (domain/*.go)   |     | (domain/*Db.go)   |
|                  |     | (service/*.go)  |     |                   |
+--------+---------+     +-------+---------+     +---------+---------+
         |                       |                         |
         v                       v                         v
+--------+---------+     +-------+---------+     +---------+---------+
|                  |     |                 |     |                   |
|  REST Controllers|---->|  Service Layer  |---->|  Repository      |
|  (CustomerHandler|     | (CustomerService|     |  Implementations |
|   AccountHandler)|     |  AccountService)|     | (CustomerRepoDb  |
|                  |     |                 |     |  AccountRepoDb)   |
+------------------+     +-----------------+     +-------------------+
```

### Components

1. **Domain Layer**
   - Contains business entities (`Customer`, `Account`, `Transaction`)
   - Defines repository interfaces (`CustomerRepository`, `AccountRepository`)
   - Pure business logic, independent of frameworks

2. **Service Layer**
   - Implements application use cases
   - Orchestrates business operations
   - Services: `CustomerService`, `AccountService`

3. **Infrastructure Layer**
   - Repository implementations (`CustomerRepositoryDb`, `AccountRepositoryDb`)
   - Database connection handling
   - External system integrations

4. **Presentation Layer**
   - HTTP handlers (`CustomerHandlers`, `AccountHandler`)
   - Request parsing and response formatting
   - Routing configuration

## Entity Relationship Diagram (ERD)

The application uses a MySQL database with the following structure:

```
+-------------+       +-------------+       +----------------+
|  Customers  |       |  Accounts   |       |  Transactions  |
+-------------+       +-------------+       +----------------+
| customer_id |<----->| account_id  |<----->| transaction_id |
| name        |       | customer_id |       | account_id     |
| city        |       | opening_date|       | amount         |
| zipcode     |       | account_type|       | transaction_type|
| date_of_birth|      | amount      |       | date            |
| status      |       | status      |       +-----------------+
+-------------+       +-------------+
```

### Relationships

- One **Customer** can have multiple **Accounts** (1:N)
- One **Account** can have multiple **Transactions** (1:N)

## Service Layer Diagram

```
+--------------------+      +--------------------+
|                    |      |                    |
|  CustomerService   |      |   AccountService   |
|                    |      |                    |
+--------+-----------+      +-----------+--------+
         |                              |
         v                              v
+--------+-----------+      +-----------+--------+
|                    |      |                    |
| CustomerRepository |      | AccountRepository  |
|                    |      |                    |
+--------------------+      +--------------------+
```

## Sequence Diagrams

### 1. Get All Customers

```
+---------+      +----------------+      +----------------+      +---------+
| Client  |      | CustomerHandler|      | CustomerService|      | Database|
+---------+      +----------------+      +----------------+      +---------+
     |                   |                      |                    |
     | GET /customers    |                      |                    |
     |------------------>|                      |                    |
     |                   | GetAllCustomers()    |                    |
     |                   |--------------------->|                    |
     |                   |                      | FindAll()          |
     |                   |                      |------------------->|
     |                   |                      |                    |
     |                   |                      |<-------------------|
     |                   |                      | Return Customers   |
     |                   |<---------------------|                    |
     |                   | Return JSON Response |                    |
     |<------------------|                      |                    |
     |                   |                      |                    |
```

### 2. Create New Account

```
+---------+      +----------------+      +----------------+      +---------+
| Client  |      | AccountHandler |      | AccountService |      | Database|
+---------+      +----------------+      +----------------+      +---------+
     |                   |                      |                    |
     | POST /customers/{id}/account|           |                    |
     |------------------>|                      |                    |
     |                   | NewAccount()         |                    |
     |                   |--------------------->|                    |
     |                   |                      | Save()             |
     |                   |                      |------------------->|
     |                   |                      |                    |
     |                   |                      |<-------------------|
     |                   |                      | Return Account ID  |
     |                   |<---------------------|                    |
     |                   | Return JSON Response |                    |
     |<------------------|                      |                    |
     |                   |                      |                    |
```

### 3. Make Transaction

```
+---------+      +----------------+      +----------------+      +---------+
| Client  |      | AccountHandler |      | AccountService |      | Database|
+---------+      +----------------+      +----------------+      +---------+
     |                   |                      |                    |
     | POST /customers/{id}/account/{id}       |                    |
     |------------------>|                      |                    |
     |                   | MakeTransaction()    |                    |
     |                   |--------------------->|                    |
     |                   |                      | GetAccount()       |
     |                   |                      |------------------->|
     |                   |                      |<-------------------|
     |                   |                      | Update Account     |
     |                   |                      |------------------->|
     |                   |                      |<-------------------|
     |                   |                      | Save Transaction   |
     |                   |                      |------------------->|
     |                   |                      |<-------------------|
     |                   |<---------------------|                    |
     |                   | Return JSON Response |                    |
     |<------------------|                      |                    |
     |                   |                      |                    |
```

## API Endpoints

| Method | Endpoint                                                       | Description                        |
|--------|----------------------------------------------------------------|------------------------------------|
| GET    | `/customers`                                                   | Get all customers                  |
| GET    | `/customers/{customer_id:[0-9]+}`                              | Get customer by ID                 |
| POST   | `/customers/{customer_id:[0-9]+}/account`                      | Create a new account for customer  |
| POST   | `/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}`  | Make a transaction                 |

## How to Run

### Prerequisites

- Docker and Docker Compose installed
- Go 1.15 or higher (for local development)

### Environment Variables

The application requires the following environment variables:

```
HOST=localhost
WEB_PORT=8080
MYSQL_HOST=localhost
MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_DATABASE=banking
```

### Running with Docker

```bash
docker-compose up --build -d
```

This will start the application and the MySQL database in containers. The API will be available at `http://localhost:8080`.

### Running Locally

1. Set up the required environment variables
2. Ensure MySQL is running and accessible
3. Run the application:

```bash
go run main.go
```

## Dependencies

This project has the following direct dependencies:

| Dependency                        | Description                                                                                                                        |
|-----------------------------------|------------------------------------------------------------------------------------------------------------------------------------|
| `github.com/gorilla/mux`          | A powerful HTTP router and URL matcher for building Go web servers. Used for routing and handling HTTP requests in the application.|
| `database/sql`                    | Go's standard library package for SQL database operations. Provides a generic interface for SQL database access.                   |
| `github.com/go-sql-driver/mysql`  | MySQL driver for Go's `database/sql` package. Enables the application to connect to MySQL databases.                               |
| `log`                             | Go's standard logging package for error and informational logging.                                                                 |
| `net/http`                        | Go's standard HTTP package for building web servers and handling HTTP requests.                                                    |
| `os`                              | Go's standard library for operating system functionality like environment variables access.                                        |

## Testing API Endpoints

You can test the API endpoints using the HTTP request files in the `http-requests/` directory.

With Visual Studio Code and the REST Client extension:
1. Open any *http file in the [http-requests] directory
2. Click on "Send Request" above each request definition

With curl:
```bash
# Get all customers
curl -X GET http://localhost:8080/customers

# Get customer by ID
curl -X GET http://localhost:8080/customers/2001

# Create new account
curl -X POST http://localhost:8080/customers/2001/account \
  -H "Content-Type: application/json" \
  -d '{"account_type":"savings","amount":5000}'

# Make transaction
curl -X POST http://localhost:8080/customers/2001/account/95470 \
  -H "Content-Type: application/json" \
  -d '{"amount":1000,"transaction_type":"withdrawal"}'
```
