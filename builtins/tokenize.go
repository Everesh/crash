package builtins

import (
	"fmt"
)

func handleTokenize(_ Registry, args []string) {
	// I could pass in raw line, but since this is already tokenized w/e
	fmt.Println("[")
	for _, arg := range args {
		fmt.Printf("  \"%s\",\n", arg)
	}
	fmt.Println("]")

}

func tldrTokenize() string {
	// TODO
	return ""
}

func manTokenize() string {
	// TODO
	return ""
}
