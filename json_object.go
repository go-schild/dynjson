package dynjson

import (
	"encoding/json"
)

type JsonObject map[string]interface{}

func NewJsonObject() JsonObject {
	return JsonObject{}
}

// ParseObject parses a string containing a json and returns a json object or an error
func ParseObject(jsonString string) (JsonObject, error) {
	obj := JsonObject{}
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// ToString returns the JSON data as string.
// Returns an empty string, when an error occurred.
func (j JsonObject) ToString() string {
	data, _ := json.Marshal(&j)
	return string(data)
}

// Has checks if a json object contains a specific field.
func (j JsonObject) Has(field string) bool {
	_, ok := j[field]
	return ok
}

// ObjectOk returns a field which contains an object and return it with true, when it matches.
func (j JsonObject) ObjectOk(field string) (JsonObject, bool) {
	if val, ok := j[field]; ok {
		return convToObject(val)
	}
	return nil, false
}

// Object returns an object inside the current object
func (j JsonObject) Object(field string) JsonObject {
	val, _ := j.ObjectOk(field)
	return val
}

// ListOk returns a list / array from the json object and a boolean which indicates, whether the result is ok.
func (j JsonObject) ListOk(field string) (JsonList, bool) {
	if val, ok := j[field]; ok {
		return convToList(val)
	}

	return nil, false
}

// List returns a list / object from the json object.
func (j JsonObject) List(field string) JsonList {
	val, _ := j.ListOk(field)
	return val
}

// StringOk returns a string from the json object and a boolean which indicates, whether the result is ok.
func (j JsonObject) StringOk(field string) (string, bool) {
	if val, ok := j[field]; ok {
		return convToString(val)
	}
	return "", false
}

func (j JsonObject) StringDefault(field, def string) string {
	val, ok := j.StringOk(field)
	if ok {
		return val
	}
	return def
}

func (j JsonObject) String(field string) string {
	val, _ := j.StringOk(field)
	return val
}

// Float64Ok returns a number from the json object, converted to float64 and a boolean which indicates, whether the
// result is ok.
func (j JsonObject) Float64Ok(field string) (float64, bool) {
	if val, ok := j[field]; ok {
		return convToFloat64(val)
	}
	return 0, false
}

func (j JsonObject) Float64Default(field string, def float64) float64 {
	val, ok := j.Float64Ok(field)
	if ok {
		return val
	}
	return def
}

func (j JsonObject) Float64(field string) float64 {
	val, _ := j.Float64Ok(field)
	return val
}

// Float32Ok returns a number from the json object, converted to float32 and a boolean which indicates, whether the
// result is ok.
func (j JsonObject) Float32Ok(field string) (float32, bool) {
	val, ok := j.Float64Ok(field)
	return float32(val), ok
}

func (j JsonObject) Float32Default(field string, def float32) float32 {
	val, ok := j.Float32Ok(field)
	if ok {
		return val
	}
	return def
}

func (j JsonObject) Float32(field string) float32 {
	val, _ := j.Float32Ok(field)
	return val
}

// IntOk returns a number from the json object, converted to int and a boolean which indicates, whether the
// result is ok.
func (j JsonObject) IntOk(field string) (int, bool) {
	val, ok := j.Float64Ok(field)
	return int(val), ok
}

func (j JsonObject) IntDefault(field string, def int) int {
	val, ok := j.IntOk(field)
	if ok {
		return val
	}
	return def
}

func (j JsonObject) Int(field string) int {
	val, _ := j.IntOk(field)
	return val
}

func (j JsonObject) BoolOk(field string) (bool, bool) {
	if val, ok := j[field]; ok {
		return convToBool(val)
	}
	return false, false
}

func (j JsonObject) BoolDefault(field string, def bool) bool {
	val, ok := j.BoolOk(field)
	if ok {
		return val
	}
	return def
}

func (j JsonObject) Bool(field string) bool {
	val, _ := j.BoolOk(field)
	return val
}

func (j JsonObject) SetObject(field string, value JsonObject) {
	j[field] = value
}

func (j JsonObject) SetList(field string, value JsonList) {
	j[field] = value
}

// SetNumber writes an integer or float into the json object.
// Json doesn't know types like "float64" or "int". Numbers are handled like "float64".
// The Get methods converts the number to the type you want, like "int".
func (j JsonObject) SetNumber(field string, value float64) {
	j[field] = value
}

func (j JsonObject) SetString(field, value string) {
	j[field] = value
}

func (j JsonObject) SetBool(field string, value bool) {
	j[field] = value
}

// Chain receives multiple field names, which represent a json object hierarchy.
// E.g. when you have an object, containing another object in field "a", which contains another object in field "b",
// you can use Chain("a", "b") to walk down the hierarchy.
// When somewhere the object does not have another object in the called field, the method will return nil.
func (j JsonObject) Chain(field ...string) JsonObject {
	var result = j

	for _, f := range field {
		result = result.Object(f)
		if result == nil {
			return nil
		}
	}

	return result
}
