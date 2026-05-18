package handlers

import (
	"net/http"
	"strconv"

	"github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/delivery"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/feedback/services"
	"github.com/RakaMurdiarta/online-shop-system/pkg/common/customs"
	"github.com/RakaMurdiarta/online-shop-system/pkg/common/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type FeedbackHandler struct {
	fs services.FeedbackService
	v  *validator.Validate
}

func NewFeedbackHandler(fs services.FeedbackService, v *validator.Validate) *FeedbackHandler {
	return &FeedbackHandler{fs: fs, v: v}
}

// Create bisa diakses publik (User kirim feedback)
func (h *FeedbackHandler) Create(c *echo.Context) error {
	req := new(delivery.CreateFeedbackRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			"Invalid Request",
			customs.HandleBindError(err)...,
		))
	}

	if err := h.v.StructCtx(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			"Validation Failed",
			*customs.NewErrorValue("validation", err.Error()),
		))
	}

	res, err := h.fs.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			err.Error(),
			*customs.NewErrorValue("business_logic", err.Error()),
		))
	}

	return c.JSON(http.StatusCreated, response.NewResponseSuccess(res, "Feedback Sent Successfully"))
}

// GetAll hanya untuk Admin
func (h *FeedbackHandler) GetAll(c *echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	search := c.QueryParam("search")

	res, err := h.fs.GetAll(c.Request().Context(), page, limit, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewResponseError(
			err.Error(),
			*customs.NewErrorValue("business_logic", err.Error()),
		))
	}

	return c.JSON(http.StatusOK, response.NewResponseSuccess(res, "Feedbacks Retrieved"))
}

// Delete hanya untuk Admin
func (h *FeedbackHandler) Delete(c *echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if id == 0 {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			"Invalid Feedback ID",
			*customs.NewErrorValue("validation", "id is required and must be numeric"),
		))
	}

	if err := h.fs.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			err.Error(),
			*customs.NewErrorValue("business_logic", err.Error()),
		))
	}

	return c.JSON(http.StatusOK, response.NewResponseSuccess[*bool](nil, "Feedback Deleted"))
}
