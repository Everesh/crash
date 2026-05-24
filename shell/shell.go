package shell

import (
	"github.com/Everesh/crash/builtins"
)

type Shell struct {
	builtins builtins.Registry
}

func New() *Shell {
	return &Shell{builtins: builtins.NewRegistry()}
}
