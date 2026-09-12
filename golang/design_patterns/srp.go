package designpatterns

import "database/sql"

// Adding the SRP implementation code

type TransactionStore interface {
	Save(transaction *Transaction) error
	Update(transaction *Transaction) error
}

type Transaction struct {
	Id string
	Approved bool
	Source string
	Destination string
}

type TransactionRepository struct {
}

func (t *Transaction) Initiate() {
	// business logic
	return
}

func (t *Transaction) Revert() {
	// business logic
	return
}

func (t *TransactionRepository) Save(transaction *Transaction) error {
	// business logic
	return nil
}

func (t *TransactionRepository) Update(transaction *Transaction) error {
	// business logic
	return nil
}
