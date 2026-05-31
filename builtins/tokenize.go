package builtins

import (
	"fmt"
	"strings"

	t "github.com/Everesh/crash/parser/tokens"
	s "github.com/Everesh/crash/streams"
)

func handleTokenize(io s.Io, args []string) {
	if len(args) < 1 {
		io.WriteErr("tokenize: missing argument(s)")
		return
	}

	var b strings.Builder

	for _, str := range args {
		b.WriteString("[\n")

		tokens, err := t.Tokenize(str)
		if err != nil {
			io.WriteErr("tokenize: %s", err)
			return
		}

		for _, tok := range tokens {
			if tok.Kind == t.Word {
				fmt.Fprintf(&b, "  %v(\"%s\"),\n", tok.Kind, tok.Value)
			} else {
				fmt.Fprintf(&b, "  %v\n", tok.Kind)
			}
		}

		b.WriteString("],\n")
	}

	fmt.Fprint(io.Out, b.String())
}

func tldrTokenize() string {
	// TODO
	return ""
}

func manTokenize() string {
	// TODO
	return ""
}
