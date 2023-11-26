package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	count := 5
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(r int) {
			defer wg.Done()
			fmt.Println(fakeHttpGet(r))
		}(i)
	}
	wg.Wait()
}

func fakeHttpGet(i int) string {
	return fakeHtppServer(i)
}
func fakeHtppServer(i int) string {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return fmt.Sprintf("Fake response #%v", i)
}
