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

	if err := resolveFlags(&parsed, args); err != nil {
		return Parsed{}, err
	}

	if err := checkFlagGuards(&parsed, spec); err != nil {
		return Parsed{}, err
	}

	if err := checkGroupGuards(&parsed, spec); err != nil {
		return Parsed{}, err
	}

	return parsed, nil
}

func resolveFlags(parsed *Parsed, args []string) error {

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == "--":
			if i+1 < len(args) {
				parsed.Operands = append(parsed.Operands, args[i+1:]...)
			}
			return nil

		case strings.HasPrefix(arg, "--"):
			if err := parseLong(parsed, &i, args); err != nil {
				return err
			}

		case strings.HasPrefix(arg, "-") && arg != "-":
			if len(arg) > 2 {
				if err := parseShortCluster(parsed, &i, args); err != nil {
					return err
				}
			} else {
				if err := parseShort(parsed, &i, []rune(arg)[1], args); err != nil {
					return err
				}
			}

		default:
			parsed.Operands = append(parsed.Operands, args[i:]...)
			return nil
		}
	}

	return nil
}

func populateAliases(parsed *Parsed, spec Spec) error {
	for _, flag := range spec.Flags {

		if flag.Short != 0 {
			shortStr := string(flag.Short)

			if _, exists := parsed.aliases[shortStr]; exists {
				return fmt.Errorf("flags: overloaded short flag -%s\n", shortStr)
			}
			parsed.aliases[shortStr] = flag
		}

		if flag.Long != "" {
			if _, exists := parsed.aliases[flag.Long]; exists {
				return fmt.Errorf("flags: overloaded long flag --%s\n", flag.Long)
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
	return Flag{}, fmt.Errorf("flags: failed to resolve flag %s\n", name)
}

func parseShort(parsed *Parsed, i *int, r rune, args []string) error {
	shortStr := string(r)
	flag, exists := parsed.aliases[shortStr]

	if !exists {
		return fmt.Errorf("flags: unknown flag -%s\n", shortStr)
	}

	if _, exists := parsed.values[flag]; exists {
		return fmt.Errorf("flags: flag -%s was set multiple times\n", shortStr)
	}

	if flag.Parametrized {
		*i++ // expedite index one position
		if *i >= len(args) {
			return fmt.Errorf("flags: expected -%s to be followed by a parameter, found EOF\n", shortStr)
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
			return fmt.Errorf("flags: unknown flag -%s\n", shortStr)
		}

		if flag.Parametrized && idx != len(runes)-1 {
			return fmt.Errorf("flags: flag -%s requiring a parameter found in the middle of -%s, flags with parameter can only be at the very end of a cluster\n", shortStr, args[*i])
		}

		if err := parseShort(parsed, i, r, args); err != nil {
			return err
		}
	}

	return nil
}

func parseLong(parsed *Parsed, i *int, args []string) error {
	str := args[*i][2:] // strip leading --
	parts := strings.SplitN(str, "=", 2)

	if len(parts[0]) < 1 {
		return fmt.Errorf("flags: parseLong: --%s: long flag of 0 length\n", str)
	}

	flag, exists := parsed.aliases[parts[0]]

	if !exists {
		return fmt.Errorf("flags: unknown flag --%s\n", parts[0])
	}

	if _, exists := parsed.values[flag]; exists {
		return fmt.Errorf("flags: flag --%s was set multiple times\n", parts[0])
	}

	if !flag.Parametrized {
		if len(parts) > 1 {
			return fmt.Errorf("flags: flag --%s does not accept a parameter\n", parts[0])
		}
		parsed.values[flag] = ""
		return nil
	}

	switch len(parts) {
	case 1:
		*i++
		if *i >= len(args) {
			return fmt.Errorf("flags: expected --%s to be followed by a parameter, found EOF\n", parts[0])
		}
		parsed.values[flag] = args[*i]
	case 2:
		parsed.values[flag] = parts[1]
	default:
		return fmt.Errorf("flags: parseLong: --%s: unexpected error, I have no idea how you could have got here, hf troubleshooting :)\n", str)
	}

	return nil
}

func checkGroupGuards(parsed *Parsed, spec Spec) error {
	for _, group := range spec.Groups {
		count := 0
		for _, member := range group.Flags {
			if parsed.Has(member) {
				count++
			}
		}

		if group.Exclusive && count > 1 {
			return fmt.Errorf("flags: multiple mutualy exclusive flags set from [%s]\n", strings.Join(group.Flags, ","))
		}

		if group.Required && count < 1 {
			return fmt.Errorf("flags: at least one of [%s] must be set\n", strings.Join(group.Flags, ","))
		}
	}

	return nil
}

func checkFlagGuards(parsed *Parsed, spec Spec) error {
	for _, flag := range spec.Flags {
		var name, dashes string
		if flag.Short != 0 {
			name = string(flag.Short)
			dashes = "-"
		} else if flag.Long != "" {
			name = flag.Long
			dashes = "--"
		} else {
			// this could throw, but I opted to not care if a flag is unreachable due to not having
			// neither short nor long set, the alias will not be created either way
			continue
		}

		if flag.Required && !parsed.Has(name) {
			return fmt.Errorf("flags: required flag not set %s%s\n", dashes, name)
		}
	}

	return nil
}
