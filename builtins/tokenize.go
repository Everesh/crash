package builtins

import (
	"fmt"
	"io"
)

func handleTokenize(out io.Writer, args []string) error {
	// I could pass in raw line, but since this is already tokenized w/e
	o := "[\n"
	for _, arg := range args {
		o += fmt.Sprintf("  \"%s\",\n", arg)
	}
	_, err := fmt.Fprintln(out, o+"]")
	return err
}

func tldrTokenize() string {
	// TODO
	return ""
}

func manTokenize() string {
	// TODO
	return ""
}
