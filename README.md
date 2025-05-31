# Hackatona 2025 - API

Este projeto utiliza Go com Gin, PostgreSQL e SQLBoiler para geração automática de modelos.

## 🚀 Setup Super Simples

### **1. Subir toda a stack:**
```bash
docker-compose up
```

### **2. Em outro terminal, configurar o banco:**
```bash
make migrate      # Aplicar migrations
make generate     # Gerar modelos SQLBoiler
```

**Pronto! 🎉**

## 🌐 Acessos

- **🚀 API Backend**: http://localhost:8080
- **🔧 Adminer (DB Admin)**: http://localhost:8081
  - **Server**: `db`
  - **Username**: `admin` 
  - **Password**: `admin`
  - **Database**: `postgres`

## 📋 Comandos Disponíveis

```bash
# Docker
docker-compose up    # Subir tudo
docker-compose down  # Parar tudo

# OU usar Makefile
make up             # Subir tudo
make down           # Parar tudo  
make logs           # Ver logs
make migrate        # Aplicar migrations
make generate       # Gerar modelos
make clean          # Limpar tudo

# Seeder (Popular banco com dados de exemplo)
go run cmd/seeder/seeder.go    # Executar seeder
```

## ❓ **Por que SEM entrypoint?**

**Mais simples e flexível:**
- ✅ Você controla quando aplicar migrations
- ✅ Pode regenerar modelos quando quiser
- ✅ Mais fácil debugar problemas
- ✅ Não precisa esperar scripts automáticos

## 🌱 Seeder - Dados de Exemplo

O seeder popula o banco com dados realísticos para desenvolvimento e testes:

### 📊 Dados Criados:
- **👥 Funcionários**: 10 funcionários com cargos variados (Desenvolvedor, Product Manager, Designer, etc.)
- **📦 Recursos**: 8 recursos (notebooks, monitores, licenças de software, salas)
- **📅 Atividades**: 10 atividades (reuniões, treinamentos, workshops, etc.)
- **💬 Feedbacks**: 15 feedbacks aleatórios entre funcionários com ratings de 1-5 estrelas

### 🚀 Como Usar:
```bash
# Executar o seeder
go run cmd/seeder/seeder.go

# O seeder detecta automaticamente se já existem dados
# e pergunta se você quer adicionar mais
```

### ✨ Funcionalidades:
- ✅ **Inteligente**: Detecta dados existentes e pergunta antes de duplicar
- ✅ **Realístico**: Dados com nomes, descrições e relacionamentos coerentes
- ✅ **Automático**: Calcula automaticamente médias e balanços de feedback
- ✅ **Resumo**: Mostra estatísticas finais após execução

## 🗄️ Estrutura do Banco

O projeto possui as seguintes tabelas:
- `employee` - Funcionários
- `resource` - Recursos  
- `activity` - Atividades

## 🔧 SQLBoiler

### Gerando Modelos

Após fazer mudanças nas migrations:

**Docker:**
```bash
# Entrar no container do backend
docker exec -it hacka_backend_dev sh
# Aplicar migrations e gerar modelos
goose -dir infra/database/migrations postgres "host=db port=5432 user=admin password=admin dbname=postgres sslmode=disable" up
sqlboiler psql
```

**Local:**
```bash
make db-migrate     # Aplicar mudanças no banco
make db-generate    # Regenerar modelos
```

### Usando os Modelos

Após gerar os modelos, você pode usar SQLBoiler assim:

```go
import (
    "hacka/core/models"
    "github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/null/v8"
)

// Criar um employee
employee := &models.Employee{
    Name:     null.StringFrom("João Silva"),
    Email:    null.StringFrom("joao@example.com"),
    Position: null.StringFrom("Developer"),
    Balance:  null.FloatFrom(1000.0),
}

err := employee.Insert(ctx, db, boil.Infer())

// Buscar employees
employees, err := models.Employees().All(ctx, db)

// Buscar com filtros
devs, err := models.Employees(
    qm.Where("position = ?", "Developer"),
).All(ctx, db)
```

## 🛣️ Rotas da API

### Employees
- `GET /api/v1/employees` - Listar todos
- `GET /api/v1/employees/:id` - Buscar por ID
- `POST /api/v1/employees` - Criar novo
- `PUT /api/v1/employees/:id` - Atualizar
- `DELETE /api/v1/employees/:id` - Deletar
- `GET /api/v1/employees/position/:position` - Buscar por cargo

### Health Check
- `GET /ping` - Status da API

## 📁 Estrutura do Projeto

```
├── api/                    # Controllers e rotas
│   ├── controllers/        # Controllers da API
│   ├── api.go             # Setup das rotas
│   └── app.go             # Configuração da aplicação
├── core/                  # Lógica de negócio
│   ├── models/            # Modelos gerados pelo SQLBoiler
│   └── services/          # Services da aplicação
├── infra/                 # Infraestrutura
│   └── database/          # Configuração do banco
│       ├── migrations/    # Migrations do banco
│       └── connection.go  # Conexão com PostgreSQL
├── scripts/               # Scripts utilitários
├── examples/              # Exemplos de uso
├── docker-compose.yml     # PostgreSQL
├── sqlboiler.toml        # Configuração do SQLBoiler
├── Makefile              # Comandos automatizados
└── main.go               # Ponto de entrada
```

## 🔐 Variáveis de Ambiente

```bash
API_PORT=8080              # Porta da API (opcional)
JWT_SECRET=my_secret_key   # Chave JWT (opcional)
```

## 🐛 Troubleshooting

### Banco não conecta
```bash
make db-down
make db-up
# Aguarde alguns segundos
make db-migrate
```

### Modelos não foram gerados
```bash
make clean
make db-generate
```

### Erro de compilação
```bash
go mod tidy
make clean
make db-setup
```

## 📚 Documentação

- [SQLBoiler](https://github.com/volatiletech/sqlboiler)
- [Gin](https://gin-gonic.com/)
- [Goose Migrations](https://github.com/pressly/goose)
