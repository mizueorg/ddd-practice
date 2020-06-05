package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/trewanek-org/ddd-practice/domain"
	"github.com/trewanek-org/ddd-practice/domain/entity"
	"github.com/trewanek-org/ddd-practice/domain/repository"
	"github.com/trewanek-org/ddd-practice/domain/service"
	"github.com/trewanek-org/ddd-practice/domain/value"
)

type User struct {
	userRepository repository.IUser
	userService    *service.User
}

func NewUser(userRepository repository.IUser, userService *service.User) *User {
	return &User{userRepository: userRepository, userService: userService}
}

func (u *User) Register(ctx context.Context, name string) error {
	user, err := entity.NewUser(value.UserID(uuid.New().String()), value.UserName(name))
	if err != nil {
		return err
	}
	exist, err := u.userService.Exists(ctx, user)
	if err != nil {
		return err
	}

	if exist {
		return &domain.ErrCanNotRegisterUser{
			Msg: "ユーザはすでに存在しています。",
		}
	}

	return u.userRepository.Save(ctx, user)
}
