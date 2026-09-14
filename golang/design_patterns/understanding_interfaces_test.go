package designpatterns

import (
	"fmt"
	"testing"
)

func TestInvocation(t *testing.T) {
	var userOps UserOperations = &UserRepository{}
	u := User{
		Id: "1001",
		Name: "John Doe",
		Email: "john.doe@mail.com",
		Phone: "+1-222-333-4444",
	}
	user := userOps.Create(u)
	fmt.Println(user)
}