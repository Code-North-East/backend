package designpatterns

import (
	"sync"
	"fmt"
)

type DummyConfig struct {
	Port                int
	Host                string
	EncoddedEncyptedKey string
}

var (
	instance *DummyConfig
	once sync.Once
)

func GetConfigInstance() *DummyConfig {
	once.Do(func() {
		// check if instance is nil
		if instance == nil {
			fmt.Println("creating a new instance")
			instance = &DummyConfig{
				Port:                8080,
				Host:                "localhost",
				EncoddedEncyptedKey: "secret-key",
			}
		}
	})
	return instance
}
