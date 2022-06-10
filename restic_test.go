package restic

import (
	"context"
	"testing"
)

func TestRestic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := New(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Log(r.CmdString)
}
