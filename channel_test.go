package go_routines

import (
	"fmt"
	"testing"
	"time"
)


func TestCreateChannel(t *testing.T) {
	wait := make(chan string)
	defer close(wait)

	go func() {
		time.Sleep(2 * time.Second)
		wait <- "Lorem ipsum dolor"
		fmt.Println("Finished.")
	}()

	data := <- wait
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

//Paramerter in channel by default pass by reference (!pointer)
func GiveMeResponse(wait chan string) {
	time.Sleep(2 * time.Second)
	wait <- "Lorem ipsum dolorr"
}

func TestChannelAsParameter(t *testing.T) {
	wait := make(chan string)

	go GiveMeResponse(wait)

	data := <-wait
	fmt.Println(data)
	close(wait)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Lorem ipsum"
}

func OnlyOut(channel <-chan string) {
	data:= <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}