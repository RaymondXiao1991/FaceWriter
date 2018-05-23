package main

import (
	"net/http"
	"plot"
	"spinner"
	"time"
	"fmt"
	"net"
	"log"
	"clock"
	"regexp"
)

func main() {
	// http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)
	// http://localhost:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/16
	// http://localhost:8000/plot?expr=sin(x*y/10)/20
	http.HandleFunc("/plot", plot.Plot)

	go spinner.Spinner(100 * time.Millisecond)
	const n = 45
	fibN := spinner.Fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

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

	http.ListenAndServe("localhost:8001", nil)
}


package main

import (
"fmt"
"regexp"
)

func main() {
	str := "test"
	matched, err := regexp.MatchString("(wosaicrm)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]", str)
	fmt.Println(matched, err)
}
