package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/Everesh/crash/builtins"
	"github.com/Everesh/crash/config"
	"github.com/Everesh/crash/parser"
)

type Shell struct {
	reader   *bufio.Reader
	builtins builtins.Registry
}

func New() *Shell {
	return &Shell{
		reader:   bufio.NewReader(os.Stdin),
		builtins: builtins.NewRegistry(),
	}
}

func (s *Shell) Repl() {
	for {
		fmt.Print(config.PS1)

		line, err := s.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				os.Exit(0)
			}

			fmt.Fprintln(os.Stderr, "error reading input:", err)
			os.Exit(1)
		}

		if err := s.Eval(line, os.Stdout); err != nil {
			fmt.Print(err)
		}
	}
}

func (s *Shell) Eval(input string, out io.Writer) error {
	rawCmd, err := parser.Tokenize(input)
	if err != nil {
		return fmt.Errorf("parse error: %w", err)
	} else if len(rawCmd) == 0 {
		return nil
	}

	cmd := rawCmd[0]
	args := rawCmd[1:]

	if builtin, exists := s.builtins[cmd]; exists {
		return builtin.Handle(out, args)
	} else if _, err := exec.LookPath(cmd); err == nil {
		child := exec.Command(cmd, args...)
		child.Stdin = os.Stdin
		child.Stdout = out
		child.Stderr = os.Stderr
		return child.Run()
	} else {
		return fmt.Errorf("%s: command not found\n", cmd)
	}
}
