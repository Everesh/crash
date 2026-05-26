package parser

type TokenKind int

const (
	Word              TokenKind = iota
	Pipe                        // |
	RedirectIn                  // <
	RedirectOut                 // >
	RedirectAppend              // >>
	RedirectErr                 // 2>
	RedirectErrAppend           // 2>>
	RedirectBoth                // &>
	Semicolon                   // ;
	And                         // &&
	Or                          // ||
)

type Token struct {
	Kind  TokenKind
	Value string
}

func (k TokenKind) String() string {
	switch k {
	case Word:
		return "Word"
	case Pipe:
		return "Pipe"
	case RedirectIn:
		return "RedirectIn"
	case RedirectOut:
		return "RedirectOut"
	case RedirectAppend:
		return "RedirectAppend"
	case RedirectErr:
		return "RedirectErr"
	case RedirectErrAppend:
		return "RedirectErrAppend"
	case RedirectBoth:
		return "RedirectBoth"
	case Semicolon:
		return "Semicolon"
	case And:
		return "And"
	case Or:
		return "Or"
	default:
		return "Unknown"
	}
}
