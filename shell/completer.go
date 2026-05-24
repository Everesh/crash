package shell

import (
	"os"
	"path/filepath"
	"strings"
)

type shellCompleter struct {
	commands []string
}

func completer(s *Shell) *shellCompleter {
	seen := make(map[string]bool)
	commands := make([]string, 0, len(s.builtins))
	for name := range s.builtins {
		seen[name] = true
		commands = append(commands, name)
	}
	for _, name := range scanPathBinaries() {
		if !seen[name] {
			seen[name] = true
			commands = append(commands, name)
		}
	}
	return &shellCompleter{commands: commands}
}

func scanPathBinaries() []string {
	seen := make(map[string]bool)
	var bins []string
	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() || seen[e.Name()] {
				continue
			}
			info, err := e.Info()
			if err != nil {
				continue
			}
			if info.Mode()&0111 != 0 {
				seen[e.Name()] = true
				bins = append(bins, e.Name())
			}
		}
	}
	return bins
}

func (c *shellCompleter) Do(line []rune, pos int) ([][]rune, int) {
	lineStr := string(line[:pos])
	endsWithSpace := len(lineStr) > 0 && lineStr[len(lineStr)-1] == ' '
	words := strings.Fields(lineStr)

	var prefix string
	firstWord := false

	switch {
	case len(words) == 0:
		firstWord = true
		prefix = ""
	case len(words) == 1 && !endsWithSpace:
		firstWord = true
		prefix = words[0]
	default:
		if endsWithSpace {
			prefix = ""
		} else {
			prefix = words[len(words)-1]
		}
	}

	if firstWord {
		return c.completeCommand(prefix)
	}
	return completeFilePath(prefix)
}

func (c *shellCompleter) completeCommand(prefix string) ([][]rune, int) {
	var result [][]rune
	for _, name := range c.commands {
		if strings.HasPrefix(name, prefix) {
			result = append(result, []rune(name[len(prefix):]))
		}
	}
	return result, len([]rune(prefix))
}

func expandTilde(path string) string {
	if path == "~" || strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return home + path[1:]
	}
	return path
}

func completeFilePath(prefix string) ([][]rune, int) {
	dir, file := filepath.Split(expandTilde(prefix))
	lookDir := dir
	if lookDir == "" {
		lookDir = "."
	}

	entries, err := os.ReadDir(lookDir)
	if err != nil {
		return nil, 0
	}

	var result [][]rune
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasPrefix(name, file) {
			continue
		}
		suffix := name[len(file):]
		if entry.IsDir() {
			suffix += "/"
		}
		result = append(result, []rune(suffix))
	}
	return result, len([]rune(file))
}
