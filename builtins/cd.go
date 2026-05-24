package builtins

import (
	"fmt"
	"os"
	"strings"
)

func handleCd(_ Registry, args []string) (string, error) {
	if len(args) > 1 {
		return "", fmt.Errorf("cd: too many arguments\n")
	}

	target := os.Getenv("HOME")
	if len(args) == 1 {
		target = strings.Replace(args[0], "~", os.Getenv("HOME"), 1)
	}

	if err := os.Chdir(target); err != nil {
		return "", fmt.Errorf("cd: %s: No such file or directory\n", target)
	}

	return "", nil
}

func tldrCd() string {
	// TODO
	return ""
}

func manCd() string {
	// TODO
	return ""
}
