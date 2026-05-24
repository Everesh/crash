package builtins

import (
	"fmt"
	"os/exec"
)

func handleType(builtins Registry, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("type: missing argument\n")
	}

	if len(args) > 1 {
		return "", fmt.Errorf("type: too many arguments\n")
	}

	cmd := args[0]

	if _, exists := builtins[cmd]; exists {
		return fmt.Sprintf("%s is a shell builtin\n", cmd), nil
	} else if path, _ := exec.LookPath(cmd); path != "" {
		return fmt.Sprintf("%s is %s\n", cmd, path), nil
	} else {
		return fmt.Sprintf("%s: not found\n", cmd), nil
	}
}

func tldrType() string {
	// TODO
	return ""
}

func manType() string {
	// TODO
	return ""
}
