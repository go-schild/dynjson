package dynjson_test

import (
	"testing"

	"github.com/go-schild/dynjson"
	"github.com/stretchr/testify/assert"
)

func TestNewJsonObject(t *testing.T) {
	a := dynjson.NewJsonObject()

	assert.NotNil(t, a)
}

func TestParseObject(t *testing.T) {
	const testData = `{"string": "hello", "bool": true, "int": 5}`
	j, _ := dynjson.ParseObject(testData)

	assert.Equal(t, j.String("string"), "hello")
	assert.Equal(t, j.Int("int"), 5)
}

func TestJsonObject_Has(t *testing.T) {
	const testData = `{"string": "hello", "bool": true, "int": 5}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, true, j.Has("string"))
	assert.Equal(t, true, j.Has("bool"))
	assert.Equal(t, true, j.Has("int"))
	assert.Equal(t, false, j.Has("none"))
	assert.Equal(t, false, j.Has(""))
}

func TestJsonObject_ObjectOk(t *testing.T) {
	const testData = `{"data": {"even_more_data": 5}}`

	j1, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	j2, ok := j1.ObjectOk("data")
	assert.True(t, ok)
	assert.NotNil(t, j2)

	a := j2.Int("even_more_data")
	assert.Equal(t, 5, a)
}

func TestJsonObject_Object(t *testing.T) {
	const testData = `{"data": {"even_more_data": 5}}`

	j1, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	j2 := j1.Object("data")
	assert.NotNil(t, j2)

	a := j2.Int("even_more_data")
	assert.Equal(t, 5, a)
}

func TestJsonObject_ListOk(t *testing.T) {
	const testData = `{"data": [0, 1, 2]}`

	j1, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	j2, ok := j1.ListOk("data")
	assert.True(t, ok)
	assert.NotNil(t, j2)

	for index, item := range j2 {
		assert.Equal(t, index, item.Int())
	}
}

func TestJsonObject_List(t *testing.T) {
	const testData = `{"data": [0, 1, 2]}`

	j1, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	j2 := j1.List("data")
	assert.NotNil(t, j2)

	for index, item := range j2 {
		assert.Equal(t, index, item.Int())
	}
}

func TestJsonObject_StringOk(t *testing.T) {
	const testData = `{"a": "Hello", "b": "Gopher", "c": false}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	a, ok := j.StringOk("a")
	assert.True(t, ok)
	assert.Equal(t, "Hello", a)

	a, ok = j.StringOk("b")
	assert.True(t, ok)
	assert.Equal(t, "Gopher", a)

	a, ok = j.StringOk("c")
	assert.False(t, ok)
	assert.Equal(t, "", a)
}

func TestJsonObject_StringDefault(t *testing.T) {
	const testData = `{"a": "Hello", "b": "Gopher", "c": false}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, "Hello", j.StringDefault("a", "default"))
	assert.Equal(t, "Gopher", j.StringDefault("b", "default"))
	assert.Equal(t, "default", j.StringDefault("c", "default"))
}

func TestJsonObject_String(t *testing.T) {
	const testData = `{"a": "Hello", "b": "Gopher", "c": false}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, "Hello", j.String("a"))
	assert.Equal(t, "Gopher", j.String("b"))
	assert.Equal(t, "", j.String("c"))
}

func TestJsonObject_Float64Ok(t *testing.T) {
	const testData = `{"a": 3.14, "b": 1337, "c": false}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	a, ok := j.Float64Ok("a")
	assert.True(t, ok)
	assert.Equal(t, float64(3.14), a)

	a, ok = j.Float64Ok("b")
	assert.True(t, ok)
	assert.Equal(t, float64(1337), a)

	a, ok = j.Float64Ok("c")
	assert.False(t, ok)
	assert.Equal(t, float64(0), a)
}

func TestJsonObject_Float64Default(t *testing.T) {
	const testData = `{"a": 1.5, "b": 15.25, "c": "none"}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, float64(1.5), j.Float64Default("a", -1))
	assert.Equal(t, float64(15.25), j.Float64Default("b", -1))
	assert.Equal(t, float64(-1), j.Float64Default("c", -1))
}

func TestJsonObject_Float64(t *testing.T) {
	const testData = `{"a": 1.5, "b": 15.25, "c": "none"}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, float64(1.5), j.Float64("a"))
	assert.Equal(t, float64(15.25), j.Float64("b"))
	assert.Equal(t, float64(0), j.Float64("c"))
}

func TestJsonObject_Float32Ok(t *testing.T) {
	const testData = `{"a": 3.14, "b": 1337, "c": false}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	a, ok := j.Float32Ok("a")
	assert.True(t, ok)
	assert.Equal(t, float32(3.14), a)

	a, ok = j.Float32Ok("b")
	assert.True(t, ok)
	assert.Equal(t, float32(1337), a)

	a, ok = j.Float32Ok("c")
	assert.False(t, ok)
	assert.Equal(t, float32(0), a)
}

func TestJsonObject_Float32Default(t *testing.T) {
	const testData = `{"a": 1.5, "b": 15.25, "c": "none"}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, float32(1.5), j.Float32Default("a", -1))
	assert.Equal(t, float32(15.25), j.Float32Default("b", -1))
	assert.Equal(t, float32(-1), j.Float32Default("c", -1))
}

func TestJsonObject_Float32(t *testing.T) {
	const testData = `{"a": 1.5, "b": 15.25, "c": "none"}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, float32(1.5), j.Float32("a"))
	assert.Equal(t, float32(15.25), j.Float32("b"))
	assert.Equal(t, float32(0), j.Float32("c"))
}

func TestJsonObject_IntOk(t *testing.T) {
	const testData = `{"a": 3, "b": 44, "c": false}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	a, ok := j.IntOk("a")
	assert.True(t, ok)
	assert.Equal(t, 3, a)

	a, ok = j.IntOk("b")
	assert.True(t, ok)
	assert.Equal(t, 44, a)

	a, ok = j.IntOk("c")
	assert.False(t, ok)
	assert.Equal(t, 0, a)
}

func TestJsonObject_IntDefault(t *testing.T) {
	const testData = `{"a": 1, "b": 15, "c": "none"}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, 1, j.IntDefault("a", -1))
	assert.Equal(t, 15, j.IntDefault("b", -1))
	assert.Equal(t, -1, j.IntDefault("c", -1))
}

func TestJsonObject_Int(t *testing.T) {
	const testData = `{"a": 1, "b": 15, "c": "none"}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, 1, j.Int("a"))
	assert.Equal(t, 15, j.Int("b"))
	assert.Equal(t, 0, j.Int("c"))
}

func TestJsonObject_SetObject(t *testing.T) {
	o1 := dynjson.NewJsonObject()
	o2 := dynjson.NewJsonObject()

	o2.SetNumber("a", 5)
	o1.SetObject("o2", o2)

	assert.Equal(t, 5, o1.Object("o2").Int("a"))
}

func TestJsonObject_SetList(t *testing.T) {
	o1 := dynjson.NewJsonObject()
	o2 := dynjson.NewJsonList(nil)

	o2.Append(5)
	o1.SetList("o2", o2)

	l, ok := o1.ListOk("o2")
	assert.True(t, ok)
	assert.NotNil(t, l)
	assert.Equal(t, 5, l[0].Int())
}

func TestJsonObject_SetNumber(t *testing.T) {
	const testData = `{"a": 5}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, 5, j.Int("a"))
	j.SetNumber("a", 10)
	assert.Equal(t, 10, j.Int("a"))
}

func TestJsonObject_SetString(t *testing.T) {
	const testData = `{"a": "Hello"}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, "Hello", j.String("a"))
	j.SetString("a", "World")
	assert.Equal(t, "World", j.String("a"))
}

func TestJsonObject_SetBool(t *testing.T) {
	const testData = `{"a": true}`

	j, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	assert.Equal(t, true, j.Bool("a"))
	j.SetBool("a", false)
	assert.Equal(t, false, j.Bool("a"))
}

func TestJsonObject_Chain(t *testing.T) {
	const testData = `{"a": {"b": {"c": 5}}}`

	j1, err := dynjson.ParseObject(testData)
	assert.Nil(t, err)

	j2 := j1.Chain("a", "b")
	assert.NotNil(t, j2)
	assert.Equal(t, 5, j2.Int("c"))

	j3 := j1.Chain("a", "b", "c") // c is not an object
	assert.Nil(t, j3)
}
