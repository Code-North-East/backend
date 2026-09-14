package designpatterns

type CardOperations interface {
	Buy() error
	Withdraw() error
}