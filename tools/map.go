// map
package tools

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"reflect"
	"strings"
	"sync"
)

// HMap is a map with lock
type HMap struct {
	lock *sync.RWMutex
	SMap map[string]interface{}
}

// NewMap returns a new safe map.
func NewMap(smap ...map[string]interface{}) *HMap {
	_SMap := map[string]interface{}{}
	if len(smap) > 0 {
		_SMap = smap[0]
	}

	return &HMap{
		lock: new(sync.RWMutex),
		SMap: _SMap,
	}
}

func (this *HMap) Get(key string) interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if val, ok := this.SMap[key]; ok {
		return val
	}

	return nil
}

func (this *HMap) Set(key string, value interface{}) bool {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.SMap == nil {
		this.SMap = map[string]interface{}{}
	}

	this.SMap[key] = value
	return true
}

func (this *HMap) Delete(key string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.SMap, key)
}

// Checker returns true is key is exist in the map, otherwise returns false.
func (this *HMap) Check(key string) bool {
	this.lock.RLock()
	defer this.lock.RUnlock()
	_, ok := this.SMap[key]
	return ok
}

// GetMap returns all item in the safe map.
func (this *HMap) GetMap() map[string]interface{} {
	return this.SMap
}

// Count returns the number of items within the map
func (this *HMap) Count() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return len(this.SMap)
}

func (this *HMap) Remove(key string, defaultValue interface{}) interface{} {
	value := this.Get(key)
	if value != nil {
		this.Delete(key)
		return value
	}

	return defaultValue
}

func (this *HMap) IToS(value interface{}) string {
	return strings.TrimSpace(String(value))
}

func (this *HMap) GetToString(key string) string {
	return strings.TrimSpace(String(this.Get(key)))
}

func (this *HMap) GetToInt(key string) int {
	return Int(this.GetToString(key))
}

func (this *HMap) GetToInt32(key string) int32 {
	return Int32(this.GetToString(key))
}

func (this *HMap) GetToInt64(key string) int64 {
	return Int64(this.GetToString(key))
}

func (this *HMap) GetToFloat32(key string) float32 {
	return Float32(this.GetToString(key))
}

func (this *HMap) GetToFloat64(key string) float64 {
	return Float64(this.GetToString(key))
}

func (this *HMap) GetToMap(key string) map[string]interface{} {
	val := this.Get(key)
	if val != nil {
		if myMap, ok := val.(map[string]interface{}); ok {
			return myMap
		}
	}

	return map[string]interface{}{}
}

// parses struct to map
func ConvertToMap(model interface{}) map[string]interface{} {
	ret := map[string]interface{}{}

	modelReflect := reflect.ValueOf(model)

	if modelReflect.Kind() == reflect.Ptr {
		modelReflect = modelReflect.Elem()
	}

	modelRefType := modelReflect.Type()
	fieldsCount := modelReflect.NumField()

	var fieldData interface{}

	for i := 0; i < fieldsCount; i++ {
		field := modelReflect.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			fallthrough
		case reflect.Ptr:
			fieldData = ConvertToMap(field.Interface())
		default:
			fieldData = field.Interface()
		}

		ret[modelRefType.Field(i).Name] = fieldData
	}

	return ret
}

// ParseXMLToMap parses xml reading from xmlReader and returns the first-level sub-node key-value set,
// if the first-level sub-node contains child nodes, skip it.
func ParseXMLToMap(xmlReader io.Reader) (m map[string]string, err error) {
	if xmlReader == nil {
		err = errors.New("nil xmlReader")
		return
	}

	m = make(map[string]string)
	var (
		d     = xml.NewDecoder(xmlReader)
		tk    xml.Token
		depth = 0 // current xml.Token depth
		key   string
		value bytes.Buffer
	)

	for {
		tk, err = d.Token()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}

		switch v := tk.(type) {
		case xml.StartElement:
			depth++
			switch depth {
			case 2:
				key = v.Name.Local
				value.Reset()
			case 3:
				if err = d.Skip(); err != nil {
					return
				}
				depth--
				key = "" // key == "" indicates that the node with depth==2 has children
			}
		case xml.CharData:
			if depth == 2 && key != "" {
				value.Write(v)
			}
		case xml.EndElement:
			if depth == 2 && key != "" {
				m[key] = value.String()
			}
			depth--
		}
	}
}

// FormatMapToXML marshal map[string]string to xmlWriter with xml format, the root node name is xml.
//  NOTE: This function assumes the key of m map[string]string are legitimate xml name string
//  that does not contain the required escape character!
func FormatMapToXML(xmlWriter io.Writer, m map[string]string) (err error) {
	if xmlWriter == nil {
		return errors.New("nil xmlWriter")
	}

	if _, err = io.WriteString(xmlWriter, "<xml>"); err != nil {
		return
	}

	for k, v := range m {
		if _, err = io.WriteString(xmlWriter, "<"+k+">"); err != nil {
			return
		}

		if err = xml.EscapeText(xmlWriter, []byte(v)); err != nil {
			return
		}

		if _, err = io.WriteString(xmlWriter, "</"+k+">"); err != nil {
			return
		}
	}

	if _, err = io.WriteString(xmlWriter, "</xml>"); err != nil {
		return
	}

	return
}

type TreeMap struct {
	PrimaryKey     string // 主键字段名
	ParentFieldKey string // 父级ID字段名
	ChildrenKey    string // 子叶Key名称

	NewTreeForData []map[string]interface{}
}

// ⊢
func (this *TreeMap) FormatTree(list []map[string]interface{}, parentField,
	parentID, space string) []map[string]interface{} {
	if len(list) == 0 {
		return this.NewTreeForData
	}

	for _, val := range list {
		valMap := NewMap(val)
		if valMap.GetToString(parentField) == parentID {
			valMap.Set("Space", space)
			this.NewTreeForData = append(this.NewTreeForData, valMap.GetMap())
			this.FormatTree(list, parentField, valMap.GetToString("ID"), space+"⏤")
		}
	}

	return this.NewTreeForData
}

func (this *TreeMap) MapToTree(originMap []map[string]interface{}, parentID string) []map[string]interface{} {
	if originMap == nil || len(originMap) == 0 {
		return nil
	}

	if len(this.PrimaryKey) == 0 {
		this.PrimaryKey = "ID"
	}

	if len(this.ParentFieldKey) == 0 {
		this.ParentFieldKey = "ParentID"
	}

	if len(this.ChildrenKey) == 0 {
		this.ChildrenKey = "Children"
	}

	newTree := []map[string]interface{}{}
	for _, _mapItem := range originMap {
		mapItem := NewMap(_mapItem)
		if parentID == mapItem.Get(this.ParentFieldKey) {
			tmpItem := this.MapToTree(originMap, mapItem.GetToString(this.PrimaryKey))
			if len(tmpItem) > 0 {
				mapItem.Set(this.ChildrenKey, tmpItem)
			}

			newTree = append(newTree, mapItem.GetMap())
		}
	}

	return newTree
}

func MapToTree(list interface{}, parentID string, childFieldName string) []map[string]interface{} {
	if len(childFieldName) == 0 {
		childFieldName = "Children"
	}

	listData, ok := list.([]map[string]interface{})
	if !ok {
		return nil
	}

	tree := []map[string]interface{}{}
	for _, val := range listData {
		if parentID == String(val["ParentID"]) {
			tmp := MapToTree(listData, String(val["ParentID"]), childFieldName)
			if len(tmp) > 0 {
				val[childFieldName] = tmp
			}

			tree = append(tree, val)
		}
	}

	return tree
}

func GetSliceMapValueByKey(params []map[string]interface{}, key string) string {
	if len(params) == 0 || len(key) == 0 {
		return ""
	}

	for _, item := range params {
		itemMap := NewMap(item)
		if itemMap.Check(key) {
			return itemMap.GetToString(key)
		}
	}

	return ""
}
