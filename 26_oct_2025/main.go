package main

import "fmt"

type Bank struct {
	balance []int64
	n       int
}

func Constructor(balance []int64) Bank {
	account := append([]int64{0}, balance...)
	return Bank{balance: account, n: len(balance)}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > this.n || account2 > this.n || this.balance[account1] < money {
		return false
	}

	this.balance[account1] -= money
	this.balance[account2] += money

	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if this.n < account {
		return false
	}
	this.balance[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if this.n < account || this.balance[account] < money {
		return false
	}
	this.balance[account] -= money
	return true
}

func main() {
	balance := []int64{10, 100, 20, 50, 30}
	obj := Constructor(balance)

	param_3 := obj.Withdraw(3, 10)
	param_1 := obj.Transfer(5, 1, 20)
	param_2 := obj.Deposit(5, 20)
	param_4 := obj.Transfer(3, 4, 15)
	param_5 := obj.Withdraw(10, 50)
	fmt.Println(param_1, param_2, param_3, param_4, param_5)
}
