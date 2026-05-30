package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	t "github.com/Everesh/crash/parser/tokens"
)

func (s *Shell) Eval(input string, out io.Writer) error {
	tokens, err := t.Tokenize(input)
	if err != nil {
		return fmt.Errorf("parse error: %s\n", err)
	}
	if len(tokens) == 0 {
		return nil
	}

	var words []string
	for _, tok := range tokens {
		switch tok.Kind {
		case t.Word:
			words = append(words, tok.Value)
		default:
			return fmt.Errorf("%v: operator not yet supported\n", tok.Kind)
		}
	}

	if len(words) == 0 {
		return nil
	}

	cmd := words[0]
	args := words[1:]

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
