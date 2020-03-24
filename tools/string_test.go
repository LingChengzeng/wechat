package tools

import (
	"testing"
)

func TestS(t *testing.T) {
	str := "[aaaaa]bbbbbbb[aaaaaa]nnnnn[aaa]"
	rs := GetStrBetweenStr(str, "[", "]")
	for _, s := range rs {
		t.Log(s, len(s))
	}
}
