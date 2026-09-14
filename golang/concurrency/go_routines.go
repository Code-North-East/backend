package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Write a basic function where a go routine is invoked.
func CreateBasicGoRoutine() {
	fmt.Println("This example demonstrates the basic usage of invoking a go routine")
	go ConvertINRToUSD(139000)
	// wait for a second to allow the go routine to complete the execution before the main function exists.
	time.Sleep(time.Second)

	// Alternatively you can can cretae a go routine in a inline function closure too.
	go func() {
		fmt.Println("Hello, I am a go routine")
	}()
	time.Sleep(time.Second)
}

// Write a basic function where a go routine is invoked using wait group.
func CreateBasicGoRoutineUsingWaitGroup() {
	// The ideal way for waiting for a go routine to fininsh is by using wait group, because it will ensure the main function waits untll all the go routines finish off execution

	// creating a wait group
	var wg sync.WaitGroup
	// adding a delta of 1 to the wait group
	wg.Add(1)
	// declaring a inline go routine closure
	go func(message string) {
		// after finishing off the execution Done will be called to decrement the wg counter
		defer wg.Done()
		fmt.Println("user says: ", message)
	}("hello")
	// Will wait till the main function is done executing the go routine.
	wg.Wait()
}

// Write a function for the loop launcher, the basic idea is to have 10 go routines launch and the print the worker id from the go routine.
func LaunchWorkerLoop() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wid uint64) {
			defer wg.Done()
			PrintWorkerID(uint64(i))
		}(uint64(i))
	}
	wg.Wait()
}

// Write a function for the worker opration, in our case, it is simply printing the worker id
// Pro tip: Keep the wg passing away from the worker function.
func PrintWorkerID(id uint64) {
	fmt.Printf("Printing from worker: %d\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Execution finished from the workder: %d\n", id)
}

// Write a function which will convert INR to USD
func ConvertINRToUSD(inr float64) (float64, error) {
	if inr <= 0 {
		return 0, fmt.Errorf("Invalid amount: %f", inr)
	}
	usd := inr * (1 / 95.4)
	fmt.Println("Converted value: $", usd)
	return usd, nil
}

// Write a function which will simulate the concurrent downloading of files from a dummy url set.
// Pro tip: Always pass the other props into the inline go routing closure, because from versions earlier than 1.22 this code will run into a race condition.
func ConcurrentDummyDownloader() {
	urls := []string{
		"https://example1.com",
		"https://example2.com",
		"https://example3.com",
		"https://example4.com",
		"https://example5.com",
	}
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(wid uint64, url string) {
			defer wg.Done()
			DownloadFromDummyURL(url, wid)
		}(uint64(i), url)
	}
	wg.Wait()
	fmt.Println("All downloads are finished!")
}

// Writing a function to mock the download operation from a dummmy url
func DownloadFromDummyURL(url string, wid uint64) {
	fmt.Println("downloading from url :", url, "using worker: ", wid)
	time.Sleep(time.Second * 1)
	fmt.Println("download complete from url: ", url)
}

// Write a function where you are spawing 1000 go routines to increment the bank balance by 1 and the final bank balance should be 1000.
func IncrementBankBalance() {
	var bal int
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i:=1; i<=1000; i++ {
		//spawn 1000 go routines to increment the bank balance
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			bal = bal + 1
		}()
	}
	wg.Wait()
	fmt.Println("Final bank balance: ", bal)
}