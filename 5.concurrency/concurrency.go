package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	//	ch <- "hi"
	go func() {
		ch <- "hi"
	}()
	msg := <-ch
	fmt.Println(msg)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("message #%v", i)
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	msg, ok := <-ch
	fmt.Printf("%#v (ok=%v)", msg, ok)

	// ch <- "hey" // panic, try to recolse will panic, send/receive to a nil channell will block forever why do we have nil channels

	/*
		for {
			msg,ok := <-ch
			if !ok {
				break
			}
			fmt.Println(msg)
		}
	*/
}
