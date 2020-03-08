package value_object

import "errors"

type UserName string

func (u *UserName) Verify() error {
	r := []rune(string(*u))
	if len(r) < 3 {
		return errors.New("username should be grater than 3 char")
	}
	return nil
}
