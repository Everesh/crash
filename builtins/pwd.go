package builtins

import (
	"fmt"
	"os"

	s "github.com/Everesh/crash/streams"
)

func handlePwd(io s.Io, args []string) {
	pwd, err := os.Getwd()
	if err != nil {
		io.WriteErr("pwd: failed to fetch working directory")
		return
	}
	if _, err = fmt.Fprintln(io.Out, pwd); err != nil {
		io.WriteErr("pwd: %s", err)
	}
}

func tldrPwd() string {
	// TODO
	return ""
}

func manPwd() string {
	// TODO
	return ""
}
