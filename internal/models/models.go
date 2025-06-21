package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Пользователь системы
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" validate:"required,min=3,max=50"`
	Password  string    `json:"password,omitempty" validate:"required,min=6" db:"password"`
	Email     string    `json:"email" validate:"required,email" db:"email"`
	Role      string    `json:"role" db:"role"` // admin, teacher, student
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Курс
type Course struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" db:"description"`
}

// Урок
type Lesson struct {
	ID        int       `json:"id" db:"id"`
	CourseID  int       `json:"course_id" db:"course_id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" db:"content"`
	Order     int       `json:"order" db:"order"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Валидация моделей
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (c *Course) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
