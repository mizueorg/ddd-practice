package domain

import (
	"fmt"
)

type ErrNotFound struct {
	Msg string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}
