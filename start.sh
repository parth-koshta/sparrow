#!/bin/sh

set -e

echo "source the environment variables"
source /app/app.env

echo "run database migrations"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the application"
exec "$@"
