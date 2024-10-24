package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunningASingleCommand(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	err := runCommand("ls", stdOut, stdErr)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !strings.Contains(stdOut.String(), "command_test.go") {
		t.Errorf("Unexpected output: %s", stdOut.String())
	}
}

func TestRunningASingleCommandWithArguments(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	err := runCommand("ls -al", stdOut, stdErr)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !strings.Contains(stdOut.String(), "command_test.go") ||
		!strings.Contains(stdOut.String(), "-rw-r--r--") {
		t.Errorf("Unexpected output: %s", stdOut.String())
	}
}

func TestRunningMultipleCommands(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	err := runCommand("ls && ps", stdOut, stdErr)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !strings.Contains(stdOut.String(), "command_test.go") ||
		!strings.Contains(stdOut.String(), "PID") {
		t.Errorf("Unexpected output: %s", stdOut.String())
	}
}

func TestRunningMultipleCommandsWithArguments(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	err := runCommand("ls -al && ps -A", stdOut, stdErr)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !strings.Contains(stdOut.String(), "command_test.go") ||
		!strings.Contains(stdOut.String(), "-rw-r--r--") ||
		!strings.Contains(stdOut.String(), "PID") {
		t.Errorf("Unexpected output: %s", stdOut.String())
	}
}
