package designpatterns

// Adding the anti-srp implementation code

type struct Transaction {
	Id string
	Approved bool
	Source string
	Destination string
}


func (t *Transaction) Initiate() {
	// business logic
}

func (t *Transaction) Revert() {
	// business logic
}

func (t *Transaction) Save() {
	// business logic
}

func (t *Transaction) Update() {
	// business logic
}