package designpatterns

import (
	"testing"
	"fmt"
)

func TestGetConfigInstance(t *testing.T) {
	i1 := GetConfigInstance()
	fmt.Println(i1)
	i2 := GetConfigInstance()
	fmt.Println(i2)
}