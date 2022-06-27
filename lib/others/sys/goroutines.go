package sys

import "runtime"

type empty struct{}

var (
	MAX_GOROUTINES = runtime.NumCPU()
	tickets        chan empty
)

func init() {
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
