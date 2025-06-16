package service

import (
	"errors"

	"github.com/brunnossanttos/libri/internal/models"
	"github.com/brunnossanttos/libri/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{Repo: r}
}

func (s *UserService) CreateUser(name, email, password string) (*models.User, error) {
	if existing, _ := s.Repo.GetByEmail(email); existing != nil {
		return nil, errors.New("email já cadastrado")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}
	if err := s.Repo.Create(u); err != nil {
		return nil, err
	}
	u.Password = ""
	return u, nil
}
