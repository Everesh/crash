package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"

	"github.com/Everesh/crash/builtins"
	"github.com/Everesh/crash/config"
	"github.com/Everesh/crash/parser"
)

type Shell struct {
	builtins builtins.Registry
}

func New() *Shell {
	return &Shell{builtins: builtins.NewRegistry()}
}

func (s *Shell) Repl() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:  config.PS1,
		VimMode: config.VimMode,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "readline:", err)
		os.Exit(1)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			continue
		}
		if err == io.EOF {
			fmt.Println()
			return
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			rl.Close()
			os.Exit(1)
		}

		for strings.HasSuffix(line, "\\") {
			line = line[:len(line)-1]
			rl.SetPrompt(config.PS2)
			cont, err := rl.Readline()
			if err == readline.ErrInterrupt || err == io.EOF {
				line = ""
				break
			}
			line += cont
		}
		rl.SetPrompt(config.PS1)

		if line == "" {
			continue
		}

		if err := s.Eval(line, os.Stdout); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}
}

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
