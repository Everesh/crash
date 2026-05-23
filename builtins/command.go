package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// TODO - handle flags -v -p

func handleCommand(_ Registry, args []string) {
	if len(args) < 1 {
		fmt.Println("command: missing argument")
		return
	}

	cmd := args[0]
	if strings.HasPrefix(cmd, "-") {
		fmt.Println("command: flags not yet implemented")
		return
	}

	if _, err := exec.LookPath(cmd); err != nil {
		fmt.Printf("command: %s: no such command in path\n", cmd)
		return
	}

	child := exec.Command(cmd, args[1:]...)
	child.Stdin = os.Stdin
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr

	if err := child.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "command: error running command:", err)
		return
	}

}

func tldrCommand() string {
	// TODO
	return ""
}

func manCommand() string {
	// TODO
	return ""
}
