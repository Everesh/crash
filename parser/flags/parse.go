package flags

import (
	"fmt"
	"strings"
)

func Parse(args []string, spec Spec) (Parsed, error) {

	parsed := Parsed{
		Operands: make([]string, 0),
		flags:    spec.Flags,
		aliases:  make(map[string]int),
		values:   make(map[int]string),
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
	for i, flag := range spec.Flags {

		if flag.Short != 0 {
			shortStr := string(flag.Short)

			if _, exists := parsed.aliases[shortStr]; exists {
				return fmt.Errorf("flags: overloaded short flag -%s", shortStr)
			}
			parsed.aliases[shortStr] = i
		}

		if flag.Long != "" {
			if _, exists := parsed.aliases[flag.Long]; exists {
				return fmt.Errorf("flags: overloaded long flag --%s", flag.Long)
			}
			parsed.aliases[flag.Long] = i
		}
	}

	return nil
}

func (p Parsed) resolve(name string) (int, error) {
	if idx, ok := p.aliases[name]; ok {
		return idx, nil
	}
	return -1, fmt.Errorf("flags: failed to resolve flag %s", name)
}

func parseShort(parsed *Parsed, i *int, r rune, args []string) error {
	shortStr := string(r)
	idx, exists := parsed.aliases[shortStr]

	if !exists {
		return fmt.Errorf("flags: unknown flag -%s", shortStr)
	}

	if _, exists := parsed.values[idx]; exists {
		return fmt.Errorf("flags: flag -%s was set multiple times", shortStr)
	}

	if parsed.flags[idx].Parametrized {
		*i++
		if *i >= len(args) {
			return fmt.Errorf("flags: expected -%s to be followed by a parameter, found EOF", shortStr)
		}
		parsed.values[idx] = args[*i]
	} else {
		parsed.values[idx] = ""
	}

	return nil
}

func parseShortCluster(parsed *Parsed, i *int, args []string) error {
	runes := []rune(args[*i][1:])

	for pos, r := range runes {
		shortStr := string(r)
		flagIdx, exists := parsed.aliases[shortStr]

		if !exists {
			return fmt.Errorf("flags: unknown flag -%s", shortStr)
		}

		if parsed.flags[flagIdx].Parametrized && pos != len(runes)-1 {
			return fmt.Errorf("flags: flag -%s requiring a parameter found in the middle of -%s, flags with parameter can only be at the very end of a cluster", shortStr, args[*i])
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
		return fmt.Errorf("flags: parseLong: --%s: long flag of 0 length", str)
	}

	flagIdx, exists := parsed.aliases[parts[0]]

	if !exists {
		return fmt.Errorf("flags: unknown flag --%s", parts[0])
	}

	if _, exists := parsed.values[flagIdx]; exists {
		return fmt.Errorf("flags: flag --%s was set multiple times", parts[0])
	}

	if !parsed.flags[flagIdx].Parametrized {
		if len(parts) > 1 {
			return fmt.Errorf("flags: flag --%s does not accept a parameter", parts[0])
		}
		parsed.values[flagIdx] = ""
		return nil
	}

	switch len(parts) {
	case 1:
		*i++
		if *i >= len(args) {
			return fmt.Errorf("flags: expected --%s to be followed by a parameter, found EOF", parts[0])
		}
		parsed.values[flagIdx] = args[*i]
	case 2:
		parsed.values[flagIdx] = parts[1]
	default:
		panic("flags: parseLong: unreachable — SplitN(s, \"=\", 2) cannot return more than 2 parts")
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
			return fmt.Errorf("flags: multiple mutually exclusive flags set from [%s]", strings.Join(group.Flags, ","))
		}

		if group.Required && count < 1 {
			return fmt.Errorf("flags: at least one of [%s] must be set", strings.Join(group.Flags, ","))
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
			continue
		}

		if flag.Required && !parsed.Has(name) {
			return fmt.Errorf("flags: required flag not set %s%s", dashes, name)
		}
	}

	return nil
}
