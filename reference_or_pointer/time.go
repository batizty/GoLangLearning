package main

import (
	"fmt"
)

type SelfTime struct {
	sec int64
}

func (st SelfTime) output() {
	fmt.Println("sec: ", st.sec)
}

func (st SelfTime) add(duration int64) SelfTime {
	st.sec += duration
	return st
}

func main() {
	var st SelfTime
	st.sec = 10000
	fmt.Println("st init")
	st.output()

	var st2 = st.add(100)
	fmt.Println("st2")
	st2.output()

	fmt.Println("st")
	st.output()

}
