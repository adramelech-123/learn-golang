package bnk

import (
	"errors"
	"fmt"
	"time"
)

type AuditInfo struct {
	CreatedAt    time.Time
	LastModified time.Time
}

type AccountNumber string

type BankAccount struct {
	AccountNumber AccountNumber
	Balance       float64  
	AuditInfo     
}

func (a AccountNumber) IsValid() bool {
	return len(string(a)) == 12
}

func NewBankAccount(acc AccountNumber) *BankAccount {
	// Creates the Account if the AccountNumber is Valid else we get a runtime error
	if acc.IsValid() {
		return &BankAccount{
			AccountNumber: acc,
			AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
		}
	}
	
	return nil
}

func (bal BankAccount) DisplayBalance() {
	fmt.Printf("Account Number: %s, Balance: $%2f\n", bal.AccountNumber, bal.Balance)
}


func (bal *BankAccount) Deposit(amt float64) {
	bal.Balance += amt
}


func (bal *BankAccount) Withdraw(amt float64) error {
	
	if bal.Balance < amt {
		fmt.Printf("Insufficient funds! Your balance is $%2f\n", bal.Balance)
		return errors.New("insufficient funds")
	}

	bal.Balance -= amt
	fmt.Printf("You have withdrawn $%2f, your current balance is %2f\n", amt, bal.Balance)

	return nil
}