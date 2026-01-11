package storage

import (
	"context"
	"hw4/internal/model"
)

type Repository interface {
	GetStudent(ctx context.Context, id int) (*model.StudentResponse, error)
	GetAllSchedules(ctx context.Context, limit int, offset int) ([]model.ClassSchedule, error)
	GetGroupSchedule(ctx context.Context, groupID string) ([]model.ClassSchedule, error)

	// hw5
	RecordVisit(ctx context.Context, vr model.VisitRecord) error
	GetAttendanceByClass(ctx context.Context, classID int) ([]model.VisitRecord, error)
	GetAttendanceByStudent(ctx context.Context, studentID int) ([]model.VisitRecord, error)
}
