package cmd

import (
	"context"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel()

	stdout, err := Exec(ctx, "echo wrong >&2 && exit 33")
	t.Logf("stdout: %v", stdout)
	t.Logf("err: %v", err)
}
