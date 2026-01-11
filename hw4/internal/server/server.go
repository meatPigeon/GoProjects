package server

import (
	"hw4/internal/handler"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) MapRoutes(e *echo.Echo, h *handler.Handler) {
	e.GET("/student/:id", h.GetStudent)
	e.GET("/all_class_schedule", h.GetAllSchedules)
	e.GET("/schedule/group/:id", h.GetGroupSchedule)

	// hw 5 routes
	e.POST("/attendance/class", h.RecordVisit)
	e.GET("/attendance/class/:id", h.GetAttendanceByClass)
	e.GET("/attendance/student/:id", h.GetAttendanceByStudent)
}
