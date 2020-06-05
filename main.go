package main

import (
	"github.com/trewanek-org/ddd-practice/di"
)

func main() {
	userUserCase := di.GetUserUseCase()
	_ = userUserCase
}
