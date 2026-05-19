package builtins

import (
	"fmt"
	"os/exec"
)

func handleType(args []string) {
	if len(args) == 0 {
		fmt.Println("type: missing argument")
		return
	}

	cmd := args[0]

	if _, exists := Registry[cmd]; exists {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else if path, _ := exec.LookPath(cmd); path != "" {
		fmt.Printf("%s is %s\n", cmd, path)
	} else {
		fmt.Printf("%s: not found\n", cmd)
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
