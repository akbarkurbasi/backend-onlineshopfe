package impl

import (
	"context"
	"fmt"
	"math"

	"github.com/RakaMurdiarta/online-shop-system/internal/models"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/delivery"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/repository"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/services"
	"github.com/RakaMurdiarta/online-shop-system/pkg/common/response"
)

type feedbackServiceImpl struct {
	repo repository.FeedbackRepository
}

func NewFeedbackService(repo repository.FeedbackRepository) services.FeedbackService {
	return &feedbackServiceImpl{repo: repo}
}

func (s *feedbackServiceImpl) Create(ctx context.Context, req *delivery.CreateFeedbackRequest) (*delivery.FeedbackResponse, error) {
	feedback := &models.Feedback{
		Name:    req.Name,
		Email:   req.Email,
		Subject: req.Subject,
		Message: req.Message,
	}

	if err := s.repo.WithTransaction(ctx, func(txCtx context.Context) error {
		return s.repo.Create(txCtx, feedback)
	}); err != nil {
		return nil, fmt.Errorf("failed to submit feedback: %w", err)
	}

	return delivery.ToFeedbackResponse(feedback), nil
}

func (s *feedbackServiceImpl) GetAll(ctx context.Context, page, limit int, search string) (*delivery.FeedbackListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	feedbacks, total, err := s.repo.FindAll(ctx, limit, offset, search)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feedbacks: %w", err)
	}

	items := make([]delivery.FeedbackResponse, 0, len(feedbacks))
	for i := range feedbacks {
		items = append(items, *delivery.ToFeedbackResponse(&feedbacks[i]))
	}

	totalPage := 0
	if total > 0 {
		totalPage = int(math.Ceil(float64(total) / float64(limit)))
	}

	return &delivery.FeedbackListResponse{
		Items: items,
		Paging: response.PagingResponse{
			CurrentPage: page,
			TotalPage:   totalPage,
			TotalItems:  total,
		},
	}, nil
}

func (s *feedbackServiceImpl) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete feedback: %w", err)
	}
	return nil
}
