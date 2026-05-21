package builtins

import (
	"fmt"
	"os"
)

func handlePwd(_ Registry, args []string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("pwd: failed to fetch working directory")
		return
	}

	fmt.Println(pwd)
}

func tldrPwd() string {
	// TODO
	return ""
}

func manPwd() string {
	// TODO
	return ""
}
