package services

import (
	"context"
	"database/sql"
	"hacka/core/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type FeedbackService struct {
	*BaseService
	employeeService *EmployeeService
}

func NewFeedbackService(db *sql.DB, employeeService *EmployeeService) *FeedbackService {
	return &FeedbackService{
		BaseService:     NewBaseService(db),
		employeeService: employeeService,
	}
}

func (s *FeedbackService) CreateFeedback(ctx context.Context, senderID, receiverID string, stars int, description string) (*models.Feedback, error) {
	_, err := s.employeeService.GetEmployeeByID(ctx, senderID)
	if err != nil {
		return nil, err
	}

	_, err = s.employeeService.GetEmployeeByID(ctx, receiverID)
	if err != nil {
		return nil, err
	}

	feedback := &models.Feedback{
		SenderID:    senderID,
		ReceiverID:  receiverID,
		Stars:       stars,
		Description: null.StringFrom(description),
	}

	err = feedback.Insert(ctx, s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	err = s.updateReceiverStats(ctx, receiverID)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

func (s *FeedbackService) GetFeedbacksByReceiver(ctx context.Context, receiverID string) (models.FeedbackSlice, error) {
	return models.Feedbacks(
		qm.Where("receiver_id = ?", receiverID),
		qm.OrderBy("created_at DESC"),
	).All(ctx, s.DB)
}

func (s *FeedbackService) GetFeedbacksBySender(ctx context.Context, senderID string) (models.FeedbackSlice, error) {
	return models.Feedbacks(
		qm.Where("sender_id = ?", senderID),
		qm.OrderBy("created_at DESC"),
	).All(ctx, s.DB)
}

func (s *FeedbackService) GetAllFeedbacks(ctx context.Context) (models.FeedbackSlice, error) {
	return models.Feedbacks(
		qm.OrderBy("created_at DESC"),
	).All(ctx, s.DB)
}

func (s *FeedbackService) GetFeedbackByID(ctx context.Context, id string) (*models.Feedback, error) {
	return models.FindFeedback(ctx, s.DB, id)
}

func (s *FeedbackService) updateReceiverStats(ctx context.Context, receiverID string) error {
	feedbacks, err := s.GetFeedbacksByReceiver(ctx, receiverID)
	if err != nil {
		return err
	}

	if len(feedbacks) == 0 {
		return nil
	}

	totalStars := 0
	for _, feedback := range feedbacks {
		totalStars += feedback.Stars
	}
	newAverage := float64(totalStars) / float64(len(feedbacks))

	newBalance := newAverage * 100

	receiver, err := s.employeeService.GetEmployeeByID(ctx, receiverID)
	if err != nil {
		return err
	}

	receiver.Average = null.Float64From(newAverage)
	receiver.Balance = null.Float64From(newBalance)

	_, err = receiver.Update(ctx, s.DB, boil.Infer())
	return err
}

func (s *FeedbackService) GetEmployeeStats(ctx context.Context, employeeID string) (map[string]interface{}, error) {
	feedbacks, err := s.GetFeedbacksByReceiver(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total_feedbacks": len(feedbacks),
		"stars_breakdown": map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
	}

	if len(feedbacks) > 0 {
		totalStars := 0
		starsBreakdown := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}

		for _, feedback := range feedbacks {
			totalStars += feedback.Stars
			starsBreakdown[feedback.Stars]++
		}

		stats["average_stars"] = float64(totalStars) / float64(len(feedbacks))
		stats["stars_breakdown"] = starsBreakdown
	} else {
		stats["average_stars"] = 0.0
	}

	return stats, nil
}
