package impl

import (
	"context"

	"github.com/RakaMurdiarta/online-shop-system/internal/models"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/repository"
	"github.com/RakaMurdiarta/online-shop-system/pkg/database"
)

type feedbackRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewFeedbackRepository(tm *database.TransactionManagerImpl) repository.FeedbackRepository {
	return &feedbackRepositoryImpl{TransactionManagerImpl: tm}
}

func (r *feedbackRepositoryImpl) Create(ctx context.Context, f *models.Feedback) error {
	return r.GetTx(ctx).Create(f).Error
}

func (r *feedbackRepositoryImpl) FindAll(ctx context.Context, limit, offset int, search string) ([]models.Feedback, int64, error) {
	var feedbacks []models.Feedback
	var total int64

	query := r.GetTx(ctx).Model(&models.Feedback{})

	if search != "" {
		// Mencari berdasarkan nama, email, atau subject
		query = query.Where("name ILIKE ? OR email ILIKE ? OR subject ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&feedbacks).Error

	return feedbacks, total, err
}

func (r *feedbackRepositoryImpl) Delete(ctx context.Context, id int) error {
	// Karena ada deleted_at, GORM otomatis melakukan Soft Delete
	return r.GetTx(ctx).Where("id = ?", id).Delete(&models.Feedback{}).Error
}
