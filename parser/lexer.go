package parser

type lexer struct {
	runes []rune
	pos   int
}

func newLexer(s string) *lexer {
	return &lexer{runes: []rune(s)}
}

func (l *lexer) next() (rune, bool) {
	if l.pos >= len(l.runes) {
		return 0, false
	}
	r := l.runes[l.pos]
	l.pos++
	return r, true
}

func (l *lexer) peek() (rune, bool) {
	if l.pos >= len(l.runes) {
		return 0, false
	}
	return l.runes[l.pos], true
}
