package builtins

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// raw os.Exit leaves readline unclosed, but every solution I came up with was
// just a bunch of external state which eww or overcomplicating the handler api
// for now Im just going to let the kernel restore the terminal attributes
func handleExit(_ io.Writer, args []string) error {
	if len(args) == 0 {
		os.Exit(0)
	}
	if len(args) > 1 {
		return fmt.Errorf("exit: too many arguments\n")
	}

	code, err := strconv.Atoi(args[0])

	if err != nil {
		return fmt.Errorf("exit: %s: invalid argument\n", args[0])
	}

	if code < 0 || code > 255 {
		return fmt.Errorf("exit: %d: out of range 0-255\n", code)
	}

	os.Exit(code)
	return fmt.Errorf("exit: os.Exit(%d) failed: this should be impossible O.o", code)
}

func tldrExit() string {
	// TODO
	return ""
}

func manExit() string {
	// TODO
	return ""
}
