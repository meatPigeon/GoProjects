package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

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

var conn *pgx.Conn

func main() {
	url := "postgres:///university"
	var err error
	conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	e := echo.New()

	e.GET("/student/:id", getStudent)
	e.GET("/all_class_schedule", getAllSchedules)
	e.GET("/schedule/group/:id", getGroupSchedule)

	e.Logger.Fatal(e.Start(":8080"))
}

func getStudent(c echo.Context) error {
	id := c.Param("id")

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

	var s StudentResponse
	err := conn.QueryRow(context.Background(), sql, id).Scan(
		&s.StudentID,
		&s.StudentName,
		&s.BirthDate,
		&s.Gender,
		&s.GroupName,
	)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
	}

	return c.JSON(http.StatusOK, s)
}

func getAllSchedules(c echo.Context) error {
	sql := `
		SELECT class_id, group_id, class_name, class_date, start_time, end_time, room, teacher_id 
		FROM class_schedule
	`

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	schedules := []ClassSchedule{}

	for rows.Next() {
		var cs ClassSchedule

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
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		schedules = append(schedules, cs)
	}

	if rows.Err() != nil {
		return c.JSON(http.StatusInternalServerError, rows.Err().Error())
	}

	return c.JSON(http.StatusOK, schedules)
}

func getGroupSchedule(c echo.Context) error {
	id := c.Param("id")

	sql := `
		SELECT class_id, group_id, class_name, class_date, start_time, end_time, room, teacher_id
		FROM class_schedule
		WHERE group_id = $1;
	`

	rows, err := conn.Query(context.Background(), sql, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	schedules := []ClassSchedule{}

	for rows.Next() {
		var cs ClassSchedule
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
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		schedules = append(schedules, cs)
	}

	return c.JSON(http.StatusOK, schedules)
}
