package logger

import (
	"sync"

	"go.uber.org/zap"
)

// Steps to be followed -

// 1. Create a singleton for building the logger object.
// 2. Create the functions for debug, info and error level logs.
// 3. All the packages inside this repository should be using the in-built logger function.

type LogHelper struct {
	File   string
	Logger *zap.Logger
}

var instance *LogHelper
var once *sync.Once

func Build(file string) (*LogHelper, error) {
	once.Do(func() {
		// build the logger
		if logger, err := zap.NewProduction(); err == nil {
			defer logger.Sync()
			instance = &LogHelper{
				File: file,
			}
		} else {
			// raise an exception or return err
			return nil, err
		}
	})
	return instance, nil
}

// todo: implement later
func BuildWithSugar() {
	// for cases where performance is not that critical...
}

func (lh *LogHelper) InfoMessage() {
	// for info level messages
}

func (lh *LogHelper) DebugMessage() {
	// for debug level messages
}
