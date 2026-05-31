package builtins

import (
	"fmt"
	"strings"

	s "github.com/Everesh/crash/streams"
)

func handleEcho(io s.Io, args []string) {
	if _, err := fmt.Fprintln(io.Out, strings.Join(args, " ")); err != nil {
		io.WriteErr("echo: %s", err)
	}
}

func tldrEcho() string {
	// TODO
	return ""
}

func manEcho() string {
	// TODO
	return ""
}
