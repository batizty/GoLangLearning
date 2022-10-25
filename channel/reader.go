package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	fmt.Println("---> Begin Init Reader")
	rand.Seed(time.Now().Unix())
	fmt.Println("<--- End Init Reader")
}

func Read(strBuffer chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_readContent, _ok := <-strBuffer
		if !_ok {
			fmt.Println("From Reader Channel is Close")
			return
		}

		fmt.Println("Read Content(", len(_readContent), "): ", _readContent)

		_sleepN := rand.Intn(10)
		fmt.Println("Reader need to Sleep ", _sleepN)
		time.Sleep(time.Duration(_sleepN))
	}
}
