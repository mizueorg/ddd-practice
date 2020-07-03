package factory

import (
	"strconv"

	"github.com/trewanek-org/ddd-practice/domain/entity"
	"github.com/trewanek-org/ddd-practice/domain/factory"
	"github.com/trewanek-org/ddd-practice/domain/value"
)

type user struct {
	currentID int
}

func (u *user) Create(name string) (*entity.User, error) {
	// ユーザが生成されるたびにインクリメントする
	u.currentID++

	return entity.NewUser(value.UserID(strconv.Itoa(u.currentID)), value.UserName(name))
}

func NewUser() factory.User {
	return &user{}
}
