package builtins

import (
	"fmt"
	"strings"
)

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}
