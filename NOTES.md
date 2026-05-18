Rough project structure

```plaintext
mysh/
├── main.go                  # Entry point, bootstraps shell
│
├── shell/
│   ├── shell.go             # Shell struct, REPL loop
│   └── config.go            # Shell config, env state
│
├── lexer/
│   └── lexer.go             # Tokenizer: raw string -> []Token
│
├── parser/
│   ├── parser.go            # Token stream -> AST
│   └── ast.go               # AST node types (Command, Pipeline, etc.)
│
├── executor/
│   ├── executor.go          # Walk AST, dispatch execution
│   ├── command.go           # Fork/exec external commands
│   ├── pipeline.go          # Connect commands via pipes
│   └── redirect.go          # File redirection logic
│
├── builtins/
│   ├── builtins.go          # Registry + dispatch for builtins
│   ├── cd.go
│   ├── echo.go
│   ├── export.go
│   └── exit.go
│
├── env/
│   └── env.go               # Environment variables, $VAR expansion
│
├── signals/
│   └── signals.go           # SIGINT, SIGTSTP, process groups
│
├── jobs/
│   └── jobs.go              # Job table, fg/bg/& support
│
└── readline/
    └── readline.go          # History, line editing wrapper
```
