.PHONY: help up down logs migrate generate clean

# Ajuda
help:
	@echo "🚀 Comandos disponíveis:"
	@echo ""
	@echo "🐳 Docker:"
	@echo "  make up           - Subir toda a stack (banco + backend + adminer)"
	@echo "  make down         - Parar todos os containers"
	@echo "  make logs         - Ver logs dos containers"
	@echo ""
	@echo "🔧 Banco de dados:"
	@echo "  make migrate      - Aplicar migrations (dentro do container)"
	@echo "  make generate     - Gerar modelos SQLBoiler (dentro do container)"
	@echo "  make seed         - Executar seeder para popular banco com dados de exemplo"
	@echo ""
	@echo "🧹 Utilitários:"
	@echo "  make clean        - Limpar containers e volumes"
	@echo "  make test-api     - Testar endpoints da API"

# Comandos principais
up:
	@echo "🚀 Subindo toda a stack..."
	docker-compose up --build

down:
	@echo "🛑 Parando todos os containers..."
	docker-compose down

logs:
	@echo "📋 Logs dos containers:"
	docker-compose logs -f

# Comandos de banco (executados dentro do container)
migrate:
	@echo "📦 Aplicando migrations..."
	docker-compose exec backend goose -dir infra/database/migrations postgres "host=db port=5432 user=admin password=admin dbname=postgres sslmode=disable" up

generate:
	@echo "🔧 Gerando modelos SQLBoiler..."
	docker-compose exec backend sqlboiler psql

seed:
	@echo "🌱 Executando seeder..."
	docker-compose exec backend go run scripts/seeder.go

# Limpeza
clean:
	@echo "🧹 Limpando containers e volumes..."
	docker-compose down -v
	docker system prune -f

# Testes
test-api:
	@echo "🧪 Testando endpoints da API..."
	@echo "📍 Testando ping:"
	curl -s http://localhost:8080/ping | jq
	@echo ""
	@echo "📍 Testando get all employees:"
	curl -s http://localhost:8080/api/v1/employees | jq '.count'
	@echo ""
	@echo "📍 Testando get employee by ID (primeiro funcionário):"
	ID=$$(curl -s http://localhost:8080/api/v1/employees | jq -r '.data[0].id') && curl -s http://localhost:8080/api/v1/employees/$$ID | jq '.data.name'
