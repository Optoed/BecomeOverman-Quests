package services

import (
	"BecomeOverMan/internal/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(username, email, password string) error {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedPassword := string(hashedPasswordBytes)
	return s.repo.CreateUser(username, email, hashedPassword)
}

func (s *UserService) Login(username, password string) (int, error) {
	// Логируем начало попытки входа
	log.Printf("User %s is attempting to log in", username)

	// Получаем пользователя по username
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		// Логируем ошибку при получении пользователя
		log.Printf("Error getting user by username %s: %v", username, err)
		return 0, errors.New("invalid username or password")
	}

	// Сравниваем пароли
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		// Логируем неудачную попытку входа
		log.Printf("Failed login attempt for user %s: invalid password", username)
		return 0, errors.New("invalid username or password")
	}

	// Логируем успешный вход
	log.Printf("User %s logged in successfully", username)

	return user.ID, nil
}
