package shell

import (
	"github.com/Everesh/crash/builtins"
	b "github.com/Everesh/crash/streams/bus"
)

type Shell struct {
	builtins builtins.Registry
	signals  *b.Queue
}

func New() *Shell {
	return &Shell{
		builtins: builtins.NewRegistry(),
		signals:  b.NewQueue(),
	}
}
