# crash
POSIX-adjacent shell, heavy WIP

![preview gif](https://github.com/user-attachments/assets/e84ba3de-0614-4577-8475-793907e1ef8d)

## Installation

1. Get yourself the [`Go` toolchain](https://go.dev/)
2. Clone and build:
```sh
git clone https://github.com/Everesh/crash
cd crash
go build -o crash .
```

## Running

```sh
./crash
```

> [!TIP]
you can set it as your shell of choice in terminal emulators (I wouldn't tho) e.g.
```toml
# ~/.config/alacritty/alacritty.toml
[terminal]
shell = "/path/to/the/bin"
```

---

# TODO
## MVP

- [x] Basic REPL
- [x] Basic Builtins
  - [x] Registry of builtins with sane api surface
  - [x] flags parser / splitter (tokenizer should not do this)
  - [x] exit
  - [x] echo
  - [x] pwd
  - [x] cd
  - [x] type
  - [x] command
    - [x] `-v` arg <!-- needed for silent bin discovery, I don't want to rely on `which` too hard here -->
  - [ ] alias
- [x] Lexer 
  - [x] Basic Quoting
  - [ ] Builtins for Help w/fallback to coresponding binaries
    - [ ] `?` tldr like
    - [ ] `???` man like
- [x] Tokenizer 
- [x] Basic Stdio Redirection
- [x] Basic Completion
  - [x] Command Completion
  - [x] Filename Completion
  - [ ] Programable Completion
- [ ] Basic Background Jobs
  - [ ] orphaned via & suffix
  - [ ] proper bg / fg
- [ ] Basic Pipelines
- [x] Basic History <-- readline gave that for free, FREE I tell you
  - [ ] History Persistence
    - [ ] XDG_DATA_HOME
- [ ] Basic Parameter Expansion
- [ ] Small Ansii lib for text formatting

## Pre-POSIX

- [ ] drop readline and just work with x/term, `chzyer/readline` is a hit or miss, some things are incredible, others drive me crazy
  - [ ] on the fly syntax highlighting (this should be doable if I own the rl state, I could tokenize on keypress and instead of passing it to eval I could just pass it to formater, save it as rl state and redraw... might be an issue stripping the formatting once applied, maybe a 2 tier system where one []rune slice holds the raw state, the other holds the displayed state? idk, future me will figure this out, right? RIGHT?!)
- [ ] config file (this could be like bashrc, just run the script on shell.New, dont reuse the XDG_DATA_HOME, this should not be .local but .config)
- [ ] some more builtins
  - [ ] alias
  - [ ] export

## POSIX
See: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/V3_chap02.html

- [ ] 2.2 Quoting
  - [ ] 2.2.1 Escape Character (Backslash)
  - [ ] 2.2.2 Single-Quotes
  - [ ] 2.2.3 Double-Quotes
  - [ ] 2.2.4 Dollar-Single-Quotes
- [ ] 2.3 Token Recognition
  - [ ] 2.3.1 Alias Substitution
- [ ] 2.4 Reserved Words
- [ ] 2.5 Parameters and Variables
  - [ ] 2.5.1 Positional Parameters
  - [ ] 2.5.2 Special Parameters
  - [ ] 2.5.3 Shell Variables
- [ ] 2.6 Word Expansions
  - [ ] 2.6.1 Tilde Expansion
  - [ ] 2.6.2 Parameter Expansion
  - [ ] 2.6.3 Command Substitution
  - [ ] 2.6.4 Arithmetic Expansion
  - [ ] 2.6.5 Field Splitting
  - [ ] 2.6.6 Pathname Expansion
  - [ ] 2.6.7 Quote Removal
- [ ] 2.7 Redirection
  - [ ] 2.7.1 Redirecting Input
  - [ ] 2.7.2 Redirecting Output
  - [ ] 2.7.3 Appending Redirected Output
  - [ ] 2.7.4 Here-Document
  - [ ] 2.7.5 Duplicating an Input File Descriptor
  - [ ] 2.7.6 Duplicating an Output File Descriptor
  - [ ] 2.7.7 Open File Descriptors for Reading and Writing
- [ ] 2.8 Exit Status and Errors
  - [ ] 2.8.1 Consequences of Shell Errors
  - [ ] 2.8.2 Exit Status for Commands
- [ ] 2.9 Shell Commands
  - [ ] 2.9.1 Simple Commands
    - [ ] 2.9.1.1 Order of Processing
    - [ ] 2.9.1.2 Variable Assignments
    - [ ] 2.9.1.3 Commands with no Command Name
    - [ ] 2.9.1.4 Command Search and Execution
    - [ ] 2.9.1.5 Standard File Descriptors
    - [ ] 2.9.1.6 Non-built-in Utility Execution
  - [ ] 2.9.2 Pipelines
  - [ ] 2.9.3 Lists
    - [ ] 2.9.3.1 Asynchronous AND-OR Lists
    - [ ] 2.9.3.2 Sequential AND-OR Lists
    - [ ] 2.9.3.3 AND Lists
    - [ ] 2.9.3.4 OR Lists
  - [ ] 2.9.4 Compound Commands
    - [ ] 2.9.4.1 Grouping Commands
    - [ ] 2.9.4.2 The for Loop
    - [ ] 2.9.4.3 Case Conditional Construct
    - [ ] 2.9.4.4 The if Conditional Construct
    - [ ] 2.9.4.5 The while Loop
    - [ ] 2.9.4.6 The until Loop
  - [ ] 2.9.5 Function Definition Command
- [ ] 2.10 Shell Grammar
  - [ ] 2.10.1 Shell Grammar Lexical Conventions
  - [ ] 2.10.2 Shell Grammar Rules
- [ ] 2.11 Job Control
- [ ] 2.12 Signals and Error Handling
- [ ] 2.13 Shell Execution Environment
  - [ ] 2.14 Pattern Matching Notation
    - [ ] 2.14.1 Patterns Matching a Single Character
    - [ ] 2.14.2 Patterns Matching Multiple Characters
    - [ ] 2.14.3 Extended Patterns
