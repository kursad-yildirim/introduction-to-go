package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	toSend := "hello routine"
	// go routine
	go fmt.Println(toSend)
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < 5; i++ {
		i := i
		go fmt.Println("index is", i)
	}
	time.Sleep(10 * time.Millisecond)

	// go channel
	ch := make(chan string)
	/*ch <- toSend
	received := <-ch
	fmt.Println(received) */
	// deadlock

	// routines and channels
	go func() {
		ch <- toSend
	}()
	received := <-ch
	fmt.Println(received)

	// range over channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("index is %d", i)
		}
		close(ch)
	}()

	for received := range ch {
		fmt.Println(received)
	}

	// closed channel
	whatsLeft, ok := <-ch
	fmt.Printf("What's left: %#v, %v\n", whatsLeft, ok)
	// PANIC: ch <- "new message"
	// PANIC: close(ch)

	// nil channels

	toChOdd := []int{1, 3, 5, 7}
	toChEven := []int{2, 4, 6, 8}

	chOdd := myChan(toChOdd)
	chEven := myChan(toChEven)

	chMerged := mergeChan(chOdd, chEven)

	for fromMerged := range chMerged {
		fmt.Println(fromMerged)
	}
	for {
		fromMerged, ok := <-chMerged
		if !ok {
			break
		}
		fmt.Println(fromMerged)
	}
}

func myChan(intToSend []int) <-chan int {
	ch := make(chan int)
	go func(ints []int) {
		for _, x := range ints {
			ch <- x
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch) // owner closes the channel
	}(intToSend)

	return ch
}

func mergeChan(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		//		closed1, closed2 := false, false
		//		for !closed1 || !closed2 {
		for ch1 != nil || ch2 != nil {
			select {
			case x, ok := <-ch1:
				if !ok {
					fmt.Println("channel 1 is closed")
					//closed1 = true
					ch1 = nil
					continue
				}
				ch <- x
			case x, ok := <-ch2:
				if !ok {
					fmt.Println("channel 2 is closed")
					//closed2 = true
					ch2 = nil
					continue
				}
				ch <- x
			}
		}
	}()

	return ch
}
