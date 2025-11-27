package repository

import (
	"errors"
	"github.com/MiladCodeLab/simple-api/entity"
	"log/slog"
	"sync"
)

var (
	ErrNotFoundUser = errors.New("user not found")
)

type UserRepository interface {
	GetAll() ([]*entity.User, error)
	Add(user entity.User) error
	GetByID(id string) (*entity.User, error)
	DeleteByID(id string) error
}

type userRepository struct {
	data   map[string]entity.User
	logger *slog.Logger
	mu     sync.Mutex
}

func NewUserRepository(logger *slog.Logger) UserRepository {
	return &userRepository{
		data:   make(map[string]entity.User),
		logger: logger.With("repository", "user"),
	}
}

func (r *userRepository) GetAll() ([]*entity.User, error) {
	lg := r.logger.With("method", "GetAll")

	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]*entity.User, 0, len(r.data))
	for _, u := range r.data {
		user := u // copy
		users = append(users, &user)
	}

	lg.Debug("fetched all users", "count", len(users))
	return users, nil
}

func (r *userRepository) Add(user entity.User) error {
	lg := r.logger.With("method", "Add", "id", user.ID)

	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[user.ID] = user
	lg.Debug("user added")
	return nil
}

func (r *userRepository) GetByID(id string) (*entity.User, error) {
	lg := r.logger.With("method", "GetByID", "id", id)

	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.data[id]
	if !ok {
		lg.Warn("user not found")
		return nil, ErrNotFoundUser
	}

	lg.Debug("user fetched")
	return &user, nil
}

func (r *userRepository) DeleteByID(id string) error {
	lg := r.logger.With("method", "DeleteByID", "id", id)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		lg.Warn("user not found")
		return ErrNotFoundUser
	}

	delete(r.data, id)
	lg.Debug("user deleted")
	return nil
}
