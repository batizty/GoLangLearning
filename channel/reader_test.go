package channel

import (
	"sync"
	"testing"
)

func TestRead(t *testing.T) {
	strBuffer := make(chan string, 10)
	var wg sync.WaitGroup
	wg.Add(1)

	strBuffer <- "hello world"

	go Read(strBuffer, &wg)
	close(strBuffer)
	wg.Wait()
	t.Log("hello OK")

}
