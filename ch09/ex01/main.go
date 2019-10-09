package main

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	result := <-withdraws
	return result > 0
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			if amount > balance {
				withdraws <- -1
			} else {
				balance -= amount
				withdraws <- amount
			}
		}
	}
}

func init() {
	go teller()
}

func main() {}
