package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestInputOutput(t *testing.T) {
	var result = []struct {
		command  string
		expected string
	}{
		{"help\n", Help},
		{"h\n", Help},
	}

	for _, res := range result {
		var buf bytes.Buffer
		StartCli(strings.NewReader(res.command), &buf)
		if buf.String() != res.expected {
			t.Errorf("Unexpected output for command %s", res.command)
		}
	}

}
