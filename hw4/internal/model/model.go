package model

import "time"

type StudentResponse struct {
	StudentID   int       `json:"student_id"`
	StudentName string    `json:"student_name"`
	BirthDate   time.Time `json:"birth_date"`
	Gender      *string   `json:"gender"`
	GroupName   string    `json:"group_name"`
}

type ClassSchedule struct {
	ClassID   int       `json:"class_id"`
	GroupID   *int      `json:"group_id"`
	ClassName string    `json:"class_name"`
	ClassDate time.Time `json:"class_date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Room      *string   `json:"room"`
	TeacherID *int      `json:"teacher_id"`
}

type Group struct {
	GroupID   int    `json:"group_id"`
	GroupName string `json:"group_name"`
}

type VisitRecord struct {
	StudentID int       `json:"student_id"`
	ClassID   int       `json:"class_id"`
	VisitDate time.Time `json:"visit_date"`
	Present   bool      `json:"present"`
}
