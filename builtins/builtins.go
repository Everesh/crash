package builtins

import s "github.com/Everesh/crash/streams"

type Builtin struct {
	Handle func(s.Io, []string)
	Tldr   func() string // tldr.sh like help
	Man    func() string // man like help
}

type Registry map[string]Builtin

func NewRegistry() Registry {
	r := make(Registry)

	r["exit"] = Builtin{Handle: handleExit, Tldr: tldrExit, Man: manExit}
	r["echo"] = Builtin{Handle: handleEcho, Tldr: tldrEcho, Man: manEcho}
	r["pwd"] = Builtin{Handle: handlePwd, Tldr: tldrPwd, Man: manPwd}
	r["cd"] = Builtin{Handle: handleCd, Tldr: tldrCd, Man: manCd}
	r["command"] = Builtin{Handle: handleCommand, Tldr: tldrCommand, Man: manCommand}
	r["tokenize"] = Builtin{Handle: handleTokenize, Tldr: tldrTokenize, Man: manTokenize}

	// registry aware
	r["type"] = Builtin{
		Handle: func(io s.Io, args []string) { handleType(io, args, r) },
		Tldr:   tldrType,
		Man:    manType}

	return r
}
