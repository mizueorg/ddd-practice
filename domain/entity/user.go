package entity

import "github.com/trewanek-org/ddd-practice/domain/value"

type User struct {
	userID   value.UserID
	userName value.UserName
}

func NewUser(userID value.UserID, userName value.UserName) (*User, error) {
	if err := userName.Verify(); err != nil {
		return nil, err
	}
	return &User{userID: userID, userName: userName}, nil
}

func (u *User) ChangeName(newName value.UserName) error {
	if err := newName.Verify(); err != nil {
		return err
	}
	u.userName = newName
	return nil
}

func (u *User) Equal(arg *User) bool {
	if arg == nil {
		return false
	}
	return u.userID == arg.userID
}
