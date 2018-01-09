package main

import (
	"flag"
	"fmt"
	"strconv"
)

type Token struct {
	data      string
	recipient int
	ttl       int
}

// call go run lab4.go 15 "message",
func main() {
	flag.Parse()
	msg := flag.Arg(0)
	rcp, _ := strconv.Atoi(flag.Arg(1))
	time, _ := strconv.Atoi(flag.Arg(2))
	numСhannels, _ := strconv.Atoi(flag.Arg(3))
	arr := make([]chan Token, numСhannels, 2*numСhannels)
	Tok := Token{msg, rcp, time}
	first_channel := make(chan Token)
	arr[0] = first_channel
	createChannels(numСhannels, arr)

	first_channel <- Tok
}
func node(in <-chan Token, out chan<- Token, num int) { // подумать над num
	selfNum := num
	Tok := <-in
	Tok.ttl -= 1
	if Tok.recipient == selfNum {
		fmt.Println("Token is delivered", Tok.data)
		close(out)
		return
	}
	if Tok.ttl == 0 {
		fmt.Println("Token died")
		close(out)
		return
	}
	out <- Tok
}

func createChannels(numСhannels int, arr []chan Token) {
	for i := 1; i < numСhannels; i++ {
		if i == 1 {
			arr[1] = make(chan Token)
			go node(arr[0], arr[1], 1)
		} else if i == numСhannels-1 {
			go node(arr[i], arr[0], i)
		} else {
			arr[i+1] = make(chan Token)
			go node(arr[i], arr[i+1], i)
		}
	}
}
