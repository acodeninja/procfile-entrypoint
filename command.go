package main

import (
	"io"
	"os/exec"
	"strings"
)

func runCommand(command string, stdOut io.Writer, stdErr io.Writer) error {
	for _, c := range strings.Split(command, "&&") {
		args := strings.Fields(c)
		cmd := exec.Command(args[0], args[1:]...)

		cmd.Stdout = stdOut
		cmd.Stderr = stdErr
		err := cmd.Run()

		if err != nil {
			return err
		}
	}
	return nil
}
