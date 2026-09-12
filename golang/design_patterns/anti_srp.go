package design_patterns

// Adding the anti-srp implementation code

type FalseTransaction struct {
	Id string
	Approved bool
	Source string
	Destination string
}


func (t *FalseTransaction) Initiate() {
	// business logic
}

func (t *FalseTransaction) Revert() {
	// business logic
}

func (t *FalseTransaction) Save() {
	// business logic
}

func (t *FalseTransaction) Update() {
	// business logic
}