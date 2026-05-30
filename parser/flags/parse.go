package flags

import "fmt"

func Parse(args []string, spec Spec) (Parsed, error) {
	return Parsed{}, fmt.Errorf("What are you expecting? You havent implemented the flags parser yet. (• - •)")
}

func (f Flag) canonical() string {
	if f.Long != "" {
		return f.Long
	}
	return string(f.Short)
}

func (p Parsed) resolve(name string) string {
	if c, ok := p.aliases[name]; ok {
		return c
	}
	return name
}
