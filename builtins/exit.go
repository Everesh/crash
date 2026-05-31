package builtins

import (
	"strconv"

	s "github.com/Everesh/crash/streams"
	b "github.com/Everesh/crash/streams/bus"
)

func handleExit(io s.Io, args []string) {
	if len(args) == 0 {
		io.Send(b.ExitCmd{Code: 0})
		return
	}
	if len(args) > 1 {
		io.WriteErr("exit: too many arguments")
		return
	}

	code, err := strconv.Atoi(args[0])
	if err != nil {
		io.WriteErr("exit: %s: invalid argument", args[0])
		return
	}
	if code < 0 || code > 255 {
		io.WriteErr("exit: %d: out of range 0-255", code)
		return
	}

	io.Send(b.ExitCmd{Code: code})
}

func tldrExit() string {
	// TODO
	return ""
}

func manExit() string {
	// TODO
	return ""
}
