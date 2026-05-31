package streams

import (
	"fmt"
	"io"

	b "github.com/Everesh/crash/streams/bus"
)

type Io struct {
	In      io.Reader
	Out     io.Writer
	Err     io.Writer
	signals *b.Queue
}

func NewIo(in io.Reader, out io.Writer, err io.Writer) Io {
	return Io{
		In:      in,
		Out:     out,
		Err:     err,
		signals: b.NewQueue(),
	}
}

func (i Io) Send(cmd b.Cmd) {
	i.signals.Send(cmd)
}

func (i Io) Drain() []b.Cmd {
	return i.signals.Drain()
}

func (i Io) WriteErr(format string, a ...any) {
	fmt.Fprintf(i.Err, format+"\n", a...)
}
