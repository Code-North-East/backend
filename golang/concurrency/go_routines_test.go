package concurrency

import (
	"fmt"
	"testing"
)

func TestConvertINRToUSD(t *testing.T) {
	if usd, err := ConvertINRToUSD(95); err == nil {
		fmt.Println("value: $", usd)
	}
}

func TestCreateBasicGoRoutine(t *testing.T) {
	CreateBasicGoRoutine()
}

func TestCreateBasicGoRoutineUsingWaitGroup(t *testing.T) {
	CreateBasicGoRoutineUsingWaitGroup()
}

func TestLaunchWorkerLoop(t *testing.T) {
	LaunchWorkerLoop()
}

func TestPrintWorkerId(t *testing.T) {
	PrintWorkerID(1)
}

func TestConcurrentDummyDownloader(t *testing.T) {
	ConcurrentDummyDownloader()
}

func TestDonwloadFromDummyURL(t *testing.T) {
	DownloadFromDummyURL("https://example.com", 1001)
}

func TestIncrementBankBalance(t *testing.T) {
	IncrementBankBalance()
}