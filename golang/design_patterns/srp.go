package designpatterns

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
	db *sql.DB
}

func (t *Transaction) Initiate() {
	// business logic
}

func (t *Transaction) Revert() {
	// business logic
}

func (t *TransactionRepository) Save(transaction *Transaction) error {
	// business logic
	return nil
}

func (t *TransactionRepository) Update(transaction *Transaction) error {
	// business logic
	return nil
}
