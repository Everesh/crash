package config

import "fmt"

var PS1 = fmt.Sprintf("\x1b[31;1m%s\x1b[0m:\x1b[33m%s\x1b[0m$ ", AppName, Version)
var PS2 = "\x1b[31;1m|\x1b[0m "
var VimMode = false
