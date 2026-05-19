package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Everesh/crash/builtins"
)

type Shell struct {
	reader *bufio.Reader
}

func New() *Shell {
	return &Shell{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (s *Shell) Run() {
	for {
		fmt.Print("$ ")

		line, err := s.reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			os.Exit(1)
		}

		rawCmd := strings.Fields(strings.TrimSpace(line))
		if len(rawCmd) == 0 {
			continue
		}

		cmd := rawCmd[0]
		args := rawCmd[1:]

		if handler, exists := builtins.Registry[cmd]; exists {
			handler.Handle(args)

		} else if _, err := exec.LookPath(cmd); err == nil {
			child := exec.Command(cmd, args...)
			child.Stdout = os.Stdout
			child.Stderr = os.Stderr

			if err := child.Run(); err != nil {
				fmt.Fprintln(os.Stderr, "error running command:", err)
			}

		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
