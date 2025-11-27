package service

import (
	"github.com/MiladCodeLab/simple-api/entity"
	"github.com/MiladCodeLab/simple-api/repository"
	"log/slog"
)

type UserService interface {
	GetAll() ([]*entity.User, error)
	GetByID(id string) (*entity.User, error)
	DeleteByID(id string) error
	Add(user entity.User) error
}

type userService struct {
	logger *slog.Logger
	repo   repository.UserRepository
}

func NewUserService(logger *slog.Logger, repo repository.UserRepository) UserService {
	return &userService{
		logger: logger.With("service", "user"),
		repo:   repo,
	}
}

func (s *userService) GetAll() ([]*entity.User, error) {
	lg := s.logger.With("method", "GetAll")

	users, err := s.repo.GetAll()
	if err != nil {
		lg.Error("failed to fetch users", "error", err)
		return nil, err
	}

	lg.Debug("users fetched", "count", len(users))
	return users, nil
}

func (s *userService) GetByID(id string) (*entity.User, error) {
	lg := s.logger.With("method", "GetByID", "id", id)

	user, err := s.repo.GetByID(id)
	if err != nil {
		lg.Warn("user not found", "error", err)
		return nil, err
	}

	lg.Debug("user fetched")
	return user, nil
}

func (s *userService) DeleteByID(id string) error {
	lg := s.logger.With("method", "DeleteByID", "id", id)

	err := s.repo.DeleteByID(id)
	if err != nil {
		lg.Warn("delete failed", "error", err)
		return err
	}

	lg.Debug("user deleted")
	return nil
}

func (s *userService) Add(user entity.User) error {
	lg := s.logger.With("method", "Add", "id", user.ID)

	err := s.repo.Add(user)
	if err != nil {
		lg.Error("failed to add user", "error", err)
		return err
	}

	lg.Debug("user added")
	return nil
}
