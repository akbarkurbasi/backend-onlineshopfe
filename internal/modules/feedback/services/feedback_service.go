package services

import (
	"context"

	"github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/delivery"
)

type FeedbackService interface {
	Create(ctx context.Context, req *delivery.CreateFeedbackRequest) (*delivery.FeedbackResponse, error)
	GetAll(ctx context.Context, page, limit int, search string) (*delivery.FeedbackListResponse, error)
	Delete(ctx context.Context, id int) error
}
