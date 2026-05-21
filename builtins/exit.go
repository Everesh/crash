package builtins

import (
	"fmt"
	"os"
	"strconv"
)

func handleExit(_ Registry, args []string) {
	if len(args) == 0 {
		os.Exit(0)
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "exit: too many arguments")
		return
	}

	code, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "exit: %s: invalid argument\n", args[0])
		return
	}

	if code < 0 || code > 255 {
		fmt.Fprintf(os.Stderr, "exit: %d: out of range 0-255\n", code)
		return
	}

	os.Exit(code)
}

func tldrExit() string {
	// TODO
	return ""
}

func manExit() string {
	// TODO
	return ""
}
