MIGRATION_DIR=db/migration
POSTGRES_CONTAINER=postgres16
DB_SOURCE=postgresql://root:secret@localhost:5432/sparrow-dev?sslmode=disable

postgres:
	docker run --name $(POSTGRES_CONTAINER) --network=sparrow-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

createdb:
	docker exec -it $(POSTGRES_CONTAINER) createdb --username=root --owner=root sparrow-dev

dropdb:
	docker exec -it $(POSTGRES_CONTAINER) dropdb sparrow-dev

migratecreate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $$name

migrateup:
	@if [ -z "$(n)" ]; then \
		migrate -path $(MIGRATION_DIR) -database $(DB_SOURCE) -verbose up; \
	else \
		migrate -path $(MIGRATION_DIR) -database $(DB_SOURCE) -verbose up $(n); \
	fi
	@if [ -z "$$GITHUB_ACTIONS" ]; then \
		$(MAKE) dumpschema; \
	fi

migratedown:
	@if [ -z "$(n)" ]; then \
		migrate -path $(MIGRATION_DIR) -database $(DB_SOURCE) -verbose down; \
	else \
		migrate -path $(MIGRATION_DIR) -database $(DB_SOURCE) -verbose down $(n); \
	fi
	@if [ -z "$$GITHUB_ACTIONS" ]; then \
		$(MAKE) dumpschema; \
	fi

sqlc:
	sqlc generate

dumpschema:
	docker exec -it $(POSTGRES_CONTAINER) pg_dump --schema-only --no-owner --file=/tmp/schema.sql sparrow-dev
	docker cp $(POSTGRES_CONTAINER):/tmp/schema.sql db/schema.sql

test:
	go test -v -cover -timeout 30s -short ./...

server:
	go run main.go

mock:
	mockery --all

generate:
	make sqlc; make dumpschema; make mock

.PHONY: postgres createdb dropdb migratecreate migrateup migratedown sqlc test server mock generate redis
