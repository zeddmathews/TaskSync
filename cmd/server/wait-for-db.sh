#!/bin/sh
# Wait for PostgreSQL before starting server
set -e

host="db"
port=5432

until nc -z "$host" "$port"; do
  echo "Waiting for PostgreSQL at $host:$port..."
  sleep 1
done

echo "PostgreSQL is up - starting server"
exec ./server