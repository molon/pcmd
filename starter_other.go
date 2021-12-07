// +build !linux

package pcmd

import (
	"context"
	"os/exec"
)

type Starter struct{}

func NewStarter() *Starter {
	return &Starter{}
}

func (s *Starter) Start(ctx context.Context, cmd *exec.Cmd) error {
	return cmd.Start()
}
