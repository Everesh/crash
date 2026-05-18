package builtins

import "os"

func handleExit(args []string) {
	os.Exit(0)
}
