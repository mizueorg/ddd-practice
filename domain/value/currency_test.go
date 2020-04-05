package value

import "testing"

func TestCurrency_IsUndefined(t *testing.T) {
	tests := []struct {
		name string
		c    Currency
		want bool
	}{
		{
			name: "JPY",
			c:    1,
			want: false,
		},
		{
			name: "USD",
			c:    2,
			want: false,
		},
		{
			name: "EUR",
			c:    3,
			want: false,
		},
		{
			name: "Undefined",
			c:    4,
			want: true,
		},
		{
			name: "Zero",
			c:    0,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsUndefined(); got != tt.want {
				t.Errorf("IsUndefined() = %v, want %v", got, tt.want)
			}
		})
	}
}
