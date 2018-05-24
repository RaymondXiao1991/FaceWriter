package channel

import "fmt"

// Counter
func Counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

// Squarer
func Squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// Printer
func Printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
