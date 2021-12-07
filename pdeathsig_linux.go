// +build linux

package pcmd

import (
	"os"
	"os/exec"
	"syscall"
)

func SetPdeathsig(cmd *exec.Cmd) {
	if _, ok := os.LookupEnv("LAMBDA_TASK_ROOT"); ok {
		// do nothing on AWS Lambda
		return
	}
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = new(syscall.SysProcAttr)
	}
	// When the parent thread dies (Go), kill the child as well.
	cmd.SysProcAttr.Setpgid = true
	cmd.SysProcAttr.Pdeathsig = syscall.SIGKILL
}
