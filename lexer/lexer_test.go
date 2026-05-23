package lexer

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Simple command with spaces",
			input:    "ls -la /tmp",
			expected: []string{"ls", "-la", "/tmp"},
		},
		{
			name:     "Multiple spaces between arguments",
			input:    "echo    hello   world",
			expected: []string{"echo", "hello", "world"},
		},
		{
			name:     "Double quoted string preserving spaces",
			input:    `echo "hello   world"`,
			expected: []string{"echo", "hello   world"},
		},
		{
			name:     "Single quoted string preserving spaces",
			input:    `echo 'hello   world'`,
			expected: []string{"echo", "hello   world"},
		},
		{
			name:     "Juxtaposed quote boundaries without spaces",
			input:    `"hello"world`,
			expected: []string{"helloworld"},
		},
		{
			name:     "Preserve backslashes inside quotes",
			input:    `echo "dir\file.txt"`,
			expected: []string{"echo", `dir\file.txt`},
		},
		{
			name:     "Escape character handling outside quotes",
			input:    `echo hello\ world`,
			expected: []string{"echo", "hello world"},
		},
		{
			name:     "POSIX unterminated quote guard returns empty slice",
			input:    `echo "unterminated`,
			expected: []string{},
		},
		{
			name:     "Empty string input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Nested quotation marks",
			input:    `echo "    blue: '           '"    #0000FF`,
			expected: []string{"echo", "    blue: '           '", "#0000FF"},
		},
		{
			name:     "Malnested quotation marks",
			input:    `echo "    blue: '           " '   #0000FF`,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.expected) == 0 && tt.input != "" {
				null, _ := os.Open(os.DevNull)
				oldStdout := os.Stdout
				os.Stdout = null

				defer func() {
					os.Stdout = oldStdout
					null.Close()
				}()
			}

			actual := Parse(tt.input)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("\nParser failed for test: %q\nInput:    %s\nExpected: %v\nActual:   %v",
					tt.name, tt.input, tt.expected, actual)
			}
		})
	}
}
