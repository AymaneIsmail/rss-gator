#!/bin/bash

set -euo pipefail

DB_URL="postgres://postgres:password@localhost:5432/gator"
SCHEMA_DIR="sql/schema"

usage() {
  echo "Usage: $0 {up|down}"
  exit 1
}

# Vérifie qu’un argument a été fourni
if [[ $# -ne 1 ]]; then
  usage
fi

ACTION="$1"

case "$ACTION" in
  up|down)
    echo "📂 Changement de répertoire: $SCHEMA_DIR"
    cd "$SCHEMA_DIR"

    echo "🚀 Exécution des migrations: $ACTION"
    goose postgres "$DB_URL" "$ACTION"
    ;;
  *)
    echo "❌ Erreur: action inconnue '$ACTION'"
    usage
    ;;
esac
