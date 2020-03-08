package value_object

import "testing"

func TestUserName_Verify(t *testing.T) {
	tests := []struct {
		name    string
		u       UserName
		wantErr bool
	}{
		{
			name:    "OK",
			u:       "hog",
			wantErr: false,
		},
		{
			name:    "OK2",
			u:       "みずし",
			wantErr: false,
		},
		{
			name:    "NG",
			u:       "ho",
			wantErr: true,
		},
		{
			name:    "NG2",
			u:       "みず",
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
