package parser

import (
	"fmt"
	"strings"
)

func Tokenize(s string) ([]string, error) {
	l := newLexer(s)
	var tokens []string

	for {
		r, ok := l.next()
		if !ok {
			break
		}
		if isSpace(r) {
			continue
		}

		word, err := readWord(l, r)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, word)
	}

	return tokens, nil
}

func readWord(l *lexer, first rune) (string, error) {
	var b strings.Builder
	r := first

	for {
		switch {
		case isSpace(r):
			return b.String(), nil
		case r == '\\':
			next, ok := l.next()
			if !ok {
				return "", fmt.Errorf("unexpected EOF after backslash")
			}
			b.WriteRune(next)
		case r == '\'':
			if err := readSingleQuote(l, &b); err != nil {
				return "", err
			}
		case r == '"':
			if err := readDoubleQuote(l, &b); err != nil {
				return "", err
			}
		default:
			b.WriteRune(r)
		}

		next, ok := l.next()
		if !ok {
			return b.String(), nil
		}
		r = next
	}
}

func readSingleQuote(l *lexer, b *strings.Builder) error {
	for {
		r, ok := l.next()
		if !ok {
			return fmt.Errorf("unterminated single quote")
		}
		if r == '\'' {
			return nil
		}
		b.WriteRune(r)
	}
}

func readDoubleQuote(l *lexer, b *strings.Builder) error {
	for {
		r, ok := l.next()
		if !ok {
			return fmt.Errorf("unterminated double quote")
		}

		switch r {
		case '"':
			return nil
		// Only escapes $, `, ", \, and \n
		// see: www.gnu.org/software/bash/manual/html_node/Double-Quotes.html
		case '\\':
			next, ok := l.next()
			if !ok {
				return fmt.Errorf("unterminated double quote")
			}
			switch next {
			case '$', '`', '"', '\\', '\n':
				b.WriteRune(next)
			default:
				b.WriteRune('\\')
				b.WriteRune(next)
			}
		default:
			b.WriteRune(r)
		}
	}
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}
