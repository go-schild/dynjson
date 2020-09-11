package dynjson

import (
	"encoding/json"
)

type JsonListRaw []interface{}

type JsonList []JsonListItem

type JsonListItem struct {
	data interface{}
}

func NewJsonList(raw JsonListRaw) JsonList {
	result := make(JsonList, 0)

	if raw != nil {
		for _, data := range raw {
			result = append(result, JsonListItem{data})
		}
	}

	return result
}

func ParseList(jsonString string) (JsonList, error) {
	obj := JsonListRaw{}
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		return nil, err
	}

	return NewJsonList(obj), nil
}

// ToString returns the JSON data as string.
// Returns an empty string, when an error occurred.
func (j JsonList) ToString() string {
	data, _ := json.Marshal(&j)
	return string(data)
}

func (j *JsonList) Append(data ...interface{}) {
	var list []JsonListItem

	for _, dataEntry := range data {
		switch dataEntry.(type) {
		case int:
			list = append(list, JsonListItem{data: float64(dataEntry.(int))})
		case float32:
			list = append(list, JsonListItem{data: float64(dataEntry.(float32))})
		default:
			list = append(list, JsonListItem{data: data})
		}
	}

	*j = append(*j, list...)
}

func (j *JsonList) Prepend(data ...interface{}) {
	var list []JsonListItem

	for _, dataEntry := range data {
		switch dataEntry.(type) {
		case int:
			list = append(list, JsonListItem{data: float64(dataEntry.(int))})
		case float32:
			list = append(list, JsonListItem{data: float64(dataEntry.(float32))})
		default:
			list = append(list, JsonListItem{data: data})
		}
	}

	*j = append(list, *j...)
}

func (j *JsonListItem) ObjectOk() (JsonObject, bool) {
	return convToObject(j.data)
}

func (j *JsonListItem) Object() JsonObject {
	val, _ := j.ObjectOk()
	return val
}

func (j *JsonListItem) ListOk() (JsonList, bool) {
	return convToList(j.data)
}

func (j *JsonListItem) List() JsonList {
	val, _ := j.ListOk()
	return val
}

func (j *JsonListItem) StringOk() (string, bool) {
	return convToString(j.data)
}

func (j *JsonListItem) StringDefault(def string) string {
	val, ok := j.StringOk()
	if ok {
		return val
	}
	return def
}

func (j *JsonListItem) String() string {
	val, _ := j.StringOk()
	return val
}

func (j *JsonListItem) Float64Ok() (float64, bool) {
	return convToFloat64(j.data)
}

func (j *JsonListItem) Float64Default(def float64) float64 {
	val, ok := j.Float64Ok()
	if ok {
		return val
	}
	return def
}

func (j *JsonListItem) Float64() float64 {
	val, _ := j.Float64Ok()
	return val
}

func (j *JsonListItem) Float32Ok() (float32, bool) {
	val, ok := j.Float64Ok()
	return float32(val), ok
}

func (j *JsonListItem) Float32Default(def float32) float32 {
	val, ok := j.Float32Ok()
	if ok {
		return val
	}
	return def
}

func (j *JsonListItem) Float32() float32 {
	val, _ := j.Float32Ok()
	return val
}

func (j *JsonListItem) IntOk() (int, bool) {
	val, ok := j.Float64Ok()
	return int(val), ok
}

func (j *JsonListItem) IntDefault(def int) int {
	val, ok := j.IntOk()
	if ok {
		return val
	}
	return def
}

func (j *JsonListItem) Int() int {
	val, _ := j.IntOk()
	return val
}
