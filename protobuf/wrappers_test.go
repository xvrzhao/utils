package protobuf

import (
	"testing"
)

func TestFalseValue(t *testing.T) {
	f := FalseValue()
	if f == nil {
		t.Fail()
		return
	}
	if f.Value != false {
		t.Fail()
		return
	}
}

func TestTrueValue(t *testing.T) {
	tr := TrueValue()
	if tr == nil {
		t.Fail()
		return
	}
	if tr.Value != true {
		t.Fail()
		return
	}
}
