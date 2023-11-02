package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Serial execution")
	start := time.Now()
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println(fakeHttpSerialGet(i))
	}
	fmt.Println("Elapsed time:", time.Since(start))
	fmt.Println("With routines and channels")
	start = time.Now()
	ch := make(chan string)
	for i := 0; i < count; i++ {
		go fakeHttpGet(i, ch)
	}
	for i := 0; i < count; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Elapsed time:", time.Since(start))
}

func fakeHttpGet(i int, ch chan string) {
	ch <- fakeHtppServer(i)
}
func fakeHttpSerialGet(i int) string {
	return fakeHtppServer(i)
}

func fakeHtppServer(i int) string {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return fmt.Sprintf("Fake response #%v", i)
}
