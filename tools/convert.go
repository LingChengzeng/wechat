// convert
/*
数据类型之间的转换

var s := "168000"
i, _ := toolbox.Int(s)
i32, _ := toolbox.Int32(s)
i64, _ := toolbox.Int64(s)
f32, _ := toolbox.Float32(s)
f64, _ := toolbox.Float64(s)
*/
package tools

import (
	"html/template"
	"strconv"
	"strings"
)

// args: value, precision(only for float)
func String(args ...interface{}) string {
	value := args[0]
	var precision int = 12 // default

	switch value.(type) {
	case string:
		v, _ := value.(string)
		return v
	case int:
		v, _ := value.(int)
		return strconv.Itoa(v)
	case int32:
		v, _ := value.(int32)
		return strconv.FormatInt(int64(v), 10)
	case int64:
		v, _ := value.(int64)
		return strconv.FormatInt(v, 10)
	case float32:
		v, _ := value.(float32)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(float64(v), 'f', precision, 64)
	case float64:
		v, _ := value.(float64)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(v, 'f', precision, 64)
	case template.HTML:
		return string(value.(template.HTML))
	default:
		return ""
	}
}

func Int(value interface{}) int {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		rs, _ := strconv.Atoi(v)
		return rs
	case int:
		v, _ := value.(int)
		return v
	case int32:
		v, _ := value.(int32)
		return int(v)
	case int64:
		v, _ := value.(int64)
		return int(v)
	case float32:
		v, _ := value.(float32)
		return int(v)
	case float64:
		v, _ := value.(float64)
		return int(v)
	default:
		return int(0)
	}
}

func Int32(value interface{}) int32 {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, _ := strconv.ParseInt(v, 10, 32)
		return int32(result)
	case int:
		v, _ := value.(int)
		return int32(v)
	case int32:
		v, _ := value.(int32)
		return int32(v)
	case int64:
		v, _ := value.(int64)
		return int32(v)
	case float32:
		v, _ := value.(float32)
		return int32(v)
	case float64:
		v, _ := value.(float64)
		return int32(v)
	default:
		return int32(0)
	}
}

func Int64(value interface{}) int64 {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		rs, _ := strconv.ParseInt(v, 10, 32)
		return rs
	case int:
		v, _ := value.(int)
		return int64(v)
	case int32:
		v, _ := value.(int32)
		return int64(v)
	case int64:
		v, _ := value.(int64)
		return v
	case float32:
		v, _ := value.(float32)
		return int64(v)
	case float64:
		v, _ := value.(float64)
		return int64(v)
	default:
		return int64(0)
	}
}

func Float32(value interface{}) float32 {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, _ := strconv.ParseFloat(v, 32)
		return float32(result)
	case int:
		v, _ := value.(int)
		return float32(v)
	case int32:
		v, _ := value.(int32)
		return float32(v)
	case int64:
		v, _ := value.(int64)
		return float32(v)
	case float32:
		v, _ := value.(float32)
		return v
	case float64:
		v, _ := value.(float64)
		return float32(v)
	default:
		return float32(0)
	}
}

func Float64(value interface{}) float64 {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		rs, _ := strconv.ParseFloat(v, 64)
		return rs
	case int:
		v, _ := value.(int)
		return float64(v)
	case int32:
		v, _ := value.(int32)
		return float64(v)
	case int64:
		v, _ := value.(int64)
		return float64(v)
	case float32:
		v, _ := value.(float32)
		return float64(v)
	case float64:
		v, _ := value.(float64)
		return v
	default:
		return float64(0)
	}
}

// //  123(true), 0(false),"-123"(true), "on"(true), "off"(false), "true"(true), "false"(false)
func Boolean(value interface{}) bool {
	switch value.(type) {
	case bool:
		return value.(bool)
	case int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32, uint64:
		return value != 0
	case []byte:
		if val, err := strconv.ParseBool(string(value.([]byte))); err == nil {
			return val
		} else if val, err := strconv.ParseFloat(string(value.([]byte)), 32); err == nil {
			return val != 0
		}

		switch strings.ToLower(strings.TrimSpace(string(value.([]byte)))) {
		case "on":
			return true
		case "off":
			return false
		case "true":
			return true
		case "false":
			return false
		}
	case string:
		if val, err := strconv.ParseBool(value.(string)); err == nil {
			return val
		} else if val, err := strconv.ParseFloat(value.(string), 32); err == nil {
			return val != 0
		}

		switch strings.ToLower(strings.TrimSpace(value.(string))) {
		case "on":
			return true
		case "off":
			return false
		case "true":
			return true
		case "false":
			return false
		}
	default:
		return false
	}

	return false
}
