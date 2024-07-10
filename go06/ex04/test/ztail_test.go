package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
	//"piscine"
)

func TestZtail(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		input    string // for stdint
		want     string
		exitCode int
	}{
		// error case exit status 1
		{
			name:     "Error Case, \"\"",
			args:     []string{""},
			want:     "ztail: only the -c option is supported.",
			exitCode: 1,
		},
		{
			name:     "Error Case, -c \"\"",
			args:     []string{"-c"},
			want:     "ztail: option requires an argument -- 'c'\n",
			exitCode: 1,
		},
		{
			name:     "Error Case, -c a",
			args:     []string{"-c", "a"},
			want:     "ztail: invalid number of bytes: 'a'\n",
			exitCode: 1,
		},
		{
			name:     "Error Case, -c 1 a",
			args:     []string{"-c", "1", "a"},
			want:     "open a: no such file or directory\n",
			exitCode: 1,
		},
		{
			name:     "Error Case, -c 3 a test.txt",
			args:     []string{"-c", "3", "a", "test.txt"},
			want:     "==> a <==\nopen a: no such file or directory\n" + "==> test.txt <==\nxyz",
			exitCode: 1,
		},

		// valid case
		{
			name:     "Valid Case, -c 5 test.txt",
			args:     []string{"-c", "5", "test.txt"},
			want:     "vwxyz",
			exitCode: 0,
		},
		{
			name:     "Valid Case, -c 10",
			args:     []string{"-c", "10"},
			input:    "10 characters",
			want:     "characters",
			exitCode: 0,
		},
		{
			name:     "Valid Case, -c 5 - test.txt",
			args:     []string{"-c", "5", "-", "test.txt"},
			input:    "tokyo",
			want:     "==> standard input <==\ntokyo" + "==> test.txt <==\nvwxyz",
			exitCode: 0,
		},
		{
			name:     "Valid Case, -c 5 test.txt -",
			args:     []string{"-c", "5", "test.txt", "-"},
			input:    "tokyo",
			want:     "==> test.txt <==\nvwxyz" + "==> standard input <==\ntokyo",
			exitCode: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command("./ztail", test.args...)

			if test.input != "" {
				cmd.Stdin = strings.NewReader(test.input)
			}

			var output bytes.Buffer
			cmd.Stdout = &output
			cmd.Stderr = &output

			err := cmd.Run()

			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() != test.exitCode {
					t.Errorf("Expected exit code %d, but got %d", test.exitCode, exitError.ExitCode())
				}
			} else if err != nil && test.exitCode == 0 {
				t.Errorf("Unexpected error: %v", err)
			}

			// 出力をチェック
			if !strings.Contains(output.String(), test.want) {
				t.Errorf("Expected output to contain %q, but got %q", test.want, output.String())
			}
		})
	}
}
