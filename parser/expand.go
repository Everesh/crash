package parser

import (
	"os"
	"strings"
)

func expand(tokens []string) []string {
	out := make([]string, len(tokens))
	for i, tok := range tokens {
		out[i] = expandTilde(tok)
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
