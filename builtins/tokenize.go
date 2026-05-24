package builtins

import (
	"fmt"
)

func handleTokenize(_ Registry, args []string) (string, error) {
	// I could pass in raw line, but since this is already tokenized w/e
	o := "[\n"
	for _, arg := range args {
		o += fmt.Sprintf("  \"%s\",\n", arg)
	}

	return o + "]\n", nil
}

func tldrTokenize() string {
	// TODO
	return ""
}

func manTokenize() string {
	// TODO
	return ""
}
