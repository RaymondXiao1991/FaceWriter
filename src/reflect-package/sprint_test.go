package reflect_package

import (
	"time"
	"testing"
)

func TestAny(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	t.Log(Any(x))                  // "1"
	t.Log(Any(d))                  // "1"
	t.Log(Any([]int64{x}))         // "[]int64 0x8202b87b0"
	t.Log(Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
}
