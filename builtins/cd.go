package builtins

import (
	"os"
	"strings"

	s "github.com/Everesh/crash/streams"
)

func handleCd(io s.Io, args []string) {
	if len(args) > 1 {
		io.WriteErr("cd: too many arguments")
		return
	}

	target := os.Getenv("HOME")
	if len(args) == 1 {
		target = strings.Replace(args[0], "~", os.Getenv("HOME"), 1)
	}

	if err := os.Chdir(target); err != nil {
		io.WriteErr("cd: %s: No such file or directory", target)
	}
}

func tldrCd() string {
	// TODO
	return ""
}

func manCd() string {
	// TODO
	return ""
}
