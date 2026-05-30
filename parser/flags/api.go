package flags

type Flag struct {
	Long         string
	Short        byte
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
	aliases  map[string]string // short/long name -> canonical name
	set      map[string]string // canonical name -> value ("" for bools)
}

func (p Parsed) Has(name string) bool {
	_, ok := p.set[p.resolve(name)]
	return ok
}

func (p Parsed) Bool(name string) bool {
	return p.Has(name)
}

func (p Parsed) Value(name string) string {
	return p.set[p.resolve(name)]
	// bools and non-existent flags return "", but since "" is also a valid value for
	// parametrized flags, Im not going to guard against this. Best of luck when debugging •ᴗ•
}
