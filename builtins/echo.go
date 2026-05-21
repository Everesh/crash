package builtins

import (
	"fmt"
	"strings"
)

func handleEcho(_ Registry, args []string) {
	fmt.Println(strings.Join(args, " "))
}

func tldrEcho() string {
	// TODO
	return ""
}

func manEcho() string {
	// TODO
	return ""
}
