package main

import (
	"fmt"
	"log"

	"tuff.local/concurrency/ssh/ess"
)

func main() {
	err := ess.Env.Check()
	if err != nil {
		log.Fatalf("ERROR:\n%v", err)
	}
	in, out, err := ess.DoSSH(ess.Env["host"], ess.Env["port"])
	if err != nil {
		log.Fatalf("Error: SSH to %#v failed with %#v", ess.Env["host"], fmt.Sprintf("%s", err))
	}
	fmt.Println("hello 0")
	fmt.Printf("%s", <-out)
	in <- "ping -c 1 localhost"
	fmt.Printf("%s", <-out)
	in <- "exit"
}
