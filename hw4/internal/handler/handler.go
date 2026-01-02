package handler

import (
	"hw4/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetStudent(c echo.Context) error {
	id := c.Param("id")
	student, err := h.service.GetStudentInfo(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) GetAllSchedules(c echo.Context) error {
	schedules, err := h.service.GetAllSchedules(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, schedules)
}

func (h *Handler) GetGroupSchedule(c echo.Context) error {
	id := c.Param("id")
	schedules, err := h.service.GetGroupSchedule(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, schedules)
}
