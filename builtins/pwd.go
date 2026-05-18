package builtins

import (
	"fmt"
	"os"
)

func handlePwd(args []string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error fetching working directory")
	}

	fmt.Println(pwd)
}
