package main

import (
	"fmt"

	"github.com/kundefined97/leango/mydict"
)

func main() {
	// accounts
	// account := accounts.NewAccount("kunde")
	// account.Deposit(10)
	// account.ChangeOwner("kan")
	// fmt.Println(account)

	// mydict
	dictionary := mydict.Dictionary{"first": "First Word"}
	// Add
	dictionary.Add("second", "Second Word")
	fmt.Println(dictionary.Search("second"))
	// Update

	// Delete
}
