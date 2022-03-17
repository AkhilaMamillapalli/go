package main

import "fmt"

type Account interface {
	AccNumber() int
	Name() string
}

type Deposit struct {
	name       string
	balance    int
	depositAmt int
}

type Withdraw struct {
	name     string
	balance  int
	withdraw int
}

func (d Deposit) AccNumber() int {

	d.balance = d.balance + d.depositAmt
	return d.balance
}

func (d Deposit) Name() string {

	return d.name
}

func (w Withdraw) AccNumber() int {
	
	    if w.balance > 0 {
		w.balance = w.balance - w.withdraw

	}
	return w.balance

}
func (w Withdraw) Name() string {
	return w.name
}
func main() {

	a := Deposit{"Akhila", 10, 20}
	b := Withdraw{"Akhila", 40, 30}
	var ac Account
	ac = a
	fmt.Println(ac.Name())
	fmt.Println(ac.AccNumber())
	ac = b
	fmt.Println(ac.Name())
	fmt.Println(ac.AccNumber())

}
