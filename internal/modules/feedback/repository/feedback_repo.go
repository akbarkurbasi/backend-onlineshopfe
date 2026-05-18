package repository

import (
	"context"

	"github.com/RakaMurdiarta/online-shop-system/internal/models"
	"github.com/RakaMurdiarta/online-shop-system/pkg/database"
)

type FeedbackRepository interface {
	Create(ctx context.Context, f *models.Feedback) error
	FindAll(ctx context.Context, limit, offset int, search string) ([]models.Feedback, int64, error)
	Delete(ctx context.Context, id int) error

	// for database transaction
	database.TransactionManager
}
