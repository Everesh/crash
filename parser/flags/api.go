package flags

type Flag struct {
	Long         string
	Short        rune
	Required     bool
	Parametrized bool
}

type Group struct {
	Flags     []string
	Required  bool // at least 1 must be set
	Exclusive bool // at most 1 can be set
}

type Spec struct {
	Flags  []Flag
	Groups []Group
}

type Parsed struct {
	Operands []string
	flags    []Flag         // indexed copy of spec.Flags
	aliases  map[string]int // name -> flags index
	values   map[int]string // flags index -> value
}

func (p Parsed) Has(name string) bool {
	idx, err := p.resolve(name)
	if err != nil {
		return false
	}

	_, ok := p.values[idx]
	return ok
}

func (p Parsed) Bool(name string) bool {
	return p.Has(name)
}

func (p Parsed) Value(name string) (string, error) {
	idx, err := p.resolve(name)
	if err != nil {
		return "", err
	}

	return p.values[idx], nil
}
