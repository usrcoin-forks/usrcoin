package sys

import (
	"fmt"
	"runtime"
)

type empty struct{}

var (
	MAX_GOROUTINES = 4 * runtime.NumCPU()
	tickets        chan empty
)

func init() {
	fmt.Println("MAX_GOROUTINES:", MAX_GOROUTINES)
	tickets = make(chan empty, MAX_GOROUTINES)
}

func GetTicket() {
	tickets <- empty{}
}

func FreeTicket() {
	<-tickets
}
