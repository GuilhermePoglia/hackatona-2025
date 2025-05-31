#!/bin/bash

echo "🚀 Iniciando setup do backend..."

# Aguardar o banco ficar totalmente disponível
echo "⏳ Aguardando banco de dados..."
sleep 20

# Verificar se o banco está disponível
until pg_isready -h db -p 5432 -U admin -d postgres; do
  echo "⏳ Banco ainda não está pronto..."
  sleep 2
done

echo "✅ Banco de dados está pronto!"

# Aplicar migrations
echo "📦 Aplicando migrations..."
goose -dir infra/database/migrations postgres "host=db port=5432 user=admin password=admin dbname=postgres sslmode=disable" up

# Verificar se as migrations foram aplicadas com sucesso
if [ $? -eq 0 ]; then
    echo "✅ Migrations aplicadas com sucesso!"
    
    # Gerar modelos SQLBoiler
    echo "🔧 Gerando modelos SQLBoiler..."
    sqlboiler psql
    
    if [ $? -eq 0 ]; then
        echo "✅ Modelos gerados com sucesso!"
    else
        echo "❌ Erro ao gerar modelos"
    fi
else
    echo "❌ Erro ao aplicar migrations"
fi

echo "🚀 Iniciando aplicação..."
exec go run main.go
