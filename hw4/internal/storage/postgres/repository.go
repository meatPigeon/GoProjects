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

func (r *Repository) GetStudent(ctx context.Context, id int) (*model.StudentResponse, error) {
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

func (r *Repository) GetAllSchedules(ctx context.Context, limit int, offset int) ([]model.ClassSchedule, error) {
	sql := `
		SELECT class_id, group_id, class_name, class_date, start_time, end_time, room, teacher_id 
		FROM class_schedule
		LIMIT $1 OFFSET $2
	`
	rows, err := r.conn.Query(ctx, sql, limit, offset)
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

// Implement the following repository methods
func (r *Repository) RecordVisit(ctx context.Context, vr model.VisitRecord) error {
	// Implementation goes here
	sql := `
		INSERT INTO attendance (student_id, class_id, visit_date, present)
		VALUES ($1, $2, $3, $4);
	`
	_, err := r.conn.Exec(ctx, sql,
		vr.StudentID,
		vr.ClassID,
		vr.VisitDate,
		vr.Present,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAttendanceByClass(ctx context.Context, classID int) ([]model.VisitRecord, error) {
	// Implementation goes here
	sql := `
		SELECT student_id, class_id, visit_date, present
		FROM attendance
		WHERE class_id = $1;
	`
	rows, err := r.conn.Query(ctx, sql, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []model.VisitRecord
	for rows.Next() {
		var vr model.VisitRecord
		err := rows.Scan(
			&vr.StudentID,
			&vr.ClassID,
			&vr.VisitDate,
			&vr.Present,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, vr)
	}
	return records, nil
}

func (r *Repository) GetAttendanceByStudent(ctx context.Context, studentID int) ([]model.VisitRecord, error) {
	// Implementation goes here
	sql := `
		SELECT student_id, class_id, visit_date, present
		FROM attendance
		WHERE student_id = $1;
	`
	rows, err := r.conn.Query(ctx, sql, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []model.VisitRecord
	for rows.Next() {
		var vr model.VisitRecord
		err := rows.Scan(
			&vr.StudentID,
			&vr.ClassID,
			&vr.VisitDate,
			&vr.Present,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, vr)
	}
	return records, nil
}
