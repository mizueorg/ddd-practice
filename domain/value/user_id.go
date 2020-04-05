package value

import "errors"

type UserID string

func (u *UserID) Verify() error {
	if string(*u) == "" {
		return errors.New("user id is not empty")
	}
	return nil
}
