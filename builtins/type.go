package builtins

import (
	"fmt"
	"os/exec"

	s "github.com/Everesh/crash/streams"
)

func handleType(io s.Io, args []string, builtins Registry) {
	if len(args) < 1 {
		io.WriteErr("type: missing argument")
		return
	}
	if len(args) > 1 {
		io.WriteErr("type: too many arguments")
		return
	}

	cmd := args[0]

	if _, exists := builtins[cmd]; exists {
		fmt.Fprintf(io.Out, "%s is a shell builtin\n", cmd)
	} else if path, _ := exec.LookPath(cmd); path != "" {
		fmt.Fprintf(io.Out, "%s is %s\n", cmd, path)
	} else {
		fmt.Fprintf(io.Out, "%s: not found\n", cmd)
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
