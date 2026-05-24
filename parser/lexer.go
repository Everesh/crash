package parser

import (
	"fmt"
	"slices"

	"github.com/Everesh/crash/config"
)

type Lexer struct {
	runes []rune
	pos   int
}

func NewLexer(str string) *Lexer {
	return &Lexer{runes: []rune(str), pos: 0}
}

func (lexer *Lexer) Next(escape bool) (rune rune, ok bool, err error) {
	if lexer.pos >= len(lexer.runes) {
		return 0, false, nil
	}

	rune = lexer.runes[lexer.pos]
	lexer.pos++

	if escape && slices.Contains(config.LexerConf.Escape, rune) {
		if lexer.pos >= len(lexer.runes) {
			return 0, false, fmt.Errorf(
				"lexer: next: tailing unescaped escape char %c",
				rune)
		}

		rune = lexer.runes[lexer.pos]
		lexer.pos++
	}

	return rune, true, nil
}

func (lexer *Lexer) Peek() (rune, bool) {
	if lexer.pos >= len(lexer.runes) {
		return 0, false
	}
	return lexer.runes[lexer.pos], true
}
