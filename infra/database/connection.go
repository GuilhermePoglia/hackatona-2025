package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func StartDB() *sql.DB {
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefaultInt("DB_PORT", 5432)
	user := getEnvOrDefault("DB_USER", "admin")
	password := getEnvOrDefault("DB_PASSWORD", "admin")
	dbname := getEnvOrDefault("DB_NAME", "postgres")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Printf("Conectando ao banco: host=%s port=%d dbname=%s user=%s", host, port, dbname, user)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")
	return db
}
