package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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

	empCount, _ := models.Employees().Count(ctx, db)
	resCount, _ := models.Resources().Count(ctx, db)
	actCount, _ := models.Activities().Count(ctx, db)
	feedCount, _ := models.Feedbacks().Count(ctx, db)

	if empCount > 0 || resCount > 0 || actCount > 0 || feedCount > 0 {
		fmt.Printf("Data already exists in database:\n")
		fmt.Printf("- Employees: %d\n", empCount)
		fmt.Printf("- Resources: %d\n", resCount)
		fmt.Printf("- Activities: %d\n", actCount)
		fmt.Printf("- Feedbacks: %d\n", feedCount)
		fmt.Printf("Continue and add more data? (y/N): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Seeder cancelled.")
			os.Exit(0)
		}
	}

	fmt.Println("🧑‍💼 Inserting employees...")
	employees := seedEmployees(ctx, db)

	fmt.Println("📦 Inserting resources...")
	seedResources(ctx, db)

	fmt.Println("📅 Inserting activities...")
	seedActivities(ctx, db)

	fmt.Println("🎁 Inserting benefits...")
	seedBenefits(ctx, db)

	fmt.Println("💬 Inserting feedbacks...")
	seedFeedbacks(ctx, db, employees)

	fmt.Println("\n🎉 Seeder complete!")
	printSummary(ctx, db)
}

func seedEmployees(ctx context.Context, db boil.ContextExecutor) []*models.Employee {
	employees := []*models.Employee{
		{
			Name:     null.StringFrom("John Silva"),
			Email:    null.StringFrom("john.silva@company.com"),
			Position: null.StringFrom("Senior Developer"),
			Balance:  null.Float64From(1250.50),
			Average:  null.Float64From(4.2),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/D4D03AQG7UaCKSF9rAg/profile-displayphoto-shrink_800_800/B4DZOuEKbrHUAg-/0/1733792148042?e=1753920000&v=beta&t=ot_27yUBbk7hEYHcsxTJp9bhMW-d2MOx2OvwiVB1MKw"),
		},
		{
			Name:     null.StringFrom("Maria Santos"),
			Email:    null.StringFrom("maria.santos@company.com"),
			Position: null.StringFrom("Product Manager"),
			Balance:  null.Float64From(2100.00),
			Average:  null.Float64From(4.6),
			Midia:    null.StringFrom("https://portal.pucrs.br/wp-content/uploads/2024/06/triple-awards-tecnopuc-1.jpeg"),
		},
		{
			Name:     null.StringFrom("Peter Oliveira"),
			Email:    null.StringFrom("peter.oliveira@company.com"),
			Position: null.StringFrom("UX/UI Designer"),
			Balance:  null.Float64From(890.75),
			Average:  null.Float64From(3.9),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/C4E03AQFFF1FSoqS-jw/profile-displayphoto-shrink_800_800/profile-displayphoto-shrink_800_800/0/1647298540286?e=1753920000&v=beta&t=MMdF9QiqHBRXjhw6fZ2Kzq1TkcyPD0bJSz28ZFIJBmc"),
		},
		{
			Name:     null.StringFrom("Anna Costa"),
			Email:    null.StringFrom("anna.costa@company.com"),
			Position: null.StringFrom("Junior Developer"),
			Balance:  null.Float64From(750.00),
			Average:  null.Float64From(3.7),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/C4E03AQG8HGulK6w8og/profile-displayphoto-shrink_800_800/profile-displayphoto-shrink_800_800/0/1582150268818?e=1753920000&v=beta&t=Ict4pJBMAoWDsMrZGCnYS9GezIdOcrStBIRIs09ixnw"),
		},
		{
			Name:     null.StringFrom("Carlos Ferreira"),
			Email:    null.StringFrom("carlos.ferreira@company.com"),
			Position: null.StringFrom("DevOps Engineer"),
			Balance:  null.Float64From(1800.25),
			Average:  null.Float64From(4.4),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/C4D03AQH-69SMOHqOOw/profile-displayphoto-shrink_800_800/profile-displayphoto-shrink_800_800/0/1656955167715?e=1753920000&v=beta&t=WduT-NE1acQIKcqdY2Ia2Pc085gLQqEwDIn-WV2xFes"),
		},
		{
			Name:     null.StringFrom("Lucy Rodrigues"),
			Email:    null.StringFrom("lucy.rodrigues@company.com"),
			Position: null.StringFrom("QA Analyst"),
			Balance:  null.Float64From(950.00),
			Average:  null.Float64From(4.0),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/D4D03AQHcwDN2s22pAA/profile-displayphoto-shrink_800_800/profile-displayphoto-shrink_800_800/0/1700490727729?e=1753920000&v=beta&t=spO6ekf4ssq4rVJ2NPlgjgVpsuOFEQBUDlm4eaBjv_s"),
		},
		{
			Name:     null.StringFrom("Rafael Almeida"),
			Email:    null.StringFrom("rafael.almeida@company.com"),
			Position: null.StringFrom("Tech Lead"),
			Balance:  null.Float64From(2500.00),
			Average:  null.Float64From(4.7),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/D4D03AQHyYHw9p1_TAQ/profile-displayphoto-shrink_800_800/profile-displayphoto-shrink_800_800/0/1718367750039?e=1753920000&v=beta&t=b5mcAX2dw49Pg108sj_GxYXrODq6aYWzIDtLuFFNbhI"),
		},
		{
			Name:     null.StringFrom("Fernanda Lima"),
			Email:    null.StringFrom("fernanda.lima@company.com"),
			Position: null.StringFrom("Scrum Master"),
			Balance:  null.Float64From(1600.80),
			Average:  null.Float64From(4.3),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/D4D03AQEFSpIOh19rAQ/profile-displayphoto-shrink_800_800/profile-displayphoto-shrink_800_800/0/1731270916510?e=1753920000&v=beta&t=X73aS4-eiQrrhlg1kkMtDPTAAHNu0K6rbyff4q5OrCs"),
		},
		{
			Name:     null.StringFrom("Bruno Cardoso"),
			Email:    null.StringFrom("bruno.cardoso@company.com"),
			Position: null.StringFrom("Full Stack Developer"),
			Balance:  null.Float64From(1400.00),
			Average:  null.Float64From(4.1),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/D4D03AQG5gjwSPA8vyw/profile-displayphoto-shrink_800_800/profile-displayphoto-shrink_800_800/0/1713146677750?e=1753920000&v=beta&t=cdypxg2R0_Z0OKLV0f7j5wkeWyXIbwUuYe243H16JnY"),
		},
		{
			Name:     null.StringFrom("Camila Mendes"),
			Email:    null.StringFrom("camila.mendes@company.com"),
			Position: null.StringFrom("Data Analyst"),
			Balance:  null.Float64From(1100.50),
			Average:  null.Float64From(4.0),
			Midia:    null.StringFrom("https://media.licdn.com/dms/image/v2/C4D03AQFZJP5J8k3E_Q/profile-displayphoto-shrink_800_800/profile-displayphoto-shrink_800_800/0/1610479651047?e=1753920000&v=beta&t=VnJ5oQ85xU1H83_4z2OFW4tIhh0J9s8k8fF4lPMGDhc"),
		},
	}

	var insertedEmployees []*models.Employee
	for i, employee := range employees {
		err := employee.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Error inserting employee %d: %v", i+1, err)
			continue
		}
		insertedEmployees = append(insertedEmployees, employee)
		fmt.Printf("✓ Employee inserted: %s (%s)\n",
			employee.Name.String,
			employee.Position.String)
	}

	return insertedEmployees
}

func seedResources(ctx context.Context, db boil.ContextExecutor) {
	resources := []*models.Resource{
		{
			Name:    null.StringFrom("Dell Inspiron 15 Laptop"),
			Type:    null.StringFrom("hardware"),
			Midia:   null.StringFrom("https://images.unsplash.com/photo-1496181133206-80ce9b88a853?w=400&h=300&fit=crop"),
			Average: null.Float64From(4.5),
		},
		{
			Name:    null.StringFrom("LG UltraWide 29\" Monitor"),
			Type:    null.StringFrom("hardware"),
			Midia:   null.StringFrom("https://images.unsplash.com/photo-1527864550417-7fd91fc51a46?w=400&h=300&fit=crop"),
			Average: null.Float64From(4.8),
		},
		{
			Name:    null.StringFrom("Adobe Creative Cloud License"),
			Type:    null.StringFrom("software"),
			Midia:   null.StringFrom("https://images.unsplash.com/photo-1611224923853-80b023f02d71?w=400&h=300&fit=crop"),
			Average: null.Float64From(4.7),
		},
		{
			Name:    null.StringFrom("Meeting Room A"),
			Type:    null.StringFrom("space"),
			Midia:   null.StringFrom("https://images.unsplash.com/photo-1497366216548-37526070297c?w=400&h=300&fit=crop"),
			Average: null.Float64From(4.2),
		},
		{
			Name:    null.StringFrom("Mac Studio Workstation"),
			Type:    null.StringFrom("hardware"),
			Midia:   null.StringFrom("https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=400&h=300&fit=crop"),
			Average: null.Float64From(4.9),
		},
		{
			Name:    null.StringFrom("JetBrains IntelliJ License"),
			Type:    null.StringFrom("software"),
			Midia:   null.StringFrom("https://images.unsplash.com/photo-1555066931-4365d14bab8c?w=400&h=300&fit=crop"),
			Average: null.Float64From(4.6),
		},
		{
			Name:    null.StringFrom("Training Room"),
			Type:    null.StringFrom("space"),
			Midia:   null.StringFrom("https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=400&h=300&fit=crop"),
			Average: null.Float64From(4.3),
		},
		{
			Name:    null.StringFrom("iPad Pro Tablet"),
			Type:    null.StringFrom("hardware"),
			Midia:   null.StringFrom("https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=400&h=300&fit=crop"),
			Average: null.Float64From(4.4),
		},
	}

	for i, resource := range resources {
		err := resource.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Error inserting resource %d: %v", i+1, err)
			continue
		}
		fmt.Printf("✓ Resource inserted: %s (%s)\n",
			resource.Name.String,
			resource.Type.String)
	}
}

func seedActivities(ctx context.Context, db boil.ContextExecutor) {
	activities := []*models.Activity{
		{
			Name:        null.StringFrom("Sprint Planning Meeting"),
			Description: null.StringFrom("Weekly meeting for sprint activity planning"),
			Type:        null.StringFrom("meeting"),
		},
		{
			Name:        null.StringFrom("Security Training"),
			Description: null.StringFrom("Mandatory training on workplace safety practices"),
			Type:        null.StringFrom("training"),
		},
		{
			Name:        null.StringFrom("Innovation Workshop"),
			Description: null.StringFrom("Workshop for developing new ideas and innovative solutions"),
			Type:        null.StringFrom("workshop"),
		},
		{
			Name:        null.StringFrom("Code Review Session"),
			Description: null.StringFrom("Collaborative code review session"),
			Type:        null.StringFrom("development"),
		},
		{
			Name:        null.StringFrom("Daily Stand-up"),
			Description: null.StringFrom("Daily team alignment meeting"),
			Type:        null.StringFrom("meeting"),
		},
		{
			Name:        null.StringFrom("Advanced Git Training"),
			Description: null.StringFrom("Training on advanced Git features"),
			Type:        null.StringFrom("training"),
		},
		{
			Name:        null.StringFrom("Sprint Retrospective"),
			Description: null.StringFrom("Meeting for process analysis and continuous improvement"),
			Type:        null.StringFrom("meeting"),
		},
		{
			Name:        null.StringFrom("Internal Hackathon"),
			Description: null.StringFrom("48-hour event for developing innovative projects"),
			Type:        null.StringFrom("event"),
		},
		{
			Name:        null.StringFrom("Project Presentation"),
			Description: null.StringFrom("Presentation of project results to stakeholders"),
			Type:        null.StringFrom("presentation"),
		},
		{
			Name:        null.StringFrom("Technical Mentorship"),
			Description: null.StringFrom("Mentoring session for technical development"),
			Type:        null.StringFrom("mentoring"),
		},
	}

	for i, activity := range activities {
		err := activity.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Error inserting activity %d: %v", i+1, err)
			continue
		}
		fmt.Printf("✓ Activity inserted: %s (%s)\n",
			activity.Name.String,
			activity.Type.String)
	}
}

func seedBenefits(ctx context.Context, db boil.ContextExecutor) {
	benefits := []*models.Benefit{
		{
			Name:        "Meal Voucher",
			Description: null.StringFrom("Food assistance during work hours"),
			Price:       500.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Health Insurance",
			Description: null.StringFrom("Complete medical coverage for employee and dependents"),
			Price:       800.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1559757148-5c350d0d3c56?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Transportation Voucher",
			Description: null.StringFrom("Commuting assistance from home to work"),
			Price:       200.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1544620347-c4fd4a3d5957?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Childcare Assistance",
			Description: null.StringFrom("Daycare expense reimbursement for children up to 5 years old"),
			Price:       400.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1503454537195-1dcabb73ffb9?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Gym Membership",
			Description: null.StringFrom("Access to gym network and physical activities"),
			Price:       150.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Language Courses",
			Description: null.StringFrom("Subsidy for English, Spanish or other language courses"),
			Price:       300.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1434030216411-0b793f4b4173?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Life Insurance",
			Description: null.StringFrom("Life insurance coverage for the employee"),
			Price:       100.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1450101499163-c8848c66ca85?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Birthday Day Off",
			Description: null.StringFrom("Paid day off on birthday"),
			Price:       0.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1513475382585-d06e58bcb0e0?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Home Office Setup",
			Description: null.StringFrom("Assistance for setting up home office"),
			Price:       600.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1486312338219-ce68d2c6f44d?w=400&h=300&fit=crop"),
		},
		{
			Name:        "Extended Paternity Leave",
			Description: null.StringFrom("20-day paternity leave (in addition to the 5 legal days)"),
			Price:       0.00,
			Image:       null.StringFrom("https://images.unsplash.com/photo-1544367567-0f2fcb009e0b?w=400&h=300&fit=crop"),
		},
	}

	for i, benefit := range benefits {
		err := benefit.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("Error inserting benefit %d: %v", i+1, err)
			continue
		}
		fmt.Printf("✓ Benefit inserted: %s ($%.2f)\n",
			benefit.Name,
			benefit.Price)
	}
}

func seedFeedbacks(ctx context.Context, db boil.ContextExecutor, employees []*models.Employee) {
	if len(employees) < 2 {
		fmt.Println("⚠️  Need at least 2 employees to create feedbacks")
		return
	}

	feedbackTexts := []string{
		"Excellent work on implementing the new system. Demonstrated great technical knowledge.",
		"Very collaborative and always willing to help colleagues. Great communication.",
		"Delivered the project on time and with exceptional quality. Congratulations!",
		"Proactive in problem solving and always seeks the best solutions.",
		"Exemplary leadership during the project. Managed to motivate the entire team.",
		"Impressive attention to detail. Code always well documented.",
		"Great learning ability and adaptation to new technologies.",
		"Constructive feedback and always seeks to improve team processes.",
		"Clear and objective presentations. Makes understanding much easier.",
		"Exceptional commitment to work quality.",
		"Very organized and efficient in task management.",
		"Creativity in solving complex problems.",
		"Excellent teamwork and integration with other departments.",
		"Solid technical mastery and always shares knowledge.",
		"Exemplary punctuality and responsibility.",
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
			log.Printf("Error inserting feedback %d: %v", i+1, err)
			continue
		}

		fmt.Printf("✓ Feedback inserted: %s → %s (⭐ %d)\n",
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
	benCount, _ := models.Benefits().Count(ctx, db)

	fmt.Printf("\n📊 FINAL SUMMARY:\n")
	fmt.Printf("==================\n")
	fmt.Printf("👥 Employees: %d\n", empCount)
	fmt.Printf("📦 Resources: %d\n", resCount)
	fmt.Printf("📅 Activities: %d\n", actCount)
	fmt.Printf("💬 Feedbacks: %d\n", feedCount)
	fmt.Printf("🎁 Benefits: %d\n", benCount)
	fmt.Printf("==================\n")
}
