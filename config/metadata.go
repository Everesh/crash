package config

import "fmt"

const AppName = "crash"
const Version = "0.0.0"

var PS1 = fmt.Sprintf("%s:%s$ ", AppName, Version)
