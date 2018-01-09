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

// call from terminal  go run lab4.go msg recipient_num lifetime number_of_elements
func main() {
	flag.Parse()
	msg := flag.Arg(0)
	rcp, _ := strconv.Atoi(flag.Arg(1))
	time, _ := strconv.Atoi(flag.Arg(2))
	numСhannels, _ := strconv.Atoi(flag.Arg(3))
	arr := make([]chan *Token, numСhannels, 2*numСhannels)
	Tok := Token{msg, rcp, time}

	for i := 0; i < numСhannels; i++ {
		arr[i] = make(chan *Token)
	}

	go node(arr[numСhannels-1], arr[0], numСhannels)

	for i := numСhannels - 2; i > 1; i-- {
		go node(arr[i], arr[i+1], i)
	}

	go node(arr[0], arr[1], 0)
	go node(arr[1], arr[2], 1)
	arr[0] <- &Tok
}

func node(in <-chan *Token, out chan<- *Token, num int) {
	selfNum := num
	for {
		Tok := <-in
		println(selfNum)
		Tok.ttl -= 1
		if Tok.recipient == selfNum {
			fmt.Print("Token is delivered to ")
			fmt.Println(Tok.recipient)
			fmt.Println(Tok.data)
			return
		}
		if Tok.ttl < 0 {
			fmt.Println("Token died")
			return
		}
		out <- Tok
	}
}
