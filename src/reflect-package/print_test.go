package reflect_package

import (
	"testing"
	"time"
	"strings"
	"fmt"
)

func TestPrint(t *testing.T) {
	Print(time.Hour)
	fmt.Println("------------------------------------")
	Print(new(strings.Replacer))
}
