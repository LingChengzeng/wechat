// slice
package tools

import (
	"strings"
)

// delete a item from slice.
func DeleteSliceItem(slice []interface{}, index int) []interface{} {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func UniqueSlice(slice *[]string) {
	found := make(map[string]bool)
	total := 0
	for key, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[key]
			total++
		}
	}

	*slice = (*slice)[:total]
}

func CheckInStrSlice(key string, strSlice []string) bool {
	hasKey := false
	if len(strSlice) == 0 {
		return hasKey
	}

	for _, item := range strSlice {
		if key == item {
			hasKey = true
			break
		}
	}

	return hasKey
}

func CheckInIntSlice(key int, intSlice []int) bool {
	hasKey := false
	if len(intSlice) == 0 {
		return hasKey
	}

	for _, item := range intSlice {
		if key == item {
			hasKey = true
			break
		}
	}

	return hasKey
}

// 根据给定的index获取slice对应的值
func GetSliceValueString(slice []string, index int, defaultValue string) string {
	sliceLen := len(slice)
	if sliceLen == 0 || index < 0 {
		if len(defaultValue) > 0 {
			return defaultValue
		}

		return ""
	}

	if sliceLen < index+1 {
		return defaultValue
	}

	return strings.TrimSpace(slice[index])
}

func StringToSlice(str, sep string) []string {
	slice := strings.Split(str, sep)
	tmpSlice := []string{}
	for _, item := range slice {
		tmpSlice = append(tmpSlice, strings.TrimSpace(item))
	}
	return tmpSlice
}
