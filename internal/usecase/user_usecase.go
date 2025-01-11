package usecase

import (
	"github.com/TomasSorgetti/go-clean/internal/entity"
	"github.com/TomasSorgetti/go-clean/internal/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(user *entity.User) error {
	return u.repo.Create(user)
}

func (u *UserUseCase) GetUserByID(id int64) (*entity.User, error) {
	return u.repo.FindByID(id)
}
