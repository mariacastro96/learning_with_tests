package wallet

import "errors"

// Wallet has a balance
type Wallet struct {
	balance int
}

// Deposit will add to the Wallet balance
func (wallet *Wallet) Deposit(value int) {
	wallet.balance += value
	return
}

// Withdraw will remove from the wallet balance
func (wallet *Wallet) Withdraw(value int) error {
	if wallet.balance < value {
		return errors.New("nope")
	}
	wallet.balance -= value
	return nil
}

// Balance will show the final balance
func (wallet *Wallet) Balance() int {
	return wallet.balance
}
