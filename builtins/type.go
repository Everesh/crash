package builtins

import (
	"fmt"
	"io"
	"os/exec"
)

func handleType(out io.Writer, args []string, builtins Registry) error {
	if len(args) < 1 {
		return fmt.Errorf("type: missing argument")
	}

	if len(args) > 1 {
		return fmt.Errorf("type: too many arguments")
	}

	cmd := args[0]

	if _, exists := builtins[cmd]; exists {
		_, err := fmt.Fprintf(out, "%s is a shell builtin\n", cmd)
		return err
	} else if path, _ := exec.LookPath(cmd); path != "" {
		_, err := fmt.Fprintf(out, "%s is %s\n", cmd, path)
		return err
	} else {
		_, err := fmt.Fprintf(out, "%s: not found\n", cmd)
		return err
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
