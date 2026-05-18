package builtins

import (
	"fmt"
	"os"
	"strings"
)

func handleCd(args []string) {
	if len(args) > 1 {
		fmt.Println("cd: invalid amount of arguments")
		return
	}

	target := os.Getenv("HOME")
	if len(args) == 1 {
		target = strings.Replace(args[0], "~", os.Getenv("HOME"), 1)
	}

	if err := os.Chdir(target); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", target)
	}
}
