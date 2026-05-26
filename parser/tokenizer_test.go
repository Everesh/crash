package parser

import "testing"

func w(s string) Token { return Token{Kind: Word, Value: s} }

func TestTokenize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Token
		wantErr bool
	}{
		{
			name:  "simple command with args",
			input: "ls -la /tmp",
			want:  []Token{w("ls"), w("-la"), w("/tmp")},
		},
		{
			name:  "multiple spaces between args",
			input: "echo    hello   world",
			want:  []Token{w("echo"), w("hello"), w("world")},
		},
		{
			name:  "double quotes preserve spaces",
			input: `echo "hello   world"`,
			want:  []Token{w("echo"), w("hello   world")},
		},
		{
			name:  "single quotes preserve spaces",
			input: `echo 'hello   world'`,
			want:  []Token{w("echo"), w("hello   world")},
		},
		{
			name:  "adjacent quotes merge into one token",
			input: `"hello"world`,
			want:  []Token{w("helloworld")},
		},
		{
			name:  "backslash is literal inside single quotes",
			input: `echo 'dir\file.txt'`,
			want:  []Token{w("echo"), w(`dir\file.txt`)},
		},
		{
			name:  "backslash before non-special char in double quotes is kept",
			input: `echo "dir\file.txt"`,
			want:  []Token{w("echo"), w(`dir\file.txt`)},
		},
		{
			name:  "backslash escapes space outside quotes",
			input: `echo hello\ world`,
			want:  []Token{w("echo"), w("hello world")},
		},
		{
			name:  "single quotes nested inside double quotes are literal",
			input: `echo "    blue: '           '"    #0000FF`,
			want:  []Token{w("echo"), w("    blue: '           '"), w("#0000FF")},
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
		// operator tokens
		{
			name:  "redirect stdout",
			input: "echo hello > file.txt",
			want:  []Token{w("echo"), w("hello"), {Kind: RedirectOut}, w("file.txt")},
		},
		{
			name:  "redirect stdout append",
			input: "echo hi >> log",
			want:  []Token{w("echo"), w("hi"), {Kind: RedirectAppend}, w("log")},
		},
		{
			name:  "redirect stdin",
			input: "cat < in.txt",
			want:  []Token{w("cat"), {Kind: RedirectIn}, w("in.txt")},
		},
		{
			name:  "redirect stderr",
			input: "cmd 2> err",
			want:  []Token{w("cmd"), {Kind: RedirectErr}, w("err")},
		},
		{
			name:  "redirect stderr append",
			input: "cmd 2>> err",
			want:  []Token{w("cmd"), {Kind: RedirectErrAppend}, w("err")},
		},
		{
			name:  "redirect both stdout and stderr",
			input: "cmd &> out",
			want:  []Token{w("cmd"), {Kind: RedirectBoth}, w("out")},
		},
		{
			name:  "pipe",
			input: "ls | grep foo",
			want:  []Token{w("ls"), {Kind: Pipe}, w("grep"), w("foo")},
		},
		{
			name:  "semicolon",
			input: "ls ; pwd",
			want:  []Token{w("ls"), {Kind: Semicolon}, w("pwd")},
		},
		{
			name:  "and",
			input: "make && ./run",
			want:  []Token{w("make"), {Kind: And}, w("./run")},
		},
		{
			name:  "or",
			input: "cmd || fallback",
			want:  []Token{w("cmd"), {Kind: Or}, w("fallback")},
		},
		{
			name:  "redirect without surrounding spaces",
			input: "echo hello>file",
			want:  []Token{w("echo"), w("hello"), {Kind: RedirectOut}, w("file")},
		},
		{
			name:  "operator inside double quotes is a plain word",
			input: `echo "hello > world"`,
			want:  []Token{w("echo"), w("hello > world")},
		},
		{
			name:  "digit not followed by redirect is a plain word",
			input: "2plus2",
			want:  []Token{w("2plus2")},
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
					t.Errorf("Tokenize(%q) token[%d]\n  got  %+v\n  want %+v", tt.input, i, got[i], tt.want[i])
				}
			}
		})
	}
}
