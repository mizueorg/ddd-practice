package value_object

import "errors"

type UserID string

func (u *UserID) Verify() error {
	if string(*u) == "" {
		return errors.New("user id is not empty")
	}
	return nil
}
