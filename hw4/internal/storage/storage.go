package storage

import (
	"context"
	"hw4/internal/model"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*model.StudentResponse, error)
	GetAllSchedules(ctx context.Context) ([]model.ClassSchedule, error)
	GetGroupSchedule(ctx context.Context, groupID string) ([]model.ClassSchedule, error)
}
