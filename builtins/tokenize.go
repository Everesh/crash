package builtins

import (
	"fmt"
	"io"
	"strings"
)

func handleTokenize(out io.Writer, args []string) error {
	// I could pass in raw line, but since this is already tokenized w/e
	var b strings.Builder
	b.WriteString("[\n")
	for _, arg := range args {
		fmt.Fprintf(&b, "  \"%s\",\n", arg)
	}
	b.WriteString("]")
	_, err := fmt.Fprintln(out, b.String())
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
