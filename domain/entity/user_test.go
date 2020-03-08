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
