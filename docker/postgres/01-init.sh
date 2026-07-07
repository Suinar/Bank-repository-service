#!/bin/bash
set -e

echo "Creating test database..."

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE bank_test;
EOSQL


echo "Applying migrations to bank..."

psql -v ON_ERROR_STOP=1 \
    --username "$POSTGRES_USER" \
    --dbname "$POSTGRES_DB" \
    -f /docker-entrypoint-initdb.d/migration/init.sql


echo "Applying migrations to bank_test..."

psql -v ON_ERROR_STOP=1 \
    --username "$POSTGRES_USER" \
    --dbname bank_test \
    -f /docker-entrypoint-initdb.d/migration/init_test.sql


echo "Database initialization completed!"