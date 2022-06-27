package sys

import (
	"fmt"
	"runtime"
)

type empty struct{}

var (
	MAX_GOROUTINES = runtime.NumCPU()
	tickets        chan empty
)

func init() {
	fmt.Println("Maximum number of hard working go-routines:", MAX_GOROUTINES)
	tickets = make(chan empty, MAX_GOROUTINES)
	for i := 0; i < cap(tickets); i++ {
		tickets <- empty{}
	}
}

func GetTicket() {
	<-tickets
}

func FreeTicket() {
	tickets <- empty{}
}
