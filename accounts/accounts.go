package accounts

import (
	"errors"
	"fmt"
)

type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("cannot withdraw")

func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// (a *account) is pointer reciever
func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil

}

func (a *account) Balance() int {
	return a.balance
}

func (a *account) Owner() string {
	return a.owner
}

func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a *account) String() string {
	return fmt.Sprint(a.owner, "'s account. \nHas: ", a.balance)
}
