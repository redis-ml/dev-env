package handler

import (
	"context"
	"testing"
)

func TestXxx(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		input    int
		expected string
	}{
		{1, "rh:723abb-1d-0000-000001"},
		{2, "rh:41ad2f-79-0000-000002"},
	}
	for idx, c := range cases {
		obtained := transformToOwnerID(ctx, c.input)
		if obtained != c.expected {
			t.Errorf("case %d: transformToOwnerID(%d) = \"%s\", expecting: \"%s\"",
				idx, c.input, transformToOwnerID(ctx, c.input), c.expected)
		}
	}
}
