package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// TODO - handle flags -v -p

func handleCommand(_ Registry, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("command: missing argument\n")
	}

	cmd := args[0]
	if strings.HasPrefix(cmd, "-") {
		return "", fmt.Errorf("command: flags not yet implemented\n")
	}

	if _, err := exec.LookPath(cmd); err != nil {
		return "", fmt.Errorf("command: %s: no such command in path\n", cmd)
	}

	child := exec.Command(cmd, args[1:]...)
	child.Stdin = os.Stdin
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr

	if err := child.Run(); err != nil {
		return "", fmt.Errorf("command: error running command:", err)
	}

	return "", nil
}

func tldrCommand() string {
	// TODO
	return ""
}

func manCommand() string {
	// TODO
	return ""
}
