package parser

import "testing"

func TestTokenize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:  "simple command with args",
			input: "ls -la /tmp",
			want:  []string{"ls", "-la", "/tmp"},
		},
		{
			name:  "multiple spaces between args",
			input: "echo    hello   world",
			want:  []string{"echo", "hello", "world"},
		},
		{
			name:  "double quotes preserve spaces",
			input: `echo "hello   world"`,
			want:  []string{"echo", "hello   world"},
		},
		{
			name:  "single quotes preserve spaces",
			input: `echo 'hello   world'`,
			want:  []string{"echo", "hello   world"},
		},
		{
			name:  "adjacent quotes merge into one token",
			input: `"hello"world`,
			want:  []string{"helloworld"},
		},
		{
			name:  "backslash is literal inside single quotes",
			input: `echo 'dir\file.txt'`,
			want:  []string{"echo", `dir\file.txt`},
		},
		{
			name:  "backslash before non-special char in double quotes is kept",
			input: `echo "dir\file.txt"`,
			want:  []string{"echo", `dir\file.txt`},
		},
		{
			name:  "backslash escapes space outside quotes",
			input: `echo hello\ world`,
			want:  []string{"echo", "hello world"},
		},
		{
			name:  "single quotes nested inside double quotes are literal",
			input: `echo "    blue: '           '"    #0000FF`,
			want:  []string{"echo", "    blue: '           '", "#0000FF"},
		},
		{
			name:  "empty input",
			input: "",
			want:  nil,
		},
		{
			name:    "unterminated double quote",
			input:   `echo "unterminated`,
			wantErr: true,
		},
		{
			name:    "unterminated single quote",
			input:   `echo 'unterminated`,
			wantErr: true,
		},
		{
			name:    "malnested quotes produce unterminated error",
			input:   `echo "    blue: '           " '   #0000FF`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tokenize(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("Tokenize(%q)\n  got  %v\n  want %v", tt.input, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Tokenize(%q) token[%d]\n  got  %q\n  want %q", tt.input, i, got[i], tt.want[i])
				}
			}
		})
	}
}
