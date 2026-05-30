package builtins

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	f "github.com/Everesh/crash/parser/flags"
)

var specCommand = f.Spec{
	Flags: []f.Flag{
		{ // show path
			Short: 'v',
		},
		{ // show path prefixed by friendly description (mostly for batching)
			Short: 'V',
		},
		{ // searches hardcoded default UNIX path
			Short: 'p',
		},
	},
	Groups: []f.Group{
		{
			Flags:     []string{"v", "V"},
			Exclusive: true,
		},
	},
}

func handleCommand(out io.Writer, args []string) error {
	parsed, err := f.Parse(args, specCommand)
	if err != nil {
		return err
	}

	if len(parsed.Operands) < 1 {
		return fmt.Errorf("command: missing argument(s)\n")
	}

	cmd := parsed.Operands[0]

	var path string
	if parsed.Bool("p") {
		path = "/usr/bin:/bin:/usr/sbin:/sbin"
	} else {
		path = os.Getenv("PATH")
	}

	bin, err := lookInCustomPath(cmd, path)
	if err != nil {
		return fmt.Errorf("command: %s: no such command in path\n", cmd)
	}

	switch {
	case parsed.Bool("V"):
		fmt.Fprintf(out, "%s is %s\n", cmd, bin)

	case parsed.Bool("v"):
		fmt.Fprintf(out, "%s\n", bin)

	default:
		child := exec.Command(bin, parsed.Operands[1:]...)
		child.Stdin = os.Stdin
		child.Stdout = out
		child.Stderr = os.Stderr

		var cleanEnv []string
		for _, env := range os.Environ() {
			if !strings.HasPrefix(env, "PATH=") {
				cleanEnv = append(cleanEnv, env)
			}
		}
		cleanEnv = append(cleanEnv, "PATH="+path)
		child.Env = cleanEnv

		if err := child.Run(); err != nil {
			return fmt.Errorf("command: error running command: %w\n", err)
		}
	}

	return nil
}

func tldrCommand() string {
	// TODO
	return ""
}

func manCommand() string {
	// TODO
	return ""
}

func lookInCustomPath(file string, customPath string) (string, error) {
	if filepath.Base(file) != file {
		return file, nil
	}

	dirs := filepath.SplitList(customPath)
	for _, dir := range dirs {
		// guard for working dir reference in path (either 2 colons back to back or a trailing one)
		if dir == "" {
			dir = "."
		}
		path := filepath.Join(dir, file)

		if d, err := os.Stat(path); err == nil {
			if m := d.Mode(); !m.IsDir() {
				return path, nil
			}
		}
	}

	return "", &exec.Error{Name: file, Err: exec.ErrNotFound}
}
