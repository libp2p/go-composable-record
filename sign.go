package sr

import (
	"io"
)

type Signed struct {
	By   Peer
	Sign Blob
	Stmt Node
	Dict
}

func (v Signed) WritePretty(w io.Writer, level int) error {
	panic("XXX")
}
