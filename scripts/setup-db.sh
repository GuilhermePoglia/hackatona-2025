#!/bin/bash

# Script para setup do banco de dados e geração de modelos

echo "🚀 Iniciando setup do banco de dados..."

# Aplicar migrations
echo "📦 Aplicando migrations..."
goose -dir infra/database/migrations postgres "host=localhost port=5432 user=admin password=admin dbname=postgres sslmode=disable" up

# Gerar modelos com SQLBoiler
echo "🔧 Gerando modelos com SQLBoiler..."
sqlboiler psql

echo "✅ Setup concluído!"
echo "📁 Modelos gerados em: core/models/"
