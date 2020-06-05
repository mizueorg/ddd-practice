// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/trewanek-org/ddd-practice/domain/service"
	"github.com/trewanek-org/ddd-practice/infrastructure/inmem"
	"github.com/trewanek-org/ddd-practice/usecase"
)

// Injectors from wire.go:

func GetUserUseCase() *usecase.User {
	iUser := inmem.NewUser()
	user := service.NewUser(iUser)
	usecaseUser := usecase.NewUser(iUser, user)
	return usecaseUser
}

// wire.go:

var serviceSet = wire.NewSet(service.NewUser)

var inMemRepositorySet = wire.NewSet(inmem.NewUser)
