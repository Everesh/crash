package shell

import (
	"fmt"
	"os"
	"os/exec"

	t "github.com/Everesh/crash/parser/tokens"
	s "github.com/Everesh/crash/streams"
)

func (sh *Shell) Eval(input string) {
	toks, err := t.Tokenize(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse error:", err)
		return
	}
	if len(toks) == 0 {
		return
	}

	words, redirs, err := parseTokens(toks)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if len(words) == 0 {
		return
	}

	io, files, err := applyRedirects(redirs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer func() {
		for _, f := range files {
			f.Close()
		}
	}()

	cmd, args := words[0], words[1:]

	if builtin, exists := sh.builtins[cmd]; exists {
		builtin.Handle(io, args)
	} else if _, err := exec.LookPath(cmd); err != nil {
		io.WriteErr("%s: command not found", cmd)
	} else {
		child := exec.Command(cmd, args...)
		child.Stdin = io.In
		child.Stdout = io.Out
		child.Stderr = io.Err
		child.Run()
	}

	for _, cmd := range io.Drain() {
		sh.signals.Send(cmd)
	}
}

type redirSpec struct {
	kind t.TokenKind
	path string
}

func parseTokens(toks []t.Token) ([]string, []redirSpec, error) {
	var words []string
	var redirs []redirSpec

	for i := 0; i < len(toks); i++ {
		tok := toks[i]
		switch tok.Kind {
		case t.Word:
			words = append(words, tok.Value)
		case t.RedirectIn, t.RedirectOut, t.RedirectAppend,
			t.RedirectErr, t.RedirectErrAppend, t.RedirectBoth:
			i++
			if i >= len(toks) || toks[i].Kind != t.Word {
				return nil, nil, fmt.Errorf("parseTokens: expected filename after redirect")
			}
			redirs = append(redirs, redirSpec{tok.Kind, toks[i].Value})
		default:
			return nil, nil, fmt.Errorf("parseTokens: %v: operator not yet supported", tok.Kind)
		}
	}

	return words, redirs, nil
}

func applyRedirects(redirs []redirSpec) (s.Io, []*os.File, error) {
	io := s.NewIo(os.Stdin, os.Stdout, os.Stderr)
	var files []*os.File

	for _, r := range redirs {
		switch r.kind {
		case t.RedirectOut:
			f, err := os.Create(r.path)
			if err != nil {
				return s.Io{}, nil, err
			}
			files = append(files, f)
			io.Out = f
		case t.RedirectAppend:
			f, err := os.OpenFile(r.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return s.Io{}, nil, err
			}
			files = append(files, f)
			io.Out = f
		case t.RedirectIn:
			f, err := os.Open(r.path)
			if err != nil {
				return s.Io{}, nil, err
			}
			files = append(files, f)
			io.In = f
		case t.RedirectErr:
			f, err := os.Create(r.path)
			if err != nil {
				return s.Io{}, nil, err
			}
			files = append(files, f)
			io.Err = f
		case t.RedirectErrAppend:
			f, err := os.OpenFile(r.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return s.Io{}, nil, err
			}
			files = append(files, f)
			io.Err = f
		case t.RedirectBoth:
			f, err := os.Create(r.path)
			if err != nil {
				return s.Io{}, nil, err
			}
			files = append(files, f)
			io.Out = f
			io.Err = f
		}
	}

	return io, files, nil
}
