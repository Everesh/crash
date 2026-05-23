package lexer

import (
	"fmt"

	"github.com/Everesh/crash/config"
)

func Parse(str string) []string {
	chunks := make([]string, 0)
	chunk := ""

	escaped := false
	skipTo := ' '

	for _, v := range str {
		if escaped {
			chunk += string(v)
			escaped = false
			continue
		}

		switch v {
		case skipTo:
			if skipTo == '"' || skipTo == '\'' {
				skipTo = ' '
				continue
			}

			if chunk != "" {
				chunks = append(chunks, chunk)
				chunk = ""
			}

		case '\\':
			if skipTo == ' ' {
				escaped = true
			} else {
				chunk += string(v)
			}

		case '"', '\'':
			if skipTo == ' ' {
				skipTo = v
			} else {
				chunk += string(v)
			}

		case '\n':
			continue

		default:
			chunk += string(v)
		}
	}

	if chunk != "" {
		chunks = append(chunks, chunk)
	}

	if skipTo != ' ' {
		fmt.Printf("%s: lexer: unterminated quoted string\n", config.AppName)
		return []string{}
	}

	return chunks
}
