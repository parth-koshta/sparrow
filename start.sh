#!/bin/sh

set -e

echo "run database migrations"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the application"
exec "$@"