package parser

import (
	"fmt"
	"slices"

	"github.com/Everesh/crash/config"
)

func Tokenize(str string) ([]string, error) {
	chunks := make([]string, 0)
	chunk := ""
	lexer := NewLexer(str)

	for {
		rune, ok, err := lexer.Next(true)
		if err != nil {
			return nil, err
		}
		if !ok {
			break // end of lexer
		}

		switch {
		case slices.Contains(config.LexerConf.Space, rune):
			if chunk != "" {
				chunks = append(chunks, chunk)
				chunk = ""
			}
		case slices.Contains(config.LexerConf.Glob, rune):
			innerChunk, err := glob(lexer, rune)
			if err != nil {
				return nil, err
			}
			chunk += innerChunk
		case slices.Contains(config.LexerConf.EvalGlob, rune):
			innerChunk, err := evalGlob(lexer, rune)
			if err != nil {
				return nil, err
			}
			chunk += innerChunk
		default:
			chunk += string(rune)
		}
	}

	if chunk != "" {
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

func glob(lexer *Lexer, delimeter rune) (string, error) {
	chunk := ""
	for {
		rune, ok, err := lexer.Next(false)
		if err != nil {
			return "", err
		}
		if !ok {
			return "", fmt.Errorf(
				"lexer: tokenizer: glob: missing closure for %c",
				delimeter) // end of lexer
		}
		if rune == delimeter {
			break
		}

		chunk += string(rune)
	}

	return chunk, nil
}

func evalGlob(lexer *Lexer, delimeter rune) (string, error) {
	chunk := ""
	for {
		rune, ok, err := lexer.Next(false)
		if err != nil {
			return "", err
		}
		if !ok {
			return "", fmt.Errorf(
				"lexer: tokenizer: evalGlob: missing closure for %c",
				delimeter) // end of lexer
		}
		if rune == delimeter {
			break
		}

		chunk += string(rune)
	}

	return chunk, nil
}
