package main

import (
	"fmt"

	"github.com/bewisesh91/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("bewisesh")
	account.Deposit(100000)
	err := account.Withdraw(1000000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
