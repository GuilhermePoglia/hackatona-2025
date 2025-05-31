package services

import (
	"context"
	"database/sql"
	"fmt"

	"hacka/core/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type BenefitService struct {
	base BaseService
}

func NewBenefitService(db *sql.DB) *BenefitService {
	return &BenefitService{
		base: BaseService{DB: db},
	}
}

func (s *BenefitService) GetAll(ctx context.Context) ([]*models.Benefit, error) {
	benefits, err := models.Benefits(
		qm.OrderBy("name ASC"),
	).All(ctx, s.base.DB)

	if err != nil {
		return nil, fmt.Errorf("failed to get benefits: %w", err)
	}

	return benefits, nil
}

func (s *BenefitService) GetByID(ctx context.Context, id string) (*models.Benefit, error) {
	benefit, err := models.FindBenefit(ctx, s.base.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("benefit not found with id: %s", id)
		}
		return nil, fmt.Errorf("failed to get benefit: %w", err)
	}

	return benefit, nil
}

func (s *BenefitService) Create(ctx context.Context, benefit *models.Benefit) error {
	err := benefit.Insert(ctx, s.base.DB, boil.Infer())
	if err != nil {
		return fmt.Errorf("failed to create benefit: %w", err)
	}

	return nil
}

func (s *BenefitService) Update(ctx context.Context, id string, benefit *models.Benefit) error {
	existingBenefit, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	benefit.ID = existingBenefit.ID
	_, err = benefit.Update(ctx, s.base.DB, boil.Infer())
	if err != nil {
		return fmt.Errorf("failed to update benefit: %w", err)
	}

	return nil
}

func (s *BenefitService) Delete(ctx context.Context, id string) error {
	benefit, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_, err = benefit.Delete(ctx, s.base.DB)
	if err != nil {
		return fmt.Errorf("failed to delete benefit: %w", err)
	}

	return nil
}

func (s *BenefitService) GetByPriceRange(ctx context.Context, minPrice, maxPrice float64) ([]*models.Benefit, error) {
	benefits, err := models.Benefits(
		qm.Where("price >= ? AND price <= ?", minPrice, maxPrice),
		qm.OrderBy("price ASC"),
	).All(ctx, s.base.DB)

	if err != nil {
		return nil, fmt.Errorf("failed to get benefits by price range: %w", err)
	}

	return benefits, nil
}
