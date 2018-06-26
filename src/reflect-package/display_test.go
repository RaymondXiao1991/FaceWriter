package reflect_package

import (
	"testing"
	"os"
	"io"
	"sync"
	"net"
	"reflect"
)

func TestDisplay(t *testing.T) {
	Display("os.Stderr", os.Stderr)

	var w io.Writer = os.Stderr
	Display("&w", &w)

	var locker sync.Locker = new(sync.Mutex)
	Display("&locker", &locker)

	Display("locker", locker)

	locker = nil
	Display("(&locker)", &locker)

	ips, _ := net.LookupHost("golang.org")
	Display("ips", ips)

	Display("rV", reflect.ValueOf(os.Stderr))

	// a pointer that points to itself
	type P *P
	var p P
	p = &p
	if false {
		Display("p", p)
	}

	// a map that contains itself
	type M map[string]M
	m := make(M)
	m[""] = m
	if false {
		Display("m", m)
	}

	// a slice that contains itself
	type S []S
	s := make(S, 1)
	s[0] = s
	if false {
		Display("s", s)
	}

	// a linked list that eats its own tail
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	if false {
		Display("c", c)
	}
}
