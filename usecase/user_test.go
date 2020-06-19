package usecase_test

import (
	"context"
	"testing"

	"github.com/trewanek-org/ddd-practice/di"
)

func TestUser_Register(t *testing.T) {
	ctx := context.Background()
	use := di.GetUserUseCase()

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx:  ctx,
				name: "mizushima",
			},
			wantErr: false,
		},
		{
			name: "文字数足りない",
			args: args{
				ctx:  ctx,
				name: "ta",
			},
			wantErr: true,
		},
		{
			name: "重複エラー",
			args: args{
				ctx:  ctx,
				name: "mizushima",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := use.Register(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
