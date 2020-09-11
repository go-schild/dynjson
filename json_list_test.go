package dynjson_test

import (
	"testing"

	"github.com/go-schild/dynjson"
	"github.com/stretchr/testify/assert"
)

func TestNewJsonList(t *testing.T) {
	a := dynjson.NewJsonList(nil)
	b := dynjson.NewJsonList(dynjson.JsonListRaw{1, 2, "Hello"})

	assert.NotNil(t, a)
	assert.NotNil(t, b)

	assert.Equal(t, 0, len(a))
	assert.Equal(t, 3, len(b))
}

func TestParseList(t *testing.T) {
	const testData = `[1, 2, 3]`
	j, err := dynjson.ParseList(testData)

	assert.NotNil(t, j)
	assert.Nil(t, err)
}

func TestJsonList_Append(t *testing.T) {
	j := dynjson.NewJsonList(nil)
	j.Append(0, 1)
	j.Append(2, 3)
	j.Append(4, 5)

	for index, item := range j {
		assert.Equal(t, index, item.Int())
	}
}

func TestJsonList_Prepend(t *testing.T) {
	j := dynjson.NewJsonList(nil)
	j.Prepend(4, 5)
	j.Prepend(2, 3)
	j.Prepend(0, 1)

	for index, item := range j {
		assert.Equal(t, index, item.Int())
	}
}

func TestJsonList_ObjectOk(t *testing.T) {
	const testData = `[{"val": 0}, {"val": 1}]`

	j1, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, j2 := range j1 {
		o, ok := j2.ObjectOk()

		assert.True(t, ok)
		assert.Equal(t, index, o.Int("val"))
	}
}

func TestJsonList_Object(t *testing.T) {
	const testData = `[{"val": 0}, {"val": 1}]`

	j1, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, j2 := range j1 {
		o := j2.Object()

		assert.Equal(t, index, o.Int("val"))
	}
}

func TestJsonList_ListOk(t *testing.T) {
	const testData = `[[0, 1], [10, 11]]`

	j1, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index1, item1 := range j1 {
		j2, ok := item1.ListOk()

		assert.True(t, ok)
		assert.NotNil(t, j2)

		for index2, item2 := range j2 {
			shouldBe := index1*10 + index2
			assert.Equal(t, shouldBe, item2.Int())
		}
	}
}

func TestJsonList_List(t *testing.T) {
	const testData = `[[0, 1], [10, 11]]`

	j1, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index1, item1 := range j1 {
		j2 := item1.List()
		assert.NotNil(t, j2)

		for index2, item2 := range j2 {
			shouldBe := index1*10 + index2
			assert.Equal(t, shouldBe, item2.Int())
		}
	}
}

func TestJsonList_StringOk(t *testing.T) {
	const testData = `[1, "Hello", "World"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a, ok := item.StringOk()
			assert.Equal(t, "", a)
			assert.False(t, ok)
		case 1:
			a, ok := item.StringOk()
			assert.Equal(t, "Hello", a)
			assert.True(t, ok)
		case 2:
			a, ok := item.StringOk()
			assert.Equal(t, "World", a)
			assert.True(t, ok)
		}
	}
}

func TestJsonList_StringDefault(t *testing.T) {
	const testData = `[1, "Hello", "World"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a := item.StringDefault("default")
			assert.Equal(t, "default", a)
		case 1:
			a := item.StringDefault("default")
			assert.Equal(t, "Hello", a)
		case 2:
			a := item.StringDefault("default")
			assert.Equal(t, "World", a)
		}
	}
}

func TestJsonList_String(t *testing.T) {
	const testData = `["Hello", true, "World"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a := item.String()
			assert.Equal(t, "Hello", a)
		case 1:
			a := item.String()
			assert.Equal(t, "", a)
		case 2:
			a := item.String()
			assert.Equal(t, "World", a)
		}
	}
}

func TestJsonList_Float64Ok(t *testing.T) {
	const testData = `[1, 2, "Hello"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a, ok := item.Float64Ok()
			assert.Equal(t, float64(1), a)
			assert.True(t, ok)
		case 1:
			a, ok := item.Float64Ok()
			assert.Equal(t, float64(2), a)
			assert.True(t, ok)
		case 2:
			a, ok := item.Float64Ok()
			assert.Equal(t, float64(0), a)
			assert.False(t, ok)
		}
	}
}

func TestJsonList_Float64Default(t *testing.T) {
	const testData = `[1, 2, "Hello"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a := item.Float64Default(-1)
			assert.Equal(t, float64(1), a)
		case 1:
			a := item.Float64Default(-1)
			assert.Equal(t, float64(2), a)
		case 2:
			a := item.Float64Default(-1)
			assert.Equal(t, float64(-1), a)
		}
	}
}

func TestJsonList_Float64(t *testing.T) {
	const testData = `[1, 2, 3]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		assert.Equal(t, float64(index+1), item.Float64())
	}
}

func TestJsonList_Float32Ok(t *testing.T) {
	const testData = `[1, 2, "Hello"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a, ok := item.Float32Ok()
			assert.Equal(t, float32(1), a)
			assert.True(t, ok)
		case 1:
			a, ok := item.Float32Ok()
			assert.Equal(t, float32(2), a)
			assert.True(t, ok)
		case 2:
			a, ok := item.Float32Ok()
			assert.Equal(t, float32(0), a)
			assert.False(t, ok)
		}
	}
}

func TestJsonList_Float32Default(t *testing.T) {
	const testData = `[1, 2, "Hello"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a := item.Float32Default(-1)
			assert.Equal(t, float32(1), a)
		case 1:
			a := item.Float32Default(-1)
			assert.Equal(t, float32(2), a)
		case 2:
			a := item.Float32Default(-1)
			assert.Equal(t, float32(-1), a)
		}
	}
}

func TestJsonList_Float32(t *testing.T) {
	const testData = `[1, 2, 3]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		assert.Equal(t, float32(index+1), item.Float32())
	}
}

func TestJsonList_IntOk(t *testing.T) {
	const testData = `[1, 2, "Hello"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a, ok := item.IntOk()
			assert.Equal(t, 1, a)
			assert.True(t, ok)
		case 1:
			a, ok := item.IntOk()
			assert.Equal(t, 2, a)
			assert.True(t, ok)
		case 2:
			a, ok := item.IntOk()
			assert.Equal(t, 0, a)
			assert.False(t, ok)
		}
	}
}

func TestJsonList_IntDefault(t *testing.T) {
	const testData = `[1, 2, "Hello"]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		switch index {
		case 0:
			a := item.IntDefault(-1)
			assert.Equal(t, 1, a)
		case 1:
			a := item.IntDefault(-1)
			assert.Equal(t, 2, a)
		case 2:
			a := item.IntDefault(-1)
			assert.Equal(t, -1, a)
		}
	}
}

func TestJsonList_Int(t *testing.T) {
	const testData = `[1, 2, 3]`

	j, err := dynjson.ParseList(testData)
	assert.Nil(t, err)

	for index, item := range j {
		assert.Equal(t, index+1, item.Int())
	}
}
