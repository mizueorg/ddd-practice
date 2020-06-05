package domain

import (
	"fmt"
)

type ErrCanNotRegisterUser struct {
	Msg string
}

func (e ErrCanNotRegisterUser) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}
