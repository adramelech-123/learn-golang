package main

import (
	"errors"
	"fmt"
	"time"
)

type AuditInfo struct {
	CreatedAt time.Time
	LastModified time.Time
}

type BankAccount struct {
	AccountNumber string
	Balance       float64  
	AuditInfo     
}

func (bal BankAccount) DisplayBalance() {
	fmt.Printf("Account Number: %s, Balance: $%2f\n", bal.AccountNumber, bal.Balance)
}

func (bal *BankAccount) Deposit(amt float64) {
	bal.Balance += amt
}

func (bal *BankAccount) Withdraw(amt float64) error {
	
	if bal.Balance < amt {
		fmt.Printf("You have Insufficient funds! Your balance is $%2f\n", bal.Balance)
		return errors.New("insufficient funds")
	}

	bal.Balance -= amt
	fmt.Printf("You have withdrawn $%2f, your current balance is %2f\n", amt, bal.Balance)

	return nil
}

type Customer struct {
	Name      string
	accounts  []BankAccount
	AuditInfo
}

// Add New account to customer
func (c *Customer) AddAccount(account BankAccount) {
	c.accounts = append(c.accounts, account)
}

// Display all acccounts of a customer
func (c Customer) DisplayAccounts() {
	fmt.Printf("Customer: %s\n", c.Name)
	for _, account := range c.accounts {
		account.DisplayBalance()
	}
}

type Bank struct {
	Name string
	Customers []Customer
} 

// Add Customer to Bank
func (b *Bank) AddCustomer (customer Customer) {
	b.Customers = append(b.Customers, customer)
}

// Display all customers in the bank
func (b Bank) DisplayCustomers() {
	fmt.Printf("Bank: %s\n", b.Name)
	for _, customer := range b.Customers {
		customer.DisplayAccounts()
	}
}


// Constructor Functions

func NewBank(name string) *Bank {
	return &Bank{Name: name}
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name: name,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}
}

func NewBankAccount(acc string) *BankAccount {
	return &BankAccount{
		AccountNumber: acc,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}
}

func main() {
	bank1 := NewBank("Hidden Leaf Bank")
	bank2 := NewBank("Broke People's Bank")

	customer1 := NewCustomer("Sasuke Uchiha")
	customer2 := NewCustomer("Naruto Uzumaki")
	customer3 := NewCustomer("Johhny Bravo")

	account1 := NewBankAccount("178001618844")
	account1.Deposit(3600)

	account2 := NewBankAccount("1780023418756")
	account2.Deposit(6850)

	account3 := NewBankAccount("1571059317788")
	account3.Deposit(13.5)


	// Adding accounts to customers
	customer1.AddAccount(*account1)
	customer2.AddAccount(*account2)
	customer3.AddAccount(*account3)

	// Adding Customers to bank
	bank1.AddCustomer(*customer1)
	bank1.AddCustomer(*customer2)
	bank2.AddCustomer(*customer3)

	bank1.DisplayCustomers()
	bank2.DisplayCustomers()


}