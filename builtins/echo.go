package builtins

import (
	"fmt"
	"io"
	"strings"
)

func handleEcho(out io.Writer, args []string) error {
	_, err := fmt.Fprintln(out, strings.Join(args, " "))
	return err
}

func tldrEcho() string {
	// TODO
	return ""
}

func manEcho() string {
	// TODO
	return ""
}
