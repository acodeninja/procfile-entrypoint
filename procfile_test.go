package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestParsingAProcfile(t *testing.T) {
	procfile := bufio.NewReader(strings.NewReader("web: echo hello"))

	parsedProcfile, err := parseProcfile(procfile)
	if err != nil {
		t.Fatalf("Failed to parse procfile: %s", err)
	}

	if parsedProcfile == nil {
		t.Fatalf("No entries found in procfile")
	}

	if parsedProcfile["web"] != "echo hello" {
		t.Fatalf("Wrong entry found in procfile: %v", parsedProcfile)
	}
}
