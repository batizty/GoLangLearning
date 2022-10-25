package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("--> Begin Init")
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Init Rand OK, rand = ", rand.ExpFloat64())
	fmt.Println("<-- End Init")
}

func player(userName string, ch chan int64) {
	defer wg.Done()

	for {
		ball, _ok := <-ch
		if !_ok {
			fmt.Println("Player ", userName, " is the winner")
			return
		}

		if rand.Intn(100)%13 == 0 {
			fmt.Println("Player ", userName, " miss this round")
			close(ch)
			return
		}

		fmt.Println("Player ", userName, " Play round ", ball)
		ball++
		ch <- ball
	}
}

func main() {
	fmt.Println("--> Begin Main")

	// Add channel
	count := make(chan int64)

	// add wait group
	wg.Add(2)

	go player("tuoyu", count)
	go player("huahua", count)

	count <- 1

	// wait all thread are done
	wg.Wait()
	fmt.Println("<-- End Main")
}
