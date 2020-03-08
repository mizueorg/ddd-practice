package entity

import "github.com/trewanek-org/ddd-practice/domain/value_object"

type User struct {
	userID   value_object.UserID
	userName value_object.UserName
}

func NewUser(userID value_object.UserID, userName value_object.UserName) (*User, error) {
	if err := userName.Verify(); err != nil {
		return nil, err
	}
	return &User{userID: userID, userName: userName}, nil
}

func (u *User) ChangeName(newName value_object.UserName) error {
	if err := newName.Verify(); err != nil {
		return err
	}
	u.userName = newName
	return nil
}
