package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var letterCount int64 = 0

func init() {
	fmt.Println("---> Begin Init ")
	rand.Seed(time.Now().Unix())
	fmt.Println("<--- End Init Reader")
}

func generateRandomString(n int) string {
	_str := make([]int32, n)
	for i := range _str {
		_str[i] = letters[rand.Intn(len(letters))]
	}
	return string(_str)

}

func Write(strBuffer chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if letterCount >= 10000 {
			fmt.Println("!!Writer write too many works, finish it")
			close(strBuffer)
			return
		}

		_randLen := rand.Intn(100)
		_randContent := generateRandomString(_randLen)

		fmt.Println("Writer send out (", _randLen, "): ", _randContent)
		strBuffer <- _randContent

		letterCount += int64(_randLen)
		time.Sleep(time.Duration(rand.Intn(10)))
	}
}
