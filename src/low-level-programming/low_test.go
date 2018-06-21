package low_level_programming

import (
	"testing"
	"reflect"
	"unsafe"
	"fmt"
	"strings"
	"bytes"
)

/*
  类型                                大小
  bool                              1个字节
  intN, uintN, floatN, complexN     N/8个字节(例如float64是8个字节)
  int, uint, uintptr                1个机器字
  *T                                1个机器字
  string                            2个机器字(data,len)
  []T                               3个机器字(data,len,cap)
  map                               1个机器字
  func                              1个机器字
  chan                              1个机器字
  interface                         2个机器字(type,value)
 */
func TestUnsafeSizeof(t *testing.T) {
	var (
		b bool
		u uint16
		i int64
		f float64
	)

	t.Log(reflect.TypeOf(unsafe.Sizeof(b)), unsafe.Sizeof(b))
	t.Log(reflect.TypeOf(unsafe.Sizeof(u)), unsafe.Sizeof(u))
	t.Log(reflect.TypeOf(unsafe.Sizeof(i)), unsafe.Sizeof(i))
	t.Log(reflect.TypeOf(unsafe.Sizeof(f)), unsafe.Sizeof(f))

	var x struct {
		a bool
		b int16
		c []int
	}
	t.Log(reflect.TypeOf(unsafe.Sizeof(x)), unsafe.Alignof(x))
	t.Log(reflect.TypeOf(unsafe.Sizeof(x.a)), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	t.Log(reflect.TypeOf(unsafe.Sizeof(x.b)), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	t.Log(reflect.TypeOf(unsafe.Sizeof(x.c)), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))
}

func TestUnsafePointer(t *testing.T) {
	fmt.Printf("%#016x\n", UnsafePointer(1.0)) // "0x3ff0000000000000"
	var x struct {
		a bool
		b int16
		c []int
	}
	pa := (*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.a)))
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pa = true
	*pb = 42
	t.Log(x.a)
	t.Log(x.b)

	t.Log(reflect.ValueOf(map[string][]int{}).Kind())
	t.Log(reflect.ValueOf(map[string][]int(nil)).Kind())

	mi1, mi2 := map[string][]int{}, map[string][]int(nil)
	vx, vy := reflect.ValueOf(mi1), reflect.ValueOf(mi2)
	//kx, ky := reflect.ValueOf(mi1).Kind(), reflect.ValueOf(mi2).Kind()
	t.Log(vx.Len())
	t.Log(vy.Len())
	t.Logf("vx.MapKeys():%v", vx.MapKeys())
	t.Logf("vy.MapKeys():%v", vy.MapKeys())
}

func TestSplit(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Log("Not equal")
	} else {
		t.Log("Equal")
	}
}

func TestDeepEqual(t *testing.T) {
	var a, b []string = nil, []string{}
	t.Log(reflect.DeepEqual(a, b)) // "false"

	var c, d map[string]int = nil, make(map[string]int)
	t.Log(reflect.DeepEqual(c, d)) // "false"
}

func TestEqual(t *testing.T) {
	one, oneAgain, two := 1, 1, 2

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1, ch2 := make(chan int), make(chan int)
	var ch1ro <-chan int = ch1

	type mystring string

	var iface1, iface1Again, iface2 interface{} = &one, &oneAgain, &two

	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		{1, 1, true},    // same values
		{1, 2, false},   // different values
		{1, 1.0, false}, //different types
		{"foo", "foo", true},
		{"foo", "bar", false},
		{mystring("foo"), "foo", false}, // different types
		// slice
		{[]string{"foo"}, []string{"foo"}, true},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string{}, []string(nil), true},
		//slice cycles
		{cycleSlice, cycleSlice, true},
		// struct cycles
		{a, a, true},
		{b, b, true},
		{c, c, true},
		{a, b, false},
		{a, c, false},
		// maps
		{map[string][]int{"foo": {1, 2, 3}}, map[string][]int{"foo": {1, 2, 3}}, true},
		{map[string][]int{"foo": {1, 2, 3}}, map[string][]int{"foo": {1, 2, 3, 4}}, false},
		{map[string][]int{}, map[string][]int(nil), true}, // 切片的index为0时panic
		// pointers
		{&one, &one, true},
		{&one, &two, false},
		{&one, &oneAgain, true},
		{&one, &oneAgain, true},
		{new(bytes.Buffer), new(bytes.Buffer), true},
		// pointer cycles
		{cyclePtr1, cyclePtr1, true},
		{cyclePtr2, cyclePtr2, true},
		{cyclePtr1, cyclePtr2, true}, // deeply equal
		// functions
		{(func())(nil), (func())(nil), true},
		{(func())(nil), func() {}, false},
		{func() {}, func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, [...]int{1, 2, 3}, true},
		{[...]int{1, 2, 3}, [...]int{1, 2, 4}, false},
		// channels
		{ch1, ch1, true},
		{ch1, ch2, false},
		{ch1ro, ch1, false},
		// interface
		{&iface1, &iface1, true},
		{&iface1, iface2, false},
		{&iface1Again, iface1, false}, // ??????????
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t", test.x, test.y, !test.want)
			/*} else {
				t.Logf("Equal(%v, %v) = %t", test.x, test.y, test.want)*/
		}
	}
}
