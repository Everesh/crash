package config

type LexerConfStruct struct {
	Escape   []rune // escape runes e.g.: \
	Glob     []rune // grouping runes e.g.: '
	EvalGlob []rune // grouping runes that eval their content e.g.: "
	Space    []rune // whitespace runes e.g.: \n \t space
}

var LexerConf = LexerConfStruct{
	Escape:   []rune{'\\'},
	Glob:     []rune{'\''},
	EvalGlob: []rune{'"'},
	Space:    []rune{' ', '\n', '\t', '\v', '\f', '\r', 0x85, 0xA0},
}
