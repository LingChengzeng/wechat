// tools
package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
)

// 浮点数四舍五入
func Round(num float64, precision int64) float64 {
	strPrecision := String(precision)
	return Float64(fmt.Sprintf("%."+strPrecision+"f", num))
}

// json to map
func DecodeJson(jsonStr string) interface{} {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err == nil {
		return data
	}
	return nil
}

func ToJSON(param interface{}) string {
	jsonByte, _ := json.Marshal(param)
	return string(jsonByte)
}

func CheckError(err error, exit bool) {
	if err != nil {
		fmt.Println("Error: ", err)
		if exit == true {
			os.Exit(1)
		}
	}
}

// 发送GET和POST请求
func Request(httpURL, method string, params map[string]string) (string, error) {
	data := url.Values{}
	if len(params) > 0 {
		for key, value := range params {
			data.Add(key, value)
		}
	}

	client := &http.Client{}
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest(method, httpURL, body)
	if err != nil {
		return "", errors.New("The http request error")
	}

	req.Header.Set("Content-Type", "text/html; charset=UTF-8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("User-Agent", "BCCS_SDK/3.0 (CentOS 6.6) Golang/1.5 (Baidu Push Server SDK V3.0.0 and so on..) Unknown")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}

// 发送POST请求, 返回Body
func PostString(url, paramStr string) ([]byte, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(paramStr))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Determine whether a variable is empty
func IsEmpty(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return IsEmpty(v.Elem().Interface())
	}

	return false
}

func StringSliceToInForSQL(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	tmpSlice := []string{}
	for _, str := range slice {
		tmpSlice = append(tmpSlice, `'`+str+`'`)
	}

	return " IN (" + strings.Join(tmpSlice, ",") + ")"
}

func TrimMapValue(data map[string]string) map[string]string {
	if data == nil || len(data) == 0 {
		return nil
	}

	lock := new(sync.RWMutex)
	lock.Lock()
	defer lock.Unlock()

	tmpMap := map[string]string{}
	for key, val := range data {
		tmpMap[key] = strings.TrimSpace(val)
	}
	return tmpMap
}

// Quote string with slashes
// Returns a string with backslashes before characters that
// need to be escaped. These characters are single quote ('),
// double quote (") and backslash (\)
func Addslashes(str string) string {
	if len(str) == 0 {
		return ""
	}

	tmpSlice := strings.Split(str, "")
	newSlice := []string{}
	for _, item := range tmpSlice {
		newItem := item
		if item == `'` || item == `"` || item == `\` {
			newItem = `\` + item
		}

		newSlice = append(newSlice, newItem)
	}

	return strings.Join(newSlice, "")
}

func AddslashesPlus(str string) string {
	if len(str) == 0 {
		return ""
	}

	tmpSlice := strings.Split(str, "")
	newSlice := []string{}
	for _, item := range tmpSlice {
		newItem := item
		if item == `'` || item == `"` || item == `\` || item == `%` || item == `_` {
			newItem = `\` + item
		}

		newSlice = append(newSlice, newItem)
	}

	return strings.Join(newSlice, "")
}

// GetIP returns the ip address of a client.
func GetIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		ipnet, ok := address.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			return ipnet.IP.String(), nil
		}
	}
	return "", nil
}

// Ticket generates a ticket for user.
func Ticket(key, Account string) string {
	originStr := Md5hash(key) + Md5hash(Account)
	return Md5hash(Md5hash(originStr[16:] + originStr[:16]))
}
