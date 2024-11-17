# Sparrow - BuzzlyAI Backend Service

Sparrow is the backend service for BuzzlyAI, an AI-powered platform that helps users create and manage engaging social media posts. Users can generate post suggestions using AI, edit and save them, and publish or schedule posts to LinkedIn effortlessly.

---

## Features

- **AI-Generated Post Suggestions**: Generate content ideas for LinkedIn based on user inputs.
- **Post Management**: Save, edit, and organize posts.
- **Scheduling**: Schedule posts to go live on LinkedIn at a specified time.
- **Social Account Integration**: Seamlessly link LinkedIn accounts to publish content.
- **Admin and Monitoring**: Monitor Redis task queues with Asynqmon.

---

## Project Setup

### Prerequisites

- Docker & Docker Compose
- Go (>= 1.21)
- PostgreSQL & Redis
- `sqlc` for database operations
- `mockery` for generating test mocks
- `migrate` for managing database migrations

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/sparrow.git
   cd sparrow
   ```
   
2. Create a Docker network:
    ```bash
    docker network create sparrow-network
    ```

3. Start PostgreSQL:
    ```bash
    make postgres
    ```

4. Start Redis:
    ```bash
    make redis
    ```

5. Create the development database:
    ```bash
    make createdb
    ```

6. Apply database migrations:
    ```bash
    make migrateup
    ```

7. Start worker UI (for monitoring Redis queues):
    ```bash
    make worker-ui
    ```

8. Run the server:
    ```bash
    make server
    ```

## Makefile Commands

| Command        | Description                                       |
|----------------|---------------------------------------------------|
| `postgres`     | Starts a PostgreSQL container.                   |
| `redis`        | Starts a Redis container.                        |
| `worker-ui`    | Runs the Asynqmon UI for monitoring Redis task queues. |
| `createdb`     | Creates the development database.                |
| `dropdb`       | Drops the development database.                  |
| `migrateup`    | Applies all database migrations.                 |
| `migratedown`  | Reverts database migrations.                     |
| `migratecreate`| Creates a new migration file with a specified name. |
| `sqlc`         | Generates SQLC Go code for database queries.     |
| `dumpschema`   | Dumps the database schema into a file.           |
| `test`         | Runs all tests.                                  |
| `server`       | Runs the application server.                     |
| `mock`         | Generates mocks using Mockery.                   |
| `generate`     | Runs SQLC, dumps the schema, and generates mocks. |


## Testing
  Run unit tests with:
  ```bash
  make test
  ```
