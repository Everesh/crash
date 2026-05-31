package bus

// Handler is implemented by anything that can receive shell signals.
// Adding a new Cmd type requires adding a method here
type Handler interface {
	Exit(code int)
}

// Cmd is implemented by signals
type Cmd interface {
	Apply(Handler)
}

// --- CMDs ---

type ExitCmd struct{ Code int }

func (c ExitCmd) Apply(h Handler) { h.Exit(c.Code) }
