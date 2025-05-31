package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"hacka/core/models"
	"hacka/infra/database"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	db := database.StartDB()
	defer db.Close()

	ctx := context.Background()

	count, err := models.Employees().Count(ctx, db)
	if err != nil {
		log.Fatal("Erro ao verificar funcionários existentes:", err)
	}

	if count > 0 {
		fmt.Printf("Já existem %d funcionários no banco. Deseja continuar? (y/N): ", count)
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Seeder cancelado.")
			os.Exit(0)
		}
	}

	employees := []*models.Employee{
		{
			Name:     null.StringFrom("João Silva"),
			Email:    null.StringFrom("joao.silva@empresa.com"),
			Position: null.StringFrom("Desenvolvedor Senior"),
			Balance:  null.Float64From(1250.50),
			Average:  null.Float64From(8.5),
		},
		{
			Name:     null.StringFrom("Maria Santos"),
			Email:    null.StringFrom("maria.santos@empresa.com"),
			Position: null.StringFrom("Product Manager"),
			Balance:  null.Float64From(2100.00),
			Average:  null.Float64From(9.2),
		},
		{
			Name:     null.StringFrom("Pedro Oliveira"),
			Email:    null.StringFrom("pedro.oliveira@empresa.com"),
			Position: null.StringFrom("Designer UX/UI"),
			Balance:  null.Float64From(890.75),
			Average:  null.Float64From(7.8),
		},
		{
			Name:     null.StringFrom("Ana Costa"),
			Email:    null.StringFrom("ana.costa@empresa.com"),
			Position: null.StringFrom("Desenvolvedor Junior"),
			Balance:  null.Float64From(750.00),
			Average:  null.Float64From(7.5),
		},
		{
			Name:     null.StringFrom("Carlos Ferreira"),
			Email:    null.StringFrom("carlos.ferreira@empresa.com"),
			Position: null.StringFrom("DevOps Engineer"),
			Balance:  null.Float64From(1800.25),
			Average:  null.Float64From(8.9),
		},
		{
			Name:     null.StringFrom("Lucia Rodrigues"),
			Email:    null.StringFrom("lucia.rodrigues@empresa.com"),
			Position: null.StringFrom("QA Analyst"),
			Balance:  null.Float64From(950.00),
			Average:  null.Float64From(8.1),
		},
		{
			Name:     null.StringFrom("Rafael Almeida"),
			Email:    null.StringFrom("rafael.almeida@empresa.com"),
			Position: null.StringFrom("Tech Lead"),
			Balance:  null.Float64From(2500.00),
			Average:  null.Float64From(9.5),
		},
		{
			Name:     null.StringFrom("Fernanda Lima"),
			Email:    null.StringFrom("fernanda.lima@empresa.com"),
			Position: null.StringFrom("Scrum Master"),
			Balance:  null.Float64From(1600.80),
			Average:  null.Float64From(8.7),
		},
	}

	fmt.Println("Inserindo funcionários...")
	for i, employee := range employees {
		err := employee.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Erro ao inserir funcionário %d: %v", i+1, err)
			continue
		}
		fmt.Printf("✓ Funcionário inserido: %s (%s)\n",
			employee.Name.String,
			employee.Position.String)
	}

	finalCount, err := models.Employees().Count(ctx, db)
	if err != nil {
		log.Fatal("Erro ao verificar total final:", err)
	}

	fmt.Printf("\n🎉 Seeder concluído! Total de funcionários no banco: %d\n", finalCount)
}
