package inmem

import (
	"context"
	"fmt"

	"github.com/trewanek-org/ddd-practice/domain"
	"github.com/trewanek-org/ddd-practice/domain/entity"
	"github.com/trewanek-org/ddd-practice/domain/repository"
	"github.com/trewanek-org/ddd-practice/domain/value"
)

var UserStore = make(map[value.UserName]*entity.User)

func NewUser() repository.IUser {
	return &user{}
}

type user struct{}

func (u *user) Save(ctx context.Context, user *entity.User) error {
	copied := *user
	UserStore[user.UserName()] = &copied
	return nil
}

func (u *user) Find(ctx context.Context, name value.UserName) (*entity.User, error) {
	user := UserStore[name]
	if user == nil {
		return nil, &domain.ErrNotFound{Msg: fmt.Sprintf("user not found error(user name: %v)", name)}
	}

	ret := *user
	return &ret, nil
}
