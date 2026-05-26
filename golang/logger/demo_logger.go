package logger

import "go.uber.org/zap"

// write all the functions and possible combinations for each of the levels and then start writing the library...

var genericInfoLogMessage = "injecting an info level log"

// To curb repetitive code, we can have sort of a helper function over here...
func GetZapLogger() *zap.Logger {
	if logger, err := zap.NewProduction(); err != nil {
		// close the application and panic
		failureMessage := "could not instantiate zap logger!"
		panic(failureMessage)
	} else {
		return logger
	}
}

func ZapLoggerDemoInfo() {
	if logger, err := zap.NewProduction(); err != nil {
		// close the application and panic
		failureMessage := "could not instantiate zap logger!"
		panic(failureMessage)
	} else {
		defer logger.Sync()
		logger.Info(genericInfoLogMessage)
	}
}

func ZapLoggerDemoInfoCustomFields() {
	// for this example we can add a string custom field, but we need to test with additional fields as well.
	logger := GetZapLogger()
	defer logger.Sync()
	logger.Info(
		genericInfoLogMessage, 
		zap.String("code", "OPTIMAL"),
	)
}