package builtins

import (
	"strings"
)

func handleEcho(_ Registry, args []string) (string, error) {
	return strings.Join(args, " ") + "\n", nil
}

func tldrEcho() string {
	// TODO
	return ""
}

func manEcho() string {
	// TODO
	return ""
}
