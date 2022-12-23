package trace

import (
	"io"
)

func New(w io.Writer) Tracer {
	return &tracer{w}
}

func Off() Tracer {
	return &nilTracer{}
}
