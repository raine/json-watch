package main

import (
	"bufio"
	"os"
)

func readLinesMap(file *os.File) (map[string]struct{}, error) {
	lines := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines[scanner.Text()] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	} else {
		return lines, nil
	}
}
