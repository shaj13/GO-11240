package bar

import (
	"strings"
	"testing"
)

func TestBar(t *testing.T) {
	str := bar()
	if !strings.Contains(str, "bar") {
		t.Errorf("expected bar to be called")
	}
}