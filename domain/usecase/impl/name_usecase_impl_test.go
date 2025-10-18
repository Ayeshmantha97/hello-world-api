package impl

import (
	"context"
	"testing"
)

func TestGetName(t *testing.T) {
	n := &nameUseCase{}

	ctx := context.Background()

	tests := []struct {
		input    string
		isValid  bool
		response string
	}{
		{"Alice", true, "Hello Alice"},
		{"Mike", true, "Hello Mike"},
		{"Zara", false, ""},
		{"", false, ""},
		{"1Alice", false, ""},
		{"@Mike", false, ""},
	}

	for _, tt := range tests {
		valid, resp := n.GetName(ctx, tt.input)
		if valid != tt.isValid || resp != tt.response {
			t.Errorf("GetName(%q) = (%v, %q), want (%v, %q)", tt.input, valid, resp, tt.isValid, tt.response)
		}
	}
}
