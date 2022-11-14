package dictionary

import (
	"fmt"

	"github.com/bewisesh91/learngo/accounts"
	"github.com/bewisesh91/learngo/mydict"
)

func dictionary() {
	account := accounts.NewAccount("bewisesh")
	account.Deposit(100000)
	err := account.Withdraw(1000000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)

	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"

	result := dictionary.Add(word, definition)
	if result != nil {
		fmt.Println(result)
	}

	hello, result := dictionary.Search(word)
	fmt.Println(hello)

	result2 := dictionary.Add(word, definition)
	if result2 != nil {
		fmt.Println(result2)
	}

	updatedResult := dictionary.Update(word, "인사")
	if updatedResult != nil {
		fmt.Print(updatedResult)
	}
	resultWord, _ := dictionary.Search(word)
	fmt.Println(resultWord)

	baseWord := "Test"
	dictionary.Add(baseWord, "Working")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)

	deletedWord, deleteError := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(deleteError)
	} else {
		fmt.Println(deletedWord)
	}
}
