package bnk

import (
	"fmt"
	"time"
)



type Customer struct {
	Name      string
	accounts  []BankAccount
	AuditInfo
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name: name,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}
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

