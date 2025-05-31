package services

import (
	"context"
	"database/sql"
	"hacka/core/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ActivityService struct {
	*BaseService
}

func NewActivityService(db *sql.DB) *ActivityService {
	return &ActivityService{
		BaseService: NewBaseService(db),
	}
}

func (s *ActivityService) GetAllActivities(ctx context.Context) (models.ActivitySlice, error) {
	return models.Activities(
		qm.OrderBy("created_at DESC"),
	).All(ctx, s.DB)
}

func (s *ActivityService) GetActivityByID(ctx context.Context, id string) (*models.Activity, error) {
	return models.FindActivity(ctx, s.DB, id)
}

func (s *ActivityService) CreateActivity(ctx context.Context, name, description, activityType string) (*models.Activity, error) {
	activity := &models.Activity{
		Name:        null.StringFrom(name),
		Description: null.StringFrom(description),
		Type:        null.StringFrom(activityType),
	}

	err := activity.Insert(ctx, s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (s *ActivityService) UpdateActivity(ctx context.Context, id, name, description, activityType string) (*models.Activity, error) {
	activity, err := s.GetActivityByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		activity.Name = null.StringFrom(name)
	}
	if description != "" {
		activity.Description = null.StringFrom(description)
	}
	if activityType != "" {
		activity.Type = null.StringFrom(activityType)
	}

	_, err = activity.Update(ctx, s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (s *ActivityService) DeleteActivity(ctx context.Context, id string) error {
	activity, err := s.GetActivityByID(ctx, id)
	if err != nil {
		return err
	}

	_, err = activity.Delete(ctx, s.DB)
	return err
}

func (s *ActivityService) GetActivitiesByType(ctx context.Context, activityType string) (models.ActivitySlice, error) {
	return models.Activities(
		qm.Where("type = ?", activityType),
		qm.OrderBy("created_at DESC"),
	).All(ctx, s.DB)
}

func (s *ActivityService) GetActivitiesByAverageRanking(ctx context.Context, limit int) (models.ActivitySlice, error) {
	query := models.Activities(
		qm.Where("average IS NOT NULL"),
		qm.OrderBy("average DESC"),
	)

	if limit > 0 {
		query = models.Activities(
			qm.Where("average IS NOT NULL"),
			qm.OrderBy("average DESC"),
			qm.Limit(limit),
		)
	}

	return query.All(ctx, s.DB)
}
