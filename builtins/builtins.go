package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CommandHandler func(args []string)

var Registry = make(map[string]CommandHandler)

func init() {
	Registry["exit"] = handleExit
	Registry["echo"] = handleEcho
	Registry["type"] = handleType
	Registry["pwd"] = handlePwd
	Registry["cd"] = handleCd
}

func handleExit(args []string) {
	os.Exit(0)
}

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func handleType(args []string) {
	if len(args) == 0 {
		fmt.Println("type: missing argument")
		return
	}

	cmd := args[0]

	if _, exists := Registry[cmd]; exists {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else if path, _ := exec.LookPath(cmd); path != "" {
		fmt.Printf("%s is %s\n", cmd, path)
	} else {
		fmt.Printf("%s: not found\n", cmd)
	}
}

func handlePwd(args []string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error fetching working directory")
	}

	fmt.Println(pwd)
}

func handleCd(args []string) {
	if len(args) > 1 {
		fmt.Println("cd: invalid amount of arguments")
		return
	}

	target := os.Getenv("HOME")
	if len(args) == 1 {
		target = strings.Replace(args[0], "~", os.Getenv("HOME"), 1)
	}

	if err := os.Chdir(target); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", target)
	}
}
