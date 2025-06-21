package services

import (
	"database/sql"

	"github.com/Muhandis2000/online-school/internal/models"
)

type CourseService struct {
	db *sql.DB
}

func NewCourseService(db *sql.DB) *CourseService {
	return &CourseService{db: db}
}

func (s *CourseService) GetCourses() ([]models.Course, error) {
	rows, err := s.db.Query("SELECT id, title, description FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []models.Course
	for rows.Next() {
		var c models.Course
		if err := rows.Scan(&c.ID, &c.Title, &c.Description); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func (s *CourseService) CreateCourse(course models.Course) (int, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO courses (title, description) VALUES ($1, $2) RETURNING id",
		course.Title, course.Description).Scan(&id)
	return id, err
}
