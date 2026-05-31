package bus

type Queue struct{ cmds []Cmd }

func NewQueue() *Queue { return &Queue{} }

func (q *Queue) Send(cmd Cmd) {
	q.cmds = append(q.cmds, cmd)
}

func (q *Queue) Drain() []Cmd {
	cmds := q.cmds
	q.cmds = nil
	return cmds
}
