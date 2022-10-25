package main

import (
	"GoLangLearning/channel"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("--> Begin Init")
	fmt.Println("<-- End Init")
}

func main() {
	fmt.Println("--> Begin Main")

	// Add channel
	strBufferForRW := make(chan string, 20)

	wg.Add(2)

	go channel.Read(strBufferForRW, &wg)
	go channel.Write(strBufferForRW, &wg)

	wg.Wait()

	fmt.Println("Finish All Worker")
	return
}
