package designpatterns

import (
	"testing"
	"time"
)

func TestSRPExecution(t *testing.T) {
	var tr TransactionOperations = &TransactionRepository{}
	var ts TransactionService = TransactionService{
		TransactionOperations: tr,
	}
	transaction := Transaction {
		Id: "123",
		Sender: "pganguli@okaxis",
		Receiver: "ankita@okaxis",
		TransactionStatus: Pending,
		CreatedOn: time.Now(),
	}
	ts.InitiateTransaction(transaction)
	transaction.TransactionStatus = Failure
	ts.RevertTransaction(transaction)
}