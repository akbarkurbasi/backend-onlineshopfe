package delivery

import (
	"time"

	"github.com/RakaMurdiarta/online-shop-system/internal/models"
	"github.com/RakaMurdiarta/online-shop-system/pkg/common/response"
)

type FeedbackResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FeedbackListResponse struct {
	Items  []FeedbackResponse      `json:"items"`
	Paging response.PagingResponse `json:"paging"`
}

func ToFeedbackResponse(f *models.Feedback) *FeedbackResponse {
	if f == nil {
		return nil
	}

	return &FeedbackResponse{
		ID:        f.ID,
		Name:      f.Name,
		Email:     f.Email,
		Subject:   f.Subject,
		Message:   f.Message,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}
