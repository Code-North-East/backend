package logger

import "testing"

func TestZapLoggerDemoInfo(t *testing.T) {
	ZapLoggerDemoInfo()	
}

func TestZapLoggerDemoCustomFields(t *testing.T) {
	ZapLoggerDemoInfoCustomFields()
}