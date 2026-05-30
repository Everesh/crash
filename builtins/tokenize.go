package builtins

import (
	"fmt"
	"io"
	"strings"

	t "github.com/Everesh/crash/parser/tokens"
)

func handleTokenize(out io.Writer, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("tokenize: missing argument(s)")
	}

	var b strings.Builder

	for _, str := range args {
		b.WriteString("[\n")

		tokens, err := t.Tokenize(str)
		if err != nil {
			return fmt.Errorf("tokenize: %s", err)
		}

		for _, tok := range tokens {
			if tok.Kind == t.Word {
				b.WriteString(fmt.Sprintf("  %v(\"%s\"),\n", tok.Kind, tok.Value))
			} else {
				b.WriteString(fmt.Sprintf("  %v\n", tok.Kind))
			}
		}

		b.WriteString("],\n")
	}

	fmt.Fprint(out, b.String())
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
