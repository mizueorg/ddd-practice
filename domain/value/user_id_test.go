package value

import "testing"

func TestUserID_Verify(t *testing.T) {
	tests := []struct {
		name    string
		u       UserID
		wantErr bool
	}{
		{
			name:    "OK",
			u:       "id1",
			wantErr: false,
		},
		{
			name:    "NG",
			u:       "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Verify(); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
