package designpatterns

import "time"

type TransactionStatus string

type DBClient string 

const (
	Postgres DBClient = "Postgres"
	MongoDB DBClient = "MongoDB"
	MySQL DBClient = "MySQL"
)

const (
	Pending TransactionStatus = "Pending"
	Completed TransactionStatus = "Completed"
	Failed TransactionStatus = "Failed"
)

type TransactionOps interface {
	Initiate (t *Transaction) error
	Revert (t *Transaction) error
}

type Transaction struct {
	Id string
	Sender string
	Receiver string
	CreatedOn time.Time
	Status TransactionStatus
}	

type TransactionRepository struct {
	DbClient DBClient

}

func (tr *TransactionRepository) Initiate(t *Transaction) error {
	return nil
}

func (tr *TransactionRepository) Revert(t *Transaction) error {
	return nil
}