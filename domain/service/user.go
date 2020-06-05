package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/trewanek-org/ddd-practice/domain"
	"github.com/trewanek-org/ddd-practice/domain/entity"
	"github.com/trewanek-org/ddd-practice/domain/repository"
)

func NewUser(userRepository repository.IUser) *User {
	return &User{userRepository: userRepository}
}

type User struct {
	userRepository repository.IUser
}

func (s *User) Exists(ctx context.Context, user *entity.User) (bool, error) {
	if user == nil {
		return false, fmt.Errorf("argument user nil")
	}

	if _, err := s.userRepository.Find(ctx, user.UserName()); err != nil {
		var errNotFound *domain.ErrNotFound
		if errors.As(err, &errNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("error: %v", err)
	}

	return true, nil
}
