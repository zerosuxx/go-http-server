package utility

import (
	"io"
	"os/exec"
)

type Shell struct {

}

func (s Shell) Run(command string, args []string, output io.Writer) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = output
	cmd.Stderr = output

	return cmd.Run()
}
