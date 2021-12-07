// +build !linux

package pcmd

import (
	"os/exec"
)

func SetPdeathsig(cmd *exec.Cmd) {
	return
}
