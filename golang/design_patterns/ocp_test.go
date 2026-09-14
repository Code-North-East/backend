package designpatterns

import (
	"testing"
	"time"
)

func TestOCPInvocati(t *testing.T) {
	var tdb TransactionDB = &TransactionMongoRepository{}
	transaction := Transaction {
		Id: "123",
		Sender: "pganguli@okaxis",
		Receiver: "ankita@okaxis",
		TransactionStatus: Pending,
		CreatedOn: time.Now(),
	}
	tdb.Create(transaction)
	tdb = &TransactionSQLRepository{}
	tdb.Modify(transaction)
}