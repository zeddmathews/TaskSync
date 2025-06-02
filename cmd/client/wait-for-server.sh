#!/bin/sh
# Wait for gRPC server before starting client
set -e

host="server"
port=50051

until nc -z "$host" "$port"; do
  echo "Waiting for gRPC server at $host:$port..."
  sleep 1
done

echo "Server is up - starting client"
exec ./client