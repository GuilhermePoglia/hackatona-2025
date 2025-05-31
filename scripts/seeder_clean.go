package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"hacka/core/models"
	"hacka/infra/database"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	db := database.StartDB()
	defer db.Close()

	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())

	empCount, _ := models.Employees().Count(ctx, db)
	resCount, _ := models.Resources().Count(ctx, db)
	actCount, _ := models.Activities().Count(ctx, db)
	feedCount, _ := models.Feedbacks().Count(ctx, db)

	if empCount > 0 || resCount > 0 || actCount > 0 || feedCount > 0 {
		fmt.Printf("Dados já existem no banco:\n")
		fmt.Printf("- Funcionários: %d\n", empCount)
		fmt.Printf("- Recursos: %d\n", resCount)
		fmt.Printf("- Atividades: %d\n", actCount)
		fmt.Printf("- Feedbacks: %d\n", feedCount)
		fmt.Printf("Deseja continuar e adicionar mais dados? (y/N): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Seeder cancelado.")
			os.Exit(0)
		}
	}

	fmt.Println("🧑‍💼 Inserindo funcionários...")
	employees := seedEmployees(ctx, db)

	fmt.Println("📦 Inserindo recursos...")
	seedResources(ctx, db)

	fmt.Println("📅 Inserindo atividades...")
	seedActivities(ctx, db)

	fmt.Println("💬 Inserindo feedbacks...")
	seedFeedbacks(ctx, db, employees)

	fmt.Println("\n🎉 Seeder completo concluído!")
	printSummary(ctx, db)
}

func seedEmployees(ctx context.Context, db boil.ContextExecutor) []*models.Employee {
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
		{
			Name:     null.StringFrom("Bruno Cardoso"),
			Email:    null.StringFrom("bruno.cardoso@empresa.com"),
			Position: null.StringFrom("Desenvolvedor Full Stack"),
			Balance:  null.Float64From(1400.00),
			Average:  null.Float64From(8.3),
		},
		{
			Name:     null.StringFrom("Camila Mendes"),
			Email:    null.StringFrom("camila.mendes@empresa.com"),
			Position: null.StringFrom("Data Analyst"),
			Balance:  null.Float64From(1100.50),
			Average:  null.Float64From(8.0),
		},
	}

	var insertedEmployees []*models.Employee
	for i, employee := range employees {
		err := employee.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Erro ao inserir funcionário %d: %v", i+1, err)
			continue
		}
		insertedEmployees = append(insertedEmployees, employee)
		fmt.Printf("✓ Funcionário inserido: %s (%s)\n",
			employee.Name.String,
			employee.Position.String)
	}

	return insertedEmployees
}

func seedResources(ctx context.Context, db boil.ContextExecutor) {
	resources := []*models.Resource{
		{
			Name:    null.StringFrom("Notebook Dell Inspiron 15"),
			Type:    null.StringFrom("hardware"),
			Midia:   null.StringFrom("Notebook para desenvolvimento com 16GB RAM e SSD 512GB"),
			Average: null.Float64From(4.5),
		},
		{
			Name:    null.StringFrom("Monitor LG UltraWide 29\""),
			Type:    null.StringFrom("hardware"),
			Midia:   null.StringFrom("Monitor ultrawide para maior produtividade"),
			Average: null.Float64From(4.8),
		},
		{
			Name:    null.StringFrom("Licença Adobe Creative Cloud"),
			Type:    null.StringFrom("software"),
			Midia:   null.StringFrom("Licença anual para suite Adobe (Photoshop, Illustrator, etc)"),
			Average: null.Float64From(4.7),
		},
		{
			Name:    null.StringFrom("Sala de Reunião A"),
			Type:    null.StringFrom("espaço"),
			Midia:   null.StringFrom("Sala com capacidade para 8 pessoas, equipada com projetor"),
			Average: null.Float64From(4.2),
		},
		{
			Name:    null.StringFrom("Workstation Mac Studio"),
			Type:    null.StringFrom("hardware"),
			Midia:   null.StringFrom("Workstation para trabalhos de design e desenvolvimento iOS"),
			Average: null.Float64From(4.9),
		},
		{
			Name:    null.StringFrom("Licença JetBrains IntelliJ"),
			Type:    null.StringFrom("software"),
			Midia:   null.StringFrom("IDE profissional para desenvolvimento Java/Kotlin"),
			Average: null.Float64From(4.6),
		},
		{
			Name:    null.StringFrom("Sala de Treinamento"),
			Type:    null.StringFrom("espaço"),
			Midia:   null.StringFrom("Auditório com capacidade para 30 pessoas"),
			Average: null.Float64From(4.3),
		},
		{
			Name:    null.StringFrom("Tablet iPad Pro"),
			Type:    null.StringFrom("hardware"),
			Midia:   null.StringFrom("Tablet para apresentações e trabalho móvel"),
			Average: null.Float64From(4.4),
		},
	}

	for i, resource := range resources {
		err := resource.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Erro ao inserir recurso %d: %v", i+1, err)
			continue
		}
		fmt.Printf("✓ Recurso inserido: %s (%s)\n",
			resource.Name.String,
			resource.Type.String)
	}
}

func seedActivities(ctx context.Context, db boil.ContextExecutor) {
	activities := []*models.Activity{
		{
			Name:        null.StringFrom("Reunião de Planejamento Sprint"),
			Description: null.StringFrom("Reunião semanal para planejamento das atividades do sprint"),
			Type:        null.StringFrom("meeting"),
		},
		{
			Name:        null.StringFrom("Treinamento de Segurança"),
			Description: null.StringFrom("Treinamento obrigatório sobre práticas de segurança no trabalho"),
			Type:        null.StringFrom("training"),
		},
		{
			Name:        null.StringFrom("Workshop de Inovação"),
			Description: null.StringFrom("Workshop para desenvolvimento de novas ideias e soluções inovadoras"),
			Type:        null.StringFrom("workshop"),
		},
		{
			Name:        null.StringFrom("Code Review Sessão"),
			Description: null.StringFrom("Sessão colaborativa de revisão de código"),
			Type:        null.StringFrom("development"),
		},
		{
			Name:        null.StringFrom("Daily Stand-up"),
			Description: null.StringFrom("Reunião diária de alinhamento da equipe"),
			Type:        null.StringFrom("meeting"),
		},
		{
			Name:        null.StringFrom("Treinamento Git Avançado"),
			Description: null.StringFrom("Capacitação em funcionalidades avançadas do Git"),
			Type:        null.StringFrom("training"),
		},
		{
			Name:        null.StringFrom("Retrospectiva do Sprint"),
			Description: null.StringFrom("Reunião para análise e melhoria contínua do processo"),
			Type:        null.StringFrom("meeting"),
		},
		{
			Name:        null.StringFrom("Hackathon Interno"),
			Description: null.StringFrom("Evento de 48h para desenvolvimento de projetos inovadores"),
			Type:        null.StringFrom("event"),
		},
		{
			Name:        null.StringFrom("Apresentação Projeto"),
			Description: null.StringFrom("Apresentação dos resultados do projeto para stakeholders"),
			Type:        null.StringFrom("presentation"),
		},
		{
			Name:        null.StringFrom("Mentoria Técnica"),
			Description: null.StringFrom("Sessão de mentoria para desenvolvimento técnico"),
			Type:        null.StringFrom("mentoring"),
		},
	}

	for i, activity := range activities {
		err := activity.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Erro ao inserir atividade %d: %v", i+1, err)
			continue
		}
		fmt.Printf("✓ Atividade inserida: %s (%s)\n",
			activity.Name.String,
			activity.Type.String)
	}
}

func seedFeedbacks(ctx context.Context, db boil.ContextExecutor, employees []*models.Employee) {
	if len(employees) < 2 {
		fmt.Println("⚠️  Precisa de pelo menos 2 funcionários para criar feedbacks")
		return
	}

	feedbackTexts := []string{
		"Excelente trabalho na implementação do novo sistema. Demonstrou grande conhecimento técnico.",
		"Muito colaborativo e sempre disposto a ajudar os colegas. Ótima comunicação.",
		"Entregou o projeto no prazo e com qualidade excepcional. Parabéns!",
		"Proativo na resolução de problemas e sempre busca as melhores soluções.",
		"Liderança exemplar durante o projeto. Conseguiu motivar toda a equipe.",
		"Atenção aos detalhes impressionante. Código sempre bem documentado.",
		"Ótima capacidade de aprendizado e adaptação às novas tecnologias.",
		"Feedback construtivo e sempre busca melhorar os processos da equipe.",
		"Apresentações claras e objetivas. Facilita muito o entendimento.",
		"Comprometimento excepcional com a qualidade do trabalho.",
		"Muito organizado e eficiente na gestão das tarefas.",
		"Criatividade na solução de problemas complexos.",
		"Excelente trabalho em equipe e integração com outros departamentos.",
		"Domínio técnico sólido e sempre compartilha conhecimento.",
		"Pontualidade e responsabilidade exemplares.",
	}

	ratings := []int{3, 4, 4, 5, 5, 4, 3, 5, 4, 4, 5, 3, 4, 5, 4}

	for i := 0; i < 15; i++ {
		senderIdx := rand.Intn(len(employees))
		receiverIdx := rand.Intn(len(employees))

		for receiverIdx == senderIdx {
			receiverIdx = rand.Intn(len(employees))
		}

		feedback := &models.Feedback{
			SenderID:    employees[senderIdx].ID,
			ReceiverID:  employees[receiverIdx].ID,
			Stars:       ratings[i],
			Description: null.StringFrom(feedbackTexts[i]),
		}

		err := feedback.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Erro ao inserir feedback %d: %v", i+1, err)
			continue
		}

		fmt.Printf("✓ Feedback inserido: %s → %s (⭐ %d)\n",
			employees[senderIdx].Name.String,
			employees[receiverIdx].Name.String,
			feedback.Stars)
	}
}

func printSummary(ctx context.Context, db boil.ContextExecutor) {
	empCount, _ := models.Employees().Count(ctx, db)
	resCount, _ := models.Resources().Count(ctx, db)
	actCount, _ := models.Activities().Count(ctx, db)
	feedCount, _ := models.Feedbacks().Count(ctx, db)

	fmt.Printf("\n📊 RESUMO FINAL:\n")
	fmt.Printf("==================\n")
	fmt.Printf("👥 Funcionários: %d\n", empCount)
	fmt.Printf("📦 Recursos: %d\n", resCount)
	fmt.Printf("📅 Atividades: %d\n", actCount)
	fmt.Printf("💬 Feedbacks: %d\n", feedCount)
	fmt.Printf("==================\n")
}
