package tokens

import (
	"fmt"
	"slices"
	"strings"
)

var spaceRunes = []rune{' ', '\t', '\n', '\r'}
var operatorRunes = []rune{'>', '<', '|', ';', '&'}

func Tokenize(s string) ([]Token, error) {
	l := newLexer(s)
	var tokens []Token

	for {
		r, ok := l.next()
		if !ok {
			break
		}
		if isSpace(r) {
			continue
		}

		if r >= '0' && r <= '9' {
			if next, ok := l.peek(); ok && next == '>' {
				tok, err := readOperator(l, r)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, tok)
				continue
			}
		}

		if isOperatorRune(r) {
			tok, err := readOperator(l, r)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, tok)
			continue
		}

		word, err := readWord(l, r)
		if err != nil {
			return nil, err
		}
		if word != "" {
			tokens = append(tokens, Token{Kind: Word, Value: word})
		}
	}

	return expand(tokens), nil
}

func readOperator(l *lexer, first rune) (Token, error) {
	switch first {
	case '>':
		if r, ok := l.peek(); ok && r == '>' {
			l.next()
			return Token{Kind: RedirectAppend}, nil
		}
		return Token{Kind: RedirectOut}, nil

	case '<':
		return Token{Kind: RedirectIn}, nil

	case '|':
		if r, ok := l.peek(); ok && r == '|' {
			l.next()
			return Token{Kind: Or}, nil
		}
		return Token{Kind: Pipe}, nil

	case ';':
		return Token{Kind: Semicolon}, nil

	case '&':
		if r, ok := l.peek(); ok && r == '>' {
			l.next()
			return Token{Kind: RedirectBoth}, nil
		}
		if r, ok := l.peek(); ok && r == '&' {
			l.next()
			return Token{Kind: And}, nil
		}
		// not checking for trailing space, `echo this &echo that` is valid POSIX syntax
		return Token{}, fmt.Errorf("bare & not yet supported") // TODO

	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		if r, ok := l.next(); ok && r != '>' { // the > should get consumed here, if it doesn't match we throw anyway
			return Token{}, fmt.Errorf("readOperator: unexpected suffix %q for rune %q", r, first)
		}
		if r, ok := l.peek(); ok && r == '>' {
			l.next()
			return Token{Kind: RedirectErrAppend}, nil
		}
		return Token{Kind: RedirectErr}, nil

	default:
		return Token{}, fmt.Errorf("readOperator: unexpected rune %q", first)
	}
}

func readWord(l *lexer, first rune) (string, error) {
	var b strings.Builder
	r := first

	for {
		switch {
		case isSpace(r):
			return b.String(), nil
		case isOperatorRune(r):
			l.backup()
			return b.String(), nil
		case r == '\\':
			next, ok := l.next()
			if !ok {
				return "", fmt.Errorf("readWord: unexpected EOF after backslash")
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
			return fmt.Errorf("readSingleQuote: unterminated single quote")
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
			return fmt.Errorf("readDoubleQuote: unterminated double quote")
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
	return slices.Contains(spaceRunes, r)
}

func isOperatorRune(r rune) bool {
	return slices.Contains(operatorRunes, r)
}
