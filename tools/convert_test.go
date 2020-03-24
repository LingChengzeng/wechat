package tools

import (
	"testing"
)

func TestBoolean(t *testing.T) {
	val := 11
	rs := Boolean(val)
	t.Log(rs)
}
