package bnk

import "fmt"

type Bank struct {
	Name      string
	Customers []Customer
}

func NewBank(name string) *Bank {
	return &Bank{Name: name}
}

// Add Customer to Bank
func (b *Bank) AddCustomer(customer Customer) {
	b.Customers = append(b.Customers, customer)
}

// Display all customers in the bank
func (b Bank) DisplayCustomers() {
	fmt.Printf("Bank: %s\n", b.Name)
	for _, customer := range b.Customers {
		customer.DisplayAccounts()
	}
}