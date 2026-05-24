package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/Everesh/crash/parser"
)

func (s *Shell) Eval(input string, out io.Writer) error {
	rawCmd, err := parser.Tokenize(input)
	if err != nil {
		return fmt.Errorf("parse error: %s\n", err)
	}
	if len(rawCmd) == 0 {
		return nil
	}

	cmd := rawCmd[0]
	args := rawCmd[1:]

	if builtin, exists := s.builtins[cmd]; exists {
		if err := builtin.Handle(out, args); err != nil {
			return err
		}
		return nil
	}

	if _, err := exec.LookPath(cmd); err != nil {
		return fmt.Errorf("%s: command not found\n", cmd)
	}

	child := exec.Command(cmd, args...)
	child.Stdin = os.Stdin
	child.Stdout = out
	child.Stderr = os.Stderr
	if err := child.Run(); err != nil {
		return err
	}
	return nil
}
