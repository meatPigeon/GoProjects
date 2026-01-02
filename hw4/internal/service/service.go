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

func (s *Service) GetStudentInfo(ctx context.Context, id string) (*model.StudentResponse, error) {
	return s.repo.GetStudent(ctx, id)
}

func (s *Service) GetAllSchedules(ctx context.Context) ([]model.ClassSchedule, error) {
	return s.repo.GetAllSchedules(ctx)
}

func (s *Service) GetGroupSchedule(ctx context.Context, id string) ([]model.ClassSchedule, error) {
	return s.repo.GetGroupSchedule(ctx, id)
}
