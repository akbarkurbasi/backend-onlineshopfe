package provider

import (
	"github.com/RakaMurdiarta/online-shop-system/internal/middlewares"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/handlers"
	feedbackRepoImpl "github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/repository/Impl"
	feedbackServiceImpl "github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/services/Impl"
	"github.com/RakaMurdiarta/online-shop-system/pkg/database"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func FeedbackProvider(
	tm *database.TransactionManagerImpl,
	privateRoute *echo.Group,
	publicRoute *echo.Group,
) {
	v := validator.New()

	repo := feedbackRepoImpl.NewFeedbackRepository(tm)
	service := feedbackServiceImpl.NewFeedbackService(repo)
	handler := handlers.NewFeedbackHandler(service, v)

	// User umum bisa kirim feedback
	publicGroup := publicRoute.Group("/feedback")
	publicGroup.POST("", handler.Create)

	// Hanya admin yang bisa kelola feedback
	privateGroup := privateRoute.Group("/feedback")
	privateGroup.GET("", handler.GetAll, middlewares.IsAdmin)
	privateGroup.DELETE("/:id", handler.Delete, middlewares.IsAdmin)
}
