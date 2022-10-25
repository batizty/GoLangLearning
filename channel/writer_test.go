package channel

import (
	"sync"
	"testing"
)

func TestWrite(t *testing.T) {
	strBuffer := make(chan string, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	strBuffer <- "hello world"

	go Write(strBuffer, &wg)
	go Read(strBuffer, &wg)

	wg.Wait()
	t.Log("hello OK")

}
