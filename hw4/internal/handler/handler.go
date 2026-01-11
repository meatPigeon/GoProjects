package handler

import (
	"hw4/internal/model"
	"hw4/internal/service"
	"net/http"
	"strconv"

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

	studentID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	student, err := h.service.GetStudentInfo(c.Request().Context(), studentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) GetAllSchedules(c echo.Context) error {

	limit := 10
	offset := 0

	if c.QueryParam("limit") != "" {
		limit, _ = strconv.Atoi(c.QueryParam("limit"))
	}
	if c.QueryParam("offset") != "" {
		offset, _ = strconv.Atoi(c.QueryParam("offset"))
	}

	schedules, err := h.service.GetAllSchedules(c.Request().Context(), limit, offset)
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

// TODO: Implement the following handler functions

func (h *Handler) RecordVisit(c echo.Context) error {
	// Implementation goes here
	var vr model.VisitRecord

	if err := c.Bind(&vr); err != nil {
		return c.JSON(400, "invalid json")
	}

	// if err := c.Validate(&vr); err != nil {
	// 	return c.JSON(400, err.Error())
	// }

	// Call service method to record visit (to be implemented)
	if err := h.service.RecordVisit(c.Request().Context(), vr); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(200, "visit recorded")
}

func (h *Handler) GetAttendanceByClass(c echo.Context) error {
	// Implementation goes here
	id := c.Param("id")

	classID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	attendance, err := h.service.GetAttendanceByClass(c.Request().Context(), classID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, attendance)
}

func (h *Handler) GetAttendanceByStudent(c echo.Context) error {
	// Implementation goes here
	id := c.Param("id")

	studentID, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	attendance, err := h.service.GetAttendanceByStudent(c.Request().Context(), studentID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, attendance)
}
