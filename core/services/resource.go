package services

import (
	"context"
	"database/sql"
	"hacka/core/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ResourceService struct {
	*BaseService
}

func NewResourceService(db *sql.DB) *ResourceService {
	return &ResourceService{
		BaseService: NewBaseService(db),
	}
}

func (s *ResourceService) GetAllResources(ctx context.Context) (models.ResourceSlice, error) {
	return models.Resources(
		qm.OrderBy("created_at DESC"),
	).All(ctx, s.DB)
}

func (s *ResourceService) GetResourceByID(ctx context.Context, id string) (*models.Resource, error) {
	return models.FindResource(ctx, s.DB, id)
}

func (s *ResourceService) CreateResource(ctx context.Context, name, resourceType, midia string) (*models.Resource, error) {
	resource := &models.Resource{
		Name:  null.StringFrom(name),
		Type:  null.StringFrom(resourceType),
		Midia: null.StringFrom(midia),
	}

	err := resource.Insert(ctx, s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return resource, nil
}

// UpdateResource atualiza um recurso existente
func (s *ResourceService) UpdateResource(ctx context.Context, id, name, resourceType, midia string) (*models.Resource, error) {
	resource, err := s.GetResourceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		resource.Name = null.StringFrom(name)
	}
	if resourceType != "" {
		resource.Type = null.StringFrom(resourceType)
	}
	if midia != "" {
		resource.Midia = null.StringFrom(midia)
	}

	_, err = resource.Update(ctx, s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (s *ResourceService) DeleteResource(ctx context.Context, id string) error {
	resource, err := s.GetResourceByID(ctx, id)
	if err != nil {
		return err
	}

	_, err = resource.Delete(ctx, s.DB)
	return err
}

func (s *ResourceService) GetResourcesByType(ctx context.Context, resourceType string) (models.ResourceSlice, error) {
	return models.Resources(
		qm.Where("type = ?", resourceType),
		qm.OrderBy("created_at DESC"),
	).All(ctx, s.DB)
}

func (s *ResourceService) GetResourcesByAverageRanking(ctx context.Context, limit int) (models.ResourceSlice, error) {
	query := models.Resources(
		qm.Where("average IS NOT NULL"),
		qm.OrderBy("average DESC"),
	)

	if limit > 0 {
		query = models.Resources(
			qm.Where("average IS NOT NULL"),
			qm.OrderBy("average DESC"),
			qm.Limit(limit),
		)
	}

	return query.All(ctx, s.DB)
}
