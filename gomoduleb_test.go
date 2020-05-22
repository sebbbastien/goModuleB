package gomoduleb

import "testing"

func Test_modB(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "base",
			want: "Calling modB() v3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modB(); got != tt.want {
				t.Errorf("modB() = %v, want %v", got, tt.want)
			}
		})
	}
}
