MIGRATION_DIR=db/migration
POSTGRES_CONTAINER=postgres16

postgres:
	docker run --name $(POSTGRES_CONTAINER) -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it $(POSTGRES_CONTAINER) createdb --username=root --owner=root sparrow-dev

dropdb:
	docker exec -it $(POSTGRES_CONTAINER) dropdb sparrow-dev

migratecreate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $$name

migrateup:
	@if [ -z "$(n)" ]; then \
		migrate -path $(MIGRATION_DIR) -database "postgresql://root:secret@localhost:5432/sparrow-dev?sslmode=disable" -verbose up; \
	else \
		migrate -path $(MIGRATION_DIR) -database "postgresql://root:secret@localhost:5432/sparrow-dev?sslmode=disable" -verbose up $(n); \
	fi
	@$(MAKE) dumpschema

migratedown:
	@if [ -z "$(n)" ]; then \
		migrate -path $(MIGRATION_DIR) -database "postgresql://root:secret@localhost:5432/sparrow-dev?sslmode=disable" -verbose down; \
	else \
		migrate -path $(MIGRATION_DIR) -database "postgresql://root:secret@localhost:5432/sparrow-dev?sslmode=disable" -verbose down $(n); \
	fi
	@$(MAKE) dumpschema

sqlc:
	sqlc generate

dumpschema:
	@if docker ps -q -f name=$(POSTGRES_CONTAINER) > /dev/null; then \
		echo "Container $(POSTGRES_CONTAINER) exists. Running pg_dump..."; \
		docker exec $(POSTGRES_CONTAINER) pg_dump --schema-only --no-owner --file=/tmp/schema.sql sparrow-dev; \
		docker cp $(POSTGRES_CONTAINER):/tmp/schema.sql db/schema.sql; \
	else \
		echo "Container $(POSTGRES_CONTAINER) does not exist. Skipping pg_dump."; \
	fi

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockery --all

.PHONY: postgres createdb dropdb migratecreate migrateup migratedown sqlc test server mock
