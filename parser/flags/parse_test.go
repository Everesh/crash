package flags

import "testing"

var (
	boolFlag  = Flag{Long: "verbose", Short: 'v'}
	paramFlag = Flag{Long: "output", Short: 'o', Parametrized: true}
)

func flagSpec(flags ...Flag) Spec { return Spec{Flags: flags} }

func TestParse(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		spec          Spec
		wantErr       bool
		wantHas       map[string]bool
		wantValues    map[string]string
		wantValueErrs []string
		wantOps       []string
	}{
		{
			name:    "bool flag long",
			args:    []string{"--verbose"},
			spec:    flagSpec(boolFlag),
			wantHas: map[string]bool{"verbose": true},
		},
		{
			name:    "bool flag short",
			args:    []string{"-v"},
			spec:    flagSpec(boolFlag),
			wantHas: map[string]bool{"v": true},
		},
		{
			name:       "parametrized flag long space-separated",
			args:       []string{"--output", "file.txt"},
			spec:       flagSpec(paramFlag),
			wantValues: map[string]string{"output": "file.txt"},
		},
		{
			name:       "parametrized flag long equals",
			args:       []string{"--output=file.txt"},
			spec:       flagSpec(paramFlag),
			wantValues: map[string]string{"output": "file.txt"},
		},
		{
			name:       "parametrized flag short",
			args:       []string{"-o", "out.txt"},
			spec:       flagSpec(paramFlag),
			wantValues: map[string]string{"o": "out.txt"},
		},
		{
			name:    "trailing operands stop flag parsing",
			args:    []string{"-v", "file1", "file2"},
			spec:    flagSpec(boolFlag),
			wantHas: map[string]bool{"verbose": true},
			wantOps: []string{"file1", "file2"},
		},
		{
			name:    "-- terminates flag parsing",
			args:    []string{"-v", "--", "--not-a-flag", "file"},
			spec:    flagSpec(boolFlag),
			wantHas: map[string]bool{"verbose": true},
			wantOps: []string{"--not-a-flag", "file"},
		},
		{
			name: "-- alone produces no operands",
			args: []string{"--"},
			spec: flagSpec(boolFlag),
		},
		{
			name:    "short cluster all bool",
			args:    []string{"-abc"},
			spec:    flagSpec(Flag{Short: 'a'}, Flag{Short: 'b'}, Flag{Short: 'c'}),
			wantHas: map[string]bool{"a": true, "b": true, "c": true},
		},
		{
			name:       "short cluster parametrized at end",
			args:       []string{"-vo", "file.txt"},
			spec:       flagSpec(Flag{Short: 'v'}, Flag{Short: 'o', Parametrized: true}),
			wantHas:    map[string]bool{"v": true},
			wantValues: map[string]string{"o": "file.txt"},
		},
		{
			name:    "empty args",
			args:    []string{},
			spec:    flagSpec(boolFlag),
			wantHas: map[string]bool{"verbose": false},
		},
		{
			name: "required flag present",
			args: []string{"-r"},
			spec: flagSpec(Flag{Long: "required", Short: 'r', Required: true}),
		},
		{
			name: "group required satisfied",
			args: []string{"--foo"},
			spec: Spec{
				Flags:  []Flag{{Long: "foo"}, {Long: "bar"}},
				Groups: []Group{{Flags: []string{"foo", "bar"}, Required: true}},
			},
		},
		{
			name: "group exclusive with one flag set",
			args: []string{"--foo"},
			spec: Spec{
				Flags:  []Flag{{Long: "foo"}, {Long: "bar"}},
				Groups: []Group{{Flags: []string{"foo", "bar"}, Exclusive: true}},
			},
		},
		{
			name:    "Has returns false for unset known flag",
			args:    []string{},
			spec:    flagSpec(boolFlag),
			wantHas: map[string]bool{"verbose": false},
		},
		{
			name:    "Has returns false for unknown flag",
			args:    []string{},
			spec:    flagSpec(boolFlag),
			wantHas: map[string]bool{"nonexistent": false},
		},
		{
			name:          "Value errors for non-parametrized flag",
			args:          []string{"-v"},
			spec:          flagSpec(boolFlag),
			wantValueErrs: []string{"v"},
		},
		{
			name:          "Value errors for unset parametrized flag",
			args:          []string{},
			spec:          flagSpec(paramFlag),
			wantValueErrs: []string{"output"},
		},
		{
			name:          "Value errors for unknown flag",
			args:          []string{},
			spec:          flagSpec(paramFlag),
			wantValueErrs: []string{"nonexistent"},
		},
		{
			name:    "unknown long flag",
			args:    []string{"--unknown"},
			spec:    flagSpec(boolFlag),
			wantErr: true,
		},
		{
			name:    "unknown short flag",
			args:    []string{"-z"},
			spec:    flagSpec(boolFlag),
			wantErr: true,
		},
		{
			name:    "duplicate long flag",
			args:    []string{"--verbose", "--verbose"},
			spec:    flagSpec(boolFlag),
			wantErr: true,
		},
		{
			name:    "duplicate short flag",
			args:    []string{"-v", "-v"},
			spec:    flagSpec(boolFlag),
			wantErr: true,
		},
		{
			name:    "parametrized flag missing value long",
			args:    []string{"--output"},
			spec:    flagSpec(paramFlag),
			wantErr: true,
		},
		{
			name:    "parametrized flag missing value short",
			args:    []string{"-o"},
			spec:    flagSpec(paramFlag),
			wantErr: true,
		},
		{
			name:    "bool flag given value via =",
			args:    []string{"--verbose=yes"},
			spec:    flagSpec(boolFlag),
			wantErr: true,
		},
		{
			name:    "parametrized flag in middle of cluster",
			args:    []string{"-ov", "file.txt"},
			spec:    flagSpec(Flag{Short: 'o', Parametrized: true}, Flag{Short: 'v'}),
			wantErr: true,
		},
		{
			name:    "required flag missing",
			args:    []string{},
			spec:    flagSpec(Flag{Long: "required", Short: 'r', Required: true}),
			wantErr: true,
		},
		{
			name:    "overloaded short flag",
			spec:    flagSpec(Flag{Short: 'v', Long: "verbose"}, Flag{Short: 'v', Long: "version"}),
			wantErr: true,
		},
		{
			name:    "overloaded long flag",
			spec:    flagSpec(Flag{Long: "output", Short: 'a'}, Flag{Long: "output", Short: 'b'}),
			wantErr: true,
		},
		{
			name: "group required not satisfied",
			args: []string{},
			spec: Spec{
				Flags:  []Flag{{Long: "foo"}, {Long: "bar"}},
				Groups: []Group{{Flags: []string{"foo", "bar"}, Required: true}},
			},
			wantErr: true,
		},
		{
			name: "group exclusive both flags set",
			args: []string{"--foo", "--bar"},
			spec: Spec{
				Flags:  []Flag{{Long: "foo"}, {Long: "bar"}},
				Groups: []Group{{Flags: []string{"foo", "bar"}, Exclusive: true}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Parse(tt.args, tt.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse(%v) error = %v, wantErr %v", tt.args, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			for name, want := range tt.wantHas {
				if got := p.Has(name); got != want {
					t.Errorf("Has(%q) = %v, want %v", name, got, want)
				}
				if got := p.Bool(name); got != want {
					t.Errorf("Bool(%q) = %v, want %v", name, got, want)
				}
			}
			for name, want := range tt.wantValues {
				got, err := p.Value(name)
				if err != nil {
					t.Errorf("Value(%q) unexpected error: %v", name, err)
					continue
				}
				if got != want {
					t.Errorf("Value(%q) = %q, want %q", name, got, want)
				}
			}
			for _, name := range tt.wantValueErrs {
				if _, err := p.Value(name); err == nil {
					t.Errorf("Value(%q): expected error, got nil", name)
				}
			}
			if len(p.Operands) != len(tt.wantOps) {
				t.Errorf("Operands = %v, want %v", p.Operands, tt.wantOps)
				return
			}
			for i, want := range tt.wantOps {
				if p.Operands[i] != want {
					t.Errorf("Operands[%d] = %q, want %q", i, p.Operands[i], want)
				}
			}
		})
	}
}
