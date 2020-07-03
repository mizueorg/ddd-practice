package usecase

import (
	"context"

	"github.com/trewanek-org/ddd-practice/domain/factory"

	"github.com/trewanek-org/ddd-practice/domain"
	"github.com/trewanek-org/ddd-practice/domain/repository"
	"github.com/trewanek-org/ddd-practice/domain/service"
)

type User struct {
	userFactory    factory.User
	userRepository repository.IUser
	userService    *service.User
}

func NewUser(userFactory factory.User, userRepository repository.IUser, userService *service.User) *User {
	return &User{userFactory: userFactory, userRepository: userRepository, userService: userService}
}

func (u *User) Register(ctx context.Context, name string) error {
	user, err := u.userFactory.Create(name)
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
