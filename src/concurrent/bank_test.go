package concurrent

import (
	"testing"
	"fmt"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		Deposit(200.00)
		fmt.Println("balance =", Balance())
		done <- struct{}{}
	}()

	go func() {
		Deposit(100.00)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := Balance(), 300.00; got != want {
		t.Errorf("Balance = %.2f, want %.2f", got, want)
	}
}
