#!/bin/sh

set -e

echo "Checking if /app/app.env exists"
if [ -f /app/app.env ]; then
    echo "File /app/app.env exists. Sourcing the environment variables."
    source /app/app.env
    echo "source: $DB_SOURCE"
else
    echo "File /app/app.env does not exist."
    exit 1
fi

echo "run database migrations"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the application"
exec "$@"
