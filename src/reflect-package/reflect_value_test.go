package reflect_package

import (
	"testing"
	"reflect"
	"os"
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
}
