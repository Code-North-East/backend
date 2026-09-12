package designpatterns

// Adding the anti-srp implementation code

type FalseTransaction struct {
	Id string
	Approved bool
	Source string
	Destination string
}


func (t *FalseTransactionn) Initiate() {
	// business logic
}

func (t *FalseTransactionn) Revert() {
	// business logic
}

func (t *FalseTransactionn) Save() {
	// business logic
}

func (t *FalseTransactionn) Update() {
	// business logic
}