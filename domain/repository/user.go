package repository

import (
	"context"

	"github.com/trewanek-org/ddd-practice/domain/entity"
	"github.com/trewanek-org/ddd-practice/domain/value"
)

type IUser interface {
	Save(ctx context.Context, user *entity.User) error
	Find(ctx context.Context, name value.UserName) (*entity.User, error)
}
