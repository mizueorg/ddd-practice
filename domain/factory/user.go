package factory

import (
	"github.com/trewanek-org/ddd-practice/domain/entity"
)

type User interface {
	Create(name string) (*entity.User, error)
}
