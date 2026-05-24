package builtins

import (
	"fmt"
	"os"
)

func handlePwd(_ Registry, args []string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("pwd: failed to fetch working directory\n")
	}

	return pwd + "\n", nil
}

func tldrPwd() string {
	// TODO
	return ""
}

func manPwd() string {
	// TODO
	return ""
}
