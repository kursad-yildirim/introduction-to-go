package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(divide(8, 2))
	fmt.Println(divide(8, 0))
}

func divide(x, y int) (q int, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	q = x / y
	return
}

/*
func divide(x, y int) int {
	return x / y
}
*/
