package main

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
)

func main() {
	if logger, err := zap.NewProduction(); err != nil {
		log.Fatalf("could not use production zap logger")
	} else {
		s := fmt.Sprintf("successfully containerized the backend code - %s", time.Now().String())
		logger.Info(s)
	}
}
