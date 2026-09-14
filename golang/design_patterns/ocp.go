package designpatterns

import (
	"database/sql"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// open closed principle states that structs should be open for extensions but closed for modifications

type TransactionDB interface {
	Create(transaction Transaction) error
	Modify(transaction Transaction) error
}

type TransactionSQLRepository struct {
	db *sql.DB
}

func (tsr *TransactionSQLRepository) Create(transaction Transaction) error {
	fmt.Println("Using the SQL database to write")
	return nil
}

func (tsr *TransactionSQLRepository) Modify(transaction Transaction) error {
	fmt.Println("Using the SQL database to write")
	return nil
}

type TransactionMongoRepository struct {
	client *mongo.Client
}

func (tmr *TransactionMongoRepository) Create(transaction Transaction) error {
	fmt.Println("Using the mongo database to write")
	return nil
}

func (tmr *TransactionMongoRepository) Modify(transaction Transaction) error {
	fmt.Println("Using the mongo database to write")
	return nil
}