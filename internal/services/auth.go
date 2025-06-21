package services

import (
	"database/sql"
	"fmt"

	"github.com/Muhandis2000/online-school/internal/models"
	"github.com/Muhandis2000/online-school/internal/utils"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(user models.User) (uint, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	var id uint
	err = s.db.QueryRow("INSERT INTO users (username, password, email, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Username, hashedPassword, user.Email, user.Role).Scan(&id)
	return id, err
}

func (s *AuthService) Login(username, password string) (string, error) {
	var user models.User
	row := s.db.QueryRow("SELECT id, password, role FROM users WHERE username = $1", username)
	if err := row.Scan(&user.ID, &user.Password, &user.Role); err != nil {
		return "", err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", fmt.Errorf("invalid password")
	}
	return utils.GenerateJWT(uint(user.ID), user.Role)
}
