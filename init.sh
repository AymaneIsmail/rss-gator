#!/bin/bash

set -e

CONTAINER_NAME="pg_database"

echo -e "Starting postgres v15 on Docker...\n"

docker rm -f -v "$CONTAINER_NAME" 2>/dev/null || true

docker run -d \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=password \
  --name "$CONTAINER_NAME" \
  postgres:15