package concurrent

import (
	"os"
	"sync"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func Du4() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDirWithCancel(root, &n, fileSizes)
	}
}

func walkDirWithCancel(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {

}
