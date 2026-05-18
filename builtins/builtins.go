package builtins

type CommandHandler func(args []string)

var Registry = make(map[string]CommandHandler)

func init() {
	Registry["exit"] = handleExit
	Registry["echo"] = handleEcho
	Registry["type"] = handleType
	Registry["pwd"] = handlePwd
	Registry["cd"] = handleCd
}
