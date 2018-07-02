package concurrent

var deposits = make(chan float64)
var balances = make(chan float64)

func Deposit(amount float64) {
	deposits <- amount
}

func Balance() float64 {
	return <-balances
}

func teller() {
	var balance float64
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
