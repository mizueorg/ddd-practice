package entity

import (
	"github.com/trewanek-org/ddd-practice/domain/value_object"
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		userID   value_object.UserID
		userName value_object.UserName
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID:   value_object.UserID("hoge-id"),
				userName: value_object.UserName("hoge-name"),
			},
			want: &User{
				userID:   value_object.UserID("hoge-id"),
				userName: value_object.UserName("hoge-name"),
			},
			wantErr: false,
		},
		{
			name: "NG",
			args: args{
				userID:   value_object.UserID("hoge-id"),
				userName: value_object.UserName("ho"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.userID, tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ChangeName(t *testing.T) {
	type fields struct {
		userID   value_object.UserID
		userName value_object.UserName
	}
	type args struct {
		newName value_object.UserName
	}
	type want struct {
		changedUserName value_object.UserName
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				userID:   "hoge-id",
				userName: "hoge-name",
			},
			args: args{
				newName: "hoge-name-new",
			},
			want: want{
				changedUserName: "hoge-name-new",
			},
			wantErr: false,
		},
		{
			name: "NG",
			fields: fields{
				userID:   "hoge-id",
				userName: "hoge-name",
			},
			args: args{
				newName: "ho",
			},
			want: want{
				changedUserName: "hoge-name",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				userID:   tt.fields.userID,
				userName: tt.fields.userName,
			}
			if err := u.ChangeName(tt.args.newName); (err != nil) != tt.wantErr {
				t.Errorf("ChangeName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if u.userName != tt.want.changedUserName {
				t.Errorf("ChangeName() it is not working, actual value = %v, want %v", u.userName, tt.want.changedUserName)
			}
		})
	}
}
