package builtins

type CommandHandler struct {
	Handle func(args []string)
	Tldr   func() string // tldr.sh like help
	Man    func() string // man like help
}

var Registry = make(map[string]CommandHandler)

func init() {
	Registry["exit"] = CommandHandler{Handle: handleExit, Tldr: tldrExit, Man: manExit}
	Registry["echo"] = CommandHandler{Handle: handleEcho, Tldr: tldrEcho, Man: manEcho}
	Registry["type"] = CommandHandler{Handle: handleType, Tldr: tldrType, Man: manType}
	Registry["pwd"] = CommandHandler{Handle: handlePwd, Tldr: tldrPwd, Man: manPwd}
	Registry["cd"] = CommandHandler{Handle: handleCd, Tldr: tldrCd, Man: manCd}
}
