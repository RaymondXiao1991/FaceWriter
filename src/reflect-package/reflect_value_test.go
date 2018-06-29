package reflect_package

import (
	"testing"
	"reflect"
	"os"
	"strings"
)

func TestReflectValue(t *testing.T) {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	t.Log(x)

	d.Set(reflect.ValueOf(4))
	t.Log(x)

	e := reflect.ValueOf(&x).Elem()
	e.SetInt(5)
	t.Log(x)

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	ry.Set(reflect.ValueOf(6))
	t.Log(y)
	ry.Set(reflect.ValueOf("hello"))
	t.Log(y)

	stdout := reflect.ValueOf(os.Stdout).Elem()
	fd := stdout.FieldByName("fd")
	t.Log(fd.CanAddr(), fd.CanSet())

	type Ab struct {
		a string
		b int64
	}
	var ab Ab
	xx := reflect.ValueOf(&ab).Elem()
	t.Log(xx)
}

func TestField(t *testing.T) {
	var ab struct {
		a string `http:"a"`
		b int    `http:"b"`
	}
	ab.b = 10

	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(&ab).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}
	t.Log(fields)
}
