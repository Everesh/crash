package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/Everesh/crash/builtins"
	"github.com/Everesh/crash/config"
	"github.com/Everesh/crash/lexer"
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

func (s *Shell) Run() {
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

		rawCmd := lexer.Parse(line)
		if len(rawCmd) == 0 {
			continue
		}

		cmd := rawCmd[0]
		args := rawCmd[1:]

		if builtin, exists := s.builtins[cmd]; exists {
			// builtins is a map, maps pass by reference, no `&s.builtins` is needed
			builtin.Handle(s.builtins, args)

		} else if _, err := exec.LookPath(cmd); err == nil {
			child := exec.Command(cmd, args...)
			child.Stdin = os.Stdin
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
