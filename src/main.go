package main

import (
	"net/http"
	"reflect-package"
	"log"
	"select-multiplex"
)

func main() {
	// http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)
	// http://localhost:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/16
	// http://localhost:8000/plot?expr=sin(x*y/10)/20
	/*
	http.HandleFunc("/plot", plot.Plot)
	*/

	/*
	go spinner.Spinner(100 * time.Millisecond)
	const n = 45
	fibN := spinner.Fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	*/

	/*
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		//clock.HandleConn(conn) // handle one connection at a time
		go clock.HandleConn(conn) // handle connections concurrently
	}
	*/

	/*
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go echo.HandleConn(conn)
	}
	*/

	/*naturals := make(chan int)
	squares := make(chan int)
	go channel.Counter(naturals)
	go channel.Squarer(squares, naturals)
	channel.Printer(squares)*/

	//http.ListenAndServe("localhost:8002", nil)
	select_multiplex.SelectMultiplexWithAbort()
	http.HandleFunc("/search", reflect_package.Search)
	log.Fatal(http.ListenAndServe(":23451", nil))
}
