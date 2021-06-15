package foo

import (
	"strings"
	"testing"
)

func TestFoo(t *testing.T) {
	str := foo()
	if !strings.Contains(str, "foo") {
		t.Errorf("expected foo to be called")
	}
}