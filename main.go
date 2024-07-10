package main

import (
	"gocourse/bnk"
)

func main() {
	bank := bnk.NewBank("Golang Bank")

	customer1 := bnk.NewCustomer("Sasuke Uchiha")
	customer2 := bnk.NewCustomer("Naruto Uzumaki")

	account1 := bnk.NewBankAccount("178001618844")
	account1.Deposit(3600)

	account2 := bnk.NewBankAccount("178002341875")
	account2.Deposit(6850)

	// Adding accounts to customers
	customer1.AddAccount(*account1)
	customer2.AddAccount(*account2)

	// Adding Customers to bank
	bank.AddCustomer(*customer1)
	bank.AddCustomer(*customer2)

	bank.DisplayCustomers()
}