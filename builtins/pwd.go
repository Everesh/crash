package builtins

import (
	"fmt"
	"io"
	"os"
)

func handlePwd(out io.Writer, args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("pwd: failed to fetch working directory")
	}

	_, err = fmt.Fprintln(out, pwd)
	return err
}

func tldrPwd() string {
	// TODO
	return ""
}

func manPwd() string {
	// TODO
	return ""
}
