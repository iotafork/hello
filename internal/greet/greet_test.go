package greet

import (
	"errors"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{"simple", "Mona", "Hello, Mona!", nil},
		{"trims spaces", "  Mona  ", "Hello, Mona!", nil},
		{"empty", "", "", ErrEmptyName},
		{"only spaces", "   ", "", ErrEmptyName},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Hello(tc.input)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Hello(%q) error = %v, want %v", tc.input, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Hello(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
