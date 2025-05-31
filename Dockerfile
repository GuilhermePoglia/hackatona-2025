# Build stage
FROM golang:1.21-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git ca-certificates tzdata

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates

# Criar usuário não-root
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /root/

# Copiar binário da stage anterior
COPY --from=builder /app/main .

# Mudar proprietário do arquivo
RUN chown appuser:appgroup main

# Usar usuário não-root
USER appuser

# Expor porta
EXPOSE 8080

# Comando para executar
CMD ["./main"]
