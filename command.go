package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func runCommand(command string, stdOut io.Writer, stdErr io.Writer) error {
	for _, c := range strings.Split(command, "&&") {
		cmd := exec.Command("bash", "-c", c)

		cmd.Stdout = stdOut
		cmd.Stderr = stdErr
		cmd.Env = os.Environ()
		err := cmd.Run()

		if err != nil {
			return err
		}
	}

	return nil
}
