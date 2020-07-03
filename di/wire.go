//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/trewanek-org/ddd-practice/domain/service"
	"github.com/trewanek-org/ddd-practice/infrastructure/factory"
	"github.com/trewanek-org/ddd-practice/infrastructure/inmem"
	"github.com/trewanek-org/ddd-practice/usecase"
)

var serviceSet = wire.NewSet(
	service.NewUser,
)

var factorySet = wire.NewSet(
	factory.NewUser,
)

var inMemRepositorySet = wire.NewSet(
	inmem.NewUser,
)

func GetUserUseCase() *usecase.User {
	wire.Build(
		usecase.NewUser,
		service.NewUser,
		factorySet,
		inMemRepositorySet,
	)

	return &usecase.User{}
}
