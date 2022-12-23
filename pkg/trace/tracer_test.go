package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buff bytes.Buffer
	tracer := New(&buff)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		msg := "Hello trace package"
		tracer.Trace(msg)
		if buff.String() != msg+"\n" {
			t.Errorf("Trace should not write '%s'.", buff.String())
		}
	}
}
