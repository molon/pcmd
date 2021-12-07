// +build linux

package pcmd

import (
	"context"
	"os/exec"
	"runtime"
	"sync"
)

type startWrapper struct {
	cmd   *exec.Cmd
	errCh chan error
}

type Starter struct {
	once sync.Once
	ch   chan *startWrapper
}

func NewStarter() *Starter {
	return &Starter{
		ch: make(chan *startWrapper),
	}
}

func (s *Starter) Start(ctx context.Context, cmd *exec.Cmd) error {
	s.once.Do(func() {
		go func() {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()

			for v := range s.ch {
				v.errCh <- v.cmd.Start()
			}
		}()
	})
	w := &startWrapper{
		cmd:   cmd,
		errCh: make(chan error, 1),
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.ch <- w:
	}
	return <-w.errCh
}
