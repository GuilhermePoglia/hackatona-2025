# API de Benefícios

Esta documentação descreve os endpoints da API de benefícios implementados no sistema.

## Estrutura do Benefit

```json
{
  "id": "uuid",
  "name": "string (obrigatório)",
  "description": "string (opcional)",
  "price": "float64 (obrigatório)",
  "image": "string (opcional)",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

## Endpoints Disponíveis

### 1. Listar Todos os Benefícios
- **GET** `/api/v1/benefits`
- **Descrição**: Retorna todos os benefícios ordenados por nome
- **Resposta**: Array de objetos Benefit

**Exemplo de Resposta:**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "Vale Refeição",
    "description": "Auxílio para alimentação durante o horário de trabalho",
    "price": 500.0,
    "image": "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=400&h=300&fit=crop",
    "created_at": "2025-05-31T12:00:00Z",
    "updated_at": "2025-05-31T12:00:00Z"
  }
]
```

### 2. Buscar Benefício por ID
- **GET** `/api/v1/benefits/:id`
- **Descrição**: Retorna um benefício específico pelo ID
- **Parâmetros**: 
  - `id` (UUID): ID do benefício
- **Resposta**: Objeto Benefit ou erro 404

### 3. Criar Novo Benefício
- **POST** `/api/v1/benefits`
- **Descrição**: Cria um novo benefício
- **Body**: JSON com os dados do benefício

**Exemplo de Request:**
```json
{
  "name": "Plano Odontológico",
  "description": "Cobertura odontológica completa",
  "price": 150.0,
  "image": "https://exemplo.com/imagem.jpg"
}
```

### 4. Atualizar Benefício
- **PUT** `/api/v1/benefits/:id`
- **Descrição**: Atualiza um benefício existente
- **Parâmetros**: 
  - `id` (UUID): ID do benefício
- **Body**: JSON com os novos dados

### 5. Deletar Benefício
- **DELETE** `/api/v1/benefits/:id`
- **Descrição**: Remove um benefício
- **Parâmetros**: 
  - `id` (UUID): ID do benefício
- **Resposta**: Status 204 em caso de sucesso

### 6. Buscar por Faixa de Preço
- **GET** `/api/v1/benefits/price-range?min=100&max=500`
- **Descrição**: Retorna benefícios dentro de uma faixa de preço
- **Query Parameters**:
  - `min` (float): Preço mínimo
  - `max` (float): Preço máximo
- **Resposta**: Array de objetos Benefit ordenados por preço

## Dados de Exemplo Disponíveis

O seeder inclui os seguintes benefícios:

1. **Vale Refeição** - R$ 500,00
2. **Plano de Saúde** - R$ 800,00
3. **Vale Transporte** - R$ 200,00
4. **Auxílio Creche** - R$ 400,00
5. **Gympass** - R$ 150,00
6. **Curso de Idiomas** - R$ 300,00
7. **Seguro de Vida** - R$ 100,00
8. **Day Off Aniversário** - R$ 0,00
9. **Home Office** - R$ 600,00
10. **Licença Paternidade Estendida** - R$ 0,00

Todos os benefícios incluem imagens do Unsplash e descrições detalhadas.

## Como Testar

1. Inicie o banco PostgreSQL
2. Execute as migrações: `goose -dir infra/database/migrations postgres "host=localhost port=5432 dbname=postgres user=admin sslmode=disable" up`
3. Execute o seeder: `go run cmd/seeder/seeder.go`
4. Inicie a API: `go run main.go`
5. Teste os endpoints usando curl ou Postman

**Exemplo de teste:**
```bash
# Listar todos os benefícios
curl http://localhost:8080/api/v1/benefits

# Buscar benefício por ID
curl http://localhost:8080/api/v1/benefits/{id}

# Buscar por faixa de preço
curl "http://localhost:8080/api/v1/benefits/price-range?min=100&max=500"
```
