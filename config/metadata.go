package config

import "fmt"

const AppName = "crash"
const Version = "0.0.0"

var PS1 = fmt.Sprintf("\x1b[31;1m%s\x1b[0m:\x1b[33m%s\x1b[0m$ ", AppName, Version)
