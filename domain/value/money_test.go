package value

import (
	"reflect"
	"testing"
)

func TestNewMoney(t *testing.T) {
	type args struct {
		amount   float64
		currency Currency
	}
	tests := []struct {
		name    string
		args    args
		want    *money
		wantErr bool
	}{
		{
			name: "OK1",
			args: args{
				amount:   1.0,
				currency: 1,
			},
			want: &money{
				amount:   1.0,
				currency: 1,
			},
			wantErr: false,
		},
		{
			name: "OK2",
			args: args{
				amount:   1.0,
				currency: 3,
			},
			want: &money{
				amount:   1.0,
				currency: 3,
			},
			wantErr: false,
		},
		{
			name: "NG1",
			args: args{
				amount:   1.0,
				currency: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NG2",
			args: args{
				amount:   1.0,
				currency: 4,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMoney(tt.args.amount, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMoney() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMoney() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_money_Add(t *testing.T) {
	type fields struct {
		amount   float64
		currency Currency
	}
	type args struct {
		arg *money
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *money
		wantErr bool
	}{
		{
			name: "OK1",
			fields: fields{
				amount:   1.0,
				currency: 1,
			},
			args: args{
				arg: &money{
					amount:   2.0,
					currency: 1,
				},
			},
			want: &money{
				amount:   3.0,
				currency: 1,
			},
			wantErr: false,
		},
		{
			name: "OK2",
			fields: fields{
				amount:   -1.0,
				currency: 1,
			},
			args: args{
				arg: &money{
					amount:   2.0,
					currency: 1,
				},
			},
			want: &money{
				amount:   1.0,
				currency: 1,
			},
			wantErr: false,
		},
		{
			name: "NG1",
			fields: fields{
				amount:   1.0,
				currency: 1,
			},
			args: args{
				arg: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NG2",
			fields: fields{
				amount:   1.0,
				currency: 1,
			},
			args: args{
				arg: &money{
					amount:   1.0,
					currency: 2,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &money{
				amount:   tt.fields.amount,
				currency: tt.fields.currency,
			}
			got, err := m.Add(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}
