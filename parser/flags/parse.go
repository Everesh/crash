package flags

import (
	"fmt"
	"strings"
)

func Parse(args []string, spec Spec) (Parsed, error) {
	parsed := Parsed{
		Operands: make([]string, 0),
		aliases:  make(map[string]Flag),
		values:   make(map[Flag]string),
	}

	if err := populateAliases(&parsed, spec); err != nil {
		return Parsed{}, err
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == "--":
			if i+1 < len(args) {
				parsed.Operands = append(parsed.Operands, args[i+1:]...)
			}
			break

		case strings.HasPrefix(arg, "--"):
			parseLong(&parsed, &i, spec, args)

		case strings.HasPrefix(arg, "-") && arg != "-":
			if len(arg) > 2 {
				if err := parseShortCluster(&parsed, &i, args); err != nil {
					return Parsed{}, err
				}
			} else {
				if err := parseShort(&parsed, &i, rune(arg[1]), args); err != nil {
					return Parsed{}, err
				}
			}

		default:
			parsed.Operands = append(parsed.Operands, args[i:]...)
			break
		}
	}

	return parsed, nil
}

func populateAliases(parsed *Parsed, spec Spec) error {
	for _, flag := range spec.Flags {

		if flag.Short != 0 {
			shortStr := string(flag.Short)

			if _, exists := parsed.aliases[shortStr]; exists {
				return fmt.Errorf("flags: overloaded short flag -%s", shortStr)
			}
			parsed.aliases[shortStr] = flag
		}

		if flag.Long != "" {
			if _, exists := parsed.aliases[flag.Long]; exists {
				return fmt.Errorf("flags: overloaded long flag --%s", flag.Long)
			}
			parsed.aliases[flag.Long] = flag
		}
	}

	return nil
}

func (p Parsed) resolve(name string) (Flag, error) {
	if c, ok := p.aliases[name]; ok {
		return c, nil
	}
	return Flag{}, fmt.Errorf("flags: failed to resolve flag %s", name)
}

func parseShort(parsed *Parsed, i *int, r rune, args []string) error {
	shortStr := string(r)
	flag, exists := parsed.aliases[shortStr]

	if !exists {
		return fmt.Errorf("flags: unknown flag -%s", shortStr)
	}

	if _, exists := parsed.values[flag]; exists {
		return fmt.Errorf("flags: flag -%s was set multiple times", shortStr)
	}

	if flag.Parametrized {
		*i++ // expedite index one position
		if *i >= len(args) {
			return fmt.Errorf("flags: expected -%s to be followed by a parameter, found EOF", shortStr)
		}
		parsed.values[flag] = args[*i]
	} else {
		parsed.values[flag] = "" // Boolean flag set
	}

	return nil
}

func parseShortCluster(parsed *Parsed, i *int, args []string) error {
	runes := []rune(args[*i][1:])

	for idx, r := range runes {
		shortStr := string(r)
		flag, exists := parsed.aliases[shortStr]

		// this check is a duplicate from the parseShort
		// since we need to see if the flag is parametrized its unavoidable here
		if !exists {
			return fmt.Errorf("flags: unknown flag -%s", shortStr)
		}

		if flag.Parametrized && idx != len(runes)-1 {
			return fmt.Errorf("flags: flag -%s requiring a parameter found in the middle of -%s, flags with parameter can only be at the very end of a cluster", shortStr, args[*i])
		}

		if err := parseShort(parsed, i, r, args); err != nil {
			return err
		}
	}

	return nil
}
