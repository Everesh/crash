package parser

import (
	"os"
	"strings"
)

func expand(tokens []Token) []Token {
	out := make([]Token, len(tokens))
	for i, tok := range tokens {
		if tok.Kind == Word {
			out[i] = Token{Kind: Word, Value: expandTilde(tok.Value)}
		} else {
			out[i] = tok
		}
	}
	return out
}

func expandTilde(s string) string {
	if s == "~" || strings.HasPrefix(s, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return s
		}
		return home + s[1:]
	}
	return s
}
