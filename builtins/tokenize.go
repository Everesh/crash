package builtins

import (
	"fmt"
	"io"
	"strings"

	"github.com/Everesh/crash/parser"
)

func handleTokenize(out io.Writer, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("tokenize: missing argument(s)\n")
	}

	var b strings.Builder

	for _, str := range args {
		b.WriteString("[\n")

		tokens, err := parser.Tokenize(str)
		if err != nil {
			return err
		}

		for _, tok := range tokens {
			if tok.Kind == parser.Word {
				b.WriteString(fmt.Sprintf("  %v(\"%s\"),\n", tok.Kind, tok.Value))
			} else {
				b.WriteString(fmt.Sprintf("  %v\n", tok.Kind))
			}
		}

		b.WriteString("],\n")
	}

	fmt.Print(b.String())
	return nil
}

func tldrTokenize() string {
	// TODO
	return ""
}

func manTokenize() string {
	// TODO
	return ""
}
