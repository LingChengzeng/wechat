package tools

import (
	"strings"
	"testing"
)

// func TestGetSliceValueString(t *testing.T) {
// 	testSlice := []string{"A", "B", "C"}
// 	rs := GetSliceValueString(testSlice, 2, "BB")
// 	t.Log(rs)
// }

func TestSp(t *testing.T) {
	str := "nasdfasfssss, aaaa"
	rs := strings.Split(str, ",")
	for _, item := range rs {
		t.Log(item)
	}

	t.Log(strings.Join(rs[2:], ","))
}
