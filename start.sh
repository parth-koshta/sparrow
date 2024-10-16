#!/bin/sh

set -e

echo "run database migrations"
source /app/app.env
echo "source: $DB_SOURCE"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the application"
exec "$@"