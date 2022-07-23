package handler

import (
	"context"
	"testing"
)

func TestXxx(t *testing.T) {
	ctx := context.Background()
	t.Log("TestXxx")

	if transformToOwnerID(ctx, 1) != "" {
		t.Errorf("transformToOwnerID(%d) != \"\"", 1)
	}
}
