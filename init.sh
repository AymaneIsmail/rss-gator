#!/bin/bash

set -e

CONTAINER_NAME="pg_database"
CONFIG_FILE="$HOME/.gatorconfig.json"

echo -e "🐘 Starting postgres v15 on Docker...\n"

# Supprimer le conteneur s’il existe déjà
docker rm -f -v "$CONTAINER_NAME" 2>/dev/null || true

# Lancer un nouveau conteneur PostgreSQL
docker run -d \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=password \
  --name "$CONTAINER_NAME" \
  postgres:15

# Créer le fichier de config gator
cat > "$CONFIG_FILE" <<EOF
{
  "db_url": "postgres://postgres:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "votre-nom-utilisateur"
}
EOF

echo "✅ Config file created at $CONFIG_FILE"
