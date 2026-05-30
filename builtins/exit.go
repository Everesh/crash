package builtins

import (
	"fmt"
	"io"
	"strconv"
)

type ExitError struct{ Code int }

func (e *ExitError) Error() string { return fmt.Sprintf("exit %d", e.Code) }

func handleExit(_ io.Writer, args []string) error {
	if len(args) == 0 {
		return &ExitError{Code: 0}
	}
	if len(args) > 1 {
		return fmt.Errorf("exit: too many arguments")
	}

	code, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("exit: %s: invalid argument", args[0])
	}
	if code < 0 || code > 255 {
		return fmt.Errorf("exit: %d: out of range 0-255", code)
	}

	return &ExitError{Code: code}
}

func tldrExit() string {
	// TODO
	return ""
}

func manExit() string {
	// TODO
	return ""
}
