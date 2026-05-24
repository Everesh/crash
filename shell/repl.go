package shell

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Everesh/crash/config"
	"github.com/chzyer/readline"
)

func (s *Shell) Repl() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       config.PS1,
		VimMode:      config.VimMode,
		AutoComplete: s.completer(),
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

func (s *Shell) completer() readline.AutoCompleter {
	items := make([]readline.PrefixCompleterInterface, 0, len(s.builtins))
	for name := range s.builtins {
		items = append(items, readline.PcItem(name))
	}
	return readline.NewPrefixCompleter(items...)
}
