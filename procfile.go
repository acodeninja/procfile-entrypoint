package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getProcfileCommand(entrypoint string) (string, error) {
	file, err := openFile("Procfile")
	if err != nil {
		return "", err
	}

	parsed, err := parseProcfile(file)
	if err != nil {
		return "", err
	}

	return parsed[entrypoint], nil
}

func parseProcfile(procfile *bufio.Reader) (map[string]string, error) {
	scanner := bufio.NewScanner(procfile)

	parsed := make(map[string]string)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" ||
			strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, "//") {
			continue
		}

		name, command, found := strings.Cut(line, ":")
		if !found {
			continue
		}

		parsed[strings.TrimSpace(name)] = strings.TrimSpace(command)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return parsed, nil
}

func openFile(filename string) (*bufio.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	return bufio.NewReader(file), nil
}
