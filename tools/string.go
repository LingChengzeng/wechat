// string
package tools

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"reflect"
	"strings"
	"unsafe"
)

func FindLower(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func Wrap(str, wraper string) string {
	return wraper + str + wraper
}

func StringToBoolean(str string) bool {
	rs := false
	switch strings.ToLower(strings.TrimSpace(str)) {
	case "1", "on", "true":
		return true
	case "0", "off", "false":
		return rs
	}

	return rs
}

// BytesToString accepts bytes and returns their string presentation
// instead of string() this method doesn't generate memory allocations,
// BUT it is not safe to use anywhere because it points
// this helps on 0 memory allocations
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

// StringToBytes accepts string and returns their []byte presentation
// instead of byte() this method doesn't generate memory allocations,
// BUT it is not safe to use anywhere because it points
// this helps on 0 memory allocations
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Serialize(m interface{}) (string, error) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(m)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func Deserialize(str string, m interface{}) error {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}

	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)

	//	d := gob.NewDecoder(bytes.NewBufferString(str))
	err = d.Decode(&m)
	if err != nil {
		return err
	}

	return nil
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rsLen := len(rs)
	end := 0

	if start < 0 {
		start = rsLen - 1 + start
	}

	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > rsLen {
		start = rsLen
	}

	if end < 0 {
		end = 0
	}

	if end > rsLen {
		end = rsLen
	}

	return string(rs[start:end])
}

// 获取两个字符串中间的字符串。此方法待完善。
// 例如：[aaaaa]bbbbbbb[aaaaaa]nnnnn[aaa], 获取中括号中的字符串
func GetStrBetweenStr(str, strA, strB string) []string {
	returnStr := []string{}
	tmpSlice1 := strings.Split(str, strA)
	if len(tmpSlice1) == 0 {
		return returnStr
	}

	for _, item := range tmpSlice1 {
		if len(item) == 0 {
			continue
		}

		itemSlice := strings.Split(item, strB)
		if len(itemSlice) > 0 {
			returnStr = append(returnStr, itemSlice[0])
		}
	}

	return returnStr
}
