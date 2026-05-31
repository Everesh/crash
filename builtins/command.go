package builtins

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	f "github.com/Everesh/crash/parser/flags"
	s "github.com/Everesh/crash/streams"
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

func handleCommand(io s.Io, args []string) {
	parsed, err := f.Parse(args, specCommand)
	if err != nil {
		io.WriteErr("%s", err)
		return
	}

	if len(parsed.Operands) < 1 {
		io.WriteErr("command: missing argument(s)")
		return
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
		io.WriteErr("command: %s: no such command in path", cmd)
		return
	}

	switch {
	case parsed.Bool("V"):
		fmt.Fprintf(io.Out, "%s is %s\n", cmd, bin)

	case parsed.Bool("v"):
		fmt.Fprintf(io.Out, "%s\n", bin)

	default:
		child := exec.Command(bin, parsed.Operands[1:]...)
		child.Stdin = io.In
		child.Stdout = io.Out
		child.Stderr = io.Err

		var cleanEnv []string
		for _, env := range os.Environ() {
			if !strings.HasPrefix(env, "PATH=") {
				cleanEnv = append(cleanEnv, env)
			}
		}
		cleanEnv = append(cleanEnv, "PATH="+path)
		child.Env = cleanEnv

		if err := child.Run(); err != nil {
			var exitErr *exec.ExitError
			if !errors.As(err, &exitErr) {
				io.WriteErr("command: %s", err)
			}
		}
	}
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
