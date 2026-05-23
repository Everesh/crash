package builtins

type Builtin struct {
	Handle func(registry Registry, args []string)
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
	r["type"] = Builtin{Handle: handleType, Tldr: tldrType, Man: manType}
	r["command"] = Builtin{Handle: handleCommand, Tldr: tldrCommand, Man: manCommand}
	r["tokenize"] = Builtin{Handle: handleTokenize, Tldr: tldrTokenize, Man: manTokenize}
	return r
}
