package service

import (
	"context"
	"hw4/internal/model"
	"hw4/internal/storage"
)

type Service struct {
	repo storage.Repository
}

func NewService(repo storage.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetStudentInfo(ctx context.Context, id int) (*model.StudentResponse, error) {
	return s.repo.GetStudent(ctx, id)
}

func (s *Service) GetAllSchedules(ctx context.Context, limit int, offset int) ([]model.ClassSchedule, error) {
	return s.repo.GetAllSchedules(ctx, limit, offset)
}

func (s *Service) GetGroupSchedule(ctx context.Context, id string) ([]model.ClassSchedule, error) {
	return s.repo.GetGroupSchedule(ctx, id)
}

// TODO: Implement the following service methods
func (s *Service) RecordVisit(ctx context.Context, vr model.VisitRecord) error {
	// Implementation goes here

	return s.repo.RecordVisit(ctx, vr)
}

func (s *Service) GetAttendanceByClass(ctx context.Context, classID int) ([]model.VisitRecord, error) {
	// Implementation goes here
	return s.repo.GetAttendanceByClass(ctx, classID)
}

func (s *Service) GetAttendanceByStudent(ctx context.Context, studentID int) ([]model.VisitRecord, error) {
	// Implementation goes here
	return s.repo.GetAttendanceByStudent(ctx, studentID)
}
