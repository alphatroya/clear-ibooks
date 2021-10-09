package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	openP  = "«"
	closeP = "»"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var br bytes.Buffer
	tee := io.TeeReader(reader, &br)
	out := read(tee)
	if out == "" {
		out = br.String()
	}
	fmt.Fprintln(os.Stdout, out)
}

func read(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)

	var result string
	captureMode := false
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if captureMode || strings.HasPrefix(line, openP) {
			captureMode = true
			result += strings.TrimLeft(line, openP)
			if strings.HasSuffix(result, closeP) {
				result = strings.TrimRight(result, closeP)
				break
			}
			result += "\n"
		}
	}

	err := scanner.Err()
	if err != nil || len(result) == 0 {
		return ""
	}
	return result
}
