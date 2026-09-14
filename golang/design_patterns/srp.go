package designpatterns

import (
	"database/sql"
	"fmt"
	"time"
)

type TransactionStatus string

const (
	Pending TransactionStatus = "PENDING"
	Success TransactionStatus = "SUCCESS"
	Failure TransactionStatus = "FAILURE"
)

type Transaction struct {
	Id string
	Sender string
	Receiver string
	TransactionStatus TransactionStatus
	CreatedOn time.Time
}

type TransactionOperations interface {
	Save(transaction Transaction) error
	Update(transaction Transaction) error
}

type TransactionRepository struct {
	db *sql.DB
}

func (tr *TransactionRepository) Save(transaction Transaction) error {
	fmt.Println("Saved transaction")
	return nil
}

func (tr *TransactionRepository) Update(transaction Transaction) error {
	fmt.Println("Updated transaction")
	return nil
}

type TransactionService struct {
	TransactionOperations TransactionOperations
}

func (ts *TransactionService) InitiateTransaction(transaction Transaction) error {
	fmt.Printf("Initiating transaction: %s\n", transaction)
	if err := ts.TransactionOperations.Save(transaction); err != nil {
		return err
	}
	return nil
}

func (ts *TransactionService) RevertTransaction(transaction Transaction) error {
	fmt.Printf("Reverting transaction: %+v\n", transaction)
	if err := ts.TransactionOperations.Update(transaction); err != nil {
		return err
	}
	return nil
}


