package shell

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Everesh/crash/config"
	"github.com/chzyer/readline"
)

func (sh *Shell) Repl() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       config.PS1,
		VimMode:      config.VimMode,
		AutoComplete: completer(sh),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "readline:", err)
		os.Exit(1)
	}
	defer rl.Close()

	handler := signalHandler{rl}

	for {
		line, err := readLine(rl)
		if err == readline.ErrInterrupt {
			continue
		}
		if err == io.EOF {
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

		sh.Eval(line)

		for _, cmd := range sh.signals.Drain() {
			cmd.Apply(handler)
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

// --- Signal Handler ---

type signalHandler struct{ rl *readline.Instance }

func (h signalHandler) Exit(code int) {
	h.rl.Close()
	os.Exit(code)
}
