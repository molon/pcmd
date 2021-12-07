package pcmd

import (
	"context"
	"os/exec"
)

var (
	defaultStarter = NewStarter()
)

func Start(ctx context.Context, cmd *exec.Cmd) error {
	return defaultStarter.Start(ctx, cmd)
}
