package builtins

import "os"

func handleExit(_ Registry, args []string) {
	os.Exit(0)
}

func tldrExit() string {
	// TODO
	return ""
}

func manExit() string {
	// TODO
	return ""
}
