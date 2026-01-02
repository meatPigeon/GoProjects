package postgres

import (
	"context"
	"hw4/internal/model"
	"hw4/internal/storage"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) storage.Repository {
	return &Repository{conn: conn}
}

func (r *Repository) GetStudent(ctx context.Context, id string) (*model.StudentResponse, error) {
	sql := `
		SELECT 
			s.student_id, 
			s.student_name, 
			s.birth_date, 
			s.gender, 
			g.group_name 
		FROM students s 
		JOIN groups g ON s.group_id = g.group_id 
		WHERE s.student_id = $1
	`
	var s model.StudentResponse
	err := r.conn.QueryRow(ctx, sql, id).Scan(
		&s.StudentID,
		&s.StudentName,
		&s.BirthDate,
		&s.Gender,
		&s.GroupName,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *Repository) GetAllSchedules(ctx context.Context) ([]model.ClassSchedule, error) {
	sql := `
		SELECT class_id, group_id, class_name, class_date, start_time, end_time, room, teacher_id 
		FROM class_schedule
	`
	rows, err := r.conn.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []model.ClassSchedule
	for rows.Next() {
		var cs model.ClassSchedule
		err := rows.Scan(
			&cs.ClassID,
			&cs.GroupID,
			&cs.ClassName,
			&cs.ClassDate,
			&cs.StartTime,
			&cs.EndTime,
			&cs.Room,
			&cs.TeacherID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, cs)
	}
	return schedules, nil
}

func (r *Repository) GetGroupSchedule(ctx context.Context, groupID string) ([]model.ClassSchedule, error) {
	sql := `
		SELECT class_id, group_id, class_name, class_date, start_time, end_time, room, teacher_id
		FROM class_schedule
		WHERE group_id = $1;
	`
	rows, err := r.conn.Query(ctx, sql, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []model.ClassSchedule
	for rows.Next() {
		var cs model.ClassSchedule
		err := rows.Scan(
			&cs.ClassID,
			&cs.GroupID,
			&cs.ClassName,
			&cs.ClassDate,
			&cs.StartTime,
			&cs.EndTime,
			&cs.Room,
			&cs.TeacherID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, cs)
	}
	return schedules, nil
}
