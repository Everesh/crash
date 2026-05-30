package shell

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Everesh/crash/builtins"
	"github.com/Everesh/crash/config"
	"github.com/chzyer/readline"
)

func (s *Shell) Repl() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       config.PS1,
		VimMode:      config.VimMode,
		AutoComplete: completer(s),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "readline:", err)
		os.Exit(1)
	}
	defer rl.Close()

	for {
		line, err := readLine(rl)
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

		if line == "" {
			continue
		}

		if err := s.Eval(line, os.Stdout); err != nil {
			var exitErr *builtins.ExitError
			if errors.As(err, &exitErr) {
				rl.Close()
				os.Exit(exitErr.Code)
			}
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func readLine(rl *readline.Instance) (string, error) {
	var sb strings.Builder
	prompt := config.PS1
	for {
		rl.SetPrompt(prompt)
		seg, err := rl.Readline()
		if err != nil {
			rl.SetPrompt(config.PS1)
			if prompt == config.PS2 {
				// error during continuation: discard partial line
				return "", readline.ErrInterrupt
			}
			return "", err
		}
		if strings.HasSuffix(seg, "\\") {
			sb.WriteString(seg[:len(seg)-1])
			prompt = config.PS2
			continue
		}
		sb.WriteString(seg)
		break
	}
	rl.SetPrompt(config.PS1)
	return sb.String(), nil
}
