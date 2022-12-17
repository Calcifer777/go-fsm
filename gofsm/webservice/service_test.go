package webservice

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"just run"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Greet()
		})
	}
}
