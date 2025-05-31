package services

import (
	"context"
	"database/sql"
	"hacka/core/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type EmployeeService struct {
	*BaseService
}

func NewEmployeeService(db *sql.DB) *EmployeeService {
	return &EmployeeService{
		BaseService: NewBaseService(db),
	}
}

func (s *EmployeeService) GetAllEmployees(ctx context.Context) (models.EmployeeSlice, error) {
	return models.Employees().All(ctx, s.DB)
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id string) (*models.Employee, error) {
	return models.FindEmployee(ctx, s.DB, id)
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, name, email, position string, balance, average float64, qrcode, passwordHash, midia string) (*models.Employee, error) {
	employee := &models.Employee{
		Name:         null.StringFrom(name),
		Email:        null.StringFrom(email),
		Position:     null.StringFrom(position),
		Balance:      null.Float64From(balance),
		Average:      null.Float64From(average),
		Qrcode:       null.StringFrom(qrcode),
		PasswordHash: null.StringFrom(passwordHash),
		Midia:        null.StringFrom(midia),
	}

	err := employee.Insert(ctx, s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// UpdateEmployee atualiza um funcionário existente
func (s *EmployeeService) UpdateEmployee(ctx context.Context, id string, updates map[string]interface{}) (*models.Employee, error) {
	employee, err := s.GetEmployeeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Aplicar updates
	if name, ok := updates["name"].(string); ok {
		employee.Name = null.StringFrom(name)
	}
	if email, ok := updates["email"].(string); ok {
		employee.Email = null.StringFrom(email)
	}
	if position, ok := updates["position"].(string); ok {
		employee.Position = null.StringFrom(position)
	}
	if balance, ok := updates["balance"].(float64); ok {
		employee.Balance = null.Float64From(balance)
	}
	if average, ok := updates["average"].(float64); ok {
		employee.Average = null.Float64From(average)
	}

	_, err = employee.Update(ctx, s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id string) error {
	employee, err := s.GetEmployeeByID(ctx, id)
	if err != nil {
		return err
	}

	_, err = employee.Delete(ctx, s.DB)
	return err
}

func (s *EmployeeService) GetEmployeesByPosition(ctx context.Context, position string) (models.EmployeeSlice, error) {
	return models.Employees(
		qm.Where("position = ?", position),
	).All(ctx, s.DB)
}

func (s *EmployeeService) GetHighBalanceEmployees(ctx context.Context, minBalance float64) (models.EmployeeSlice, error) {
	return models.Employees(
		qm.Where("balance > ?", minBalance),
		qm.OrderBy("balance DESC"),
	).All(ctx, s.DB)
}

func (s *EmployeeService) GetEmployeesByAverageRanking(ctx context.Context, limit int) (models.EmployeeSlice, error) {
	query := models.Employees(
		qm.Where("average IS NOT NULL"),
		qm.OrderBy("average DESC"),
	)

	if limit > 0 {
		query = models.Employees(
			qm.Where("average IS NOT NULL"),
			qm.OrderBy("average DESC"),
			qm.Limit(limit),
		)
	}

	return query.All(ctx, s.DB)
}
