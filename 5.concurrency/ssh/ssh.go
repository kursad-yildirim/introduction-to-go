package main

import (
	"fmt"
	"log"

	"tuff.local/concurrency/ssh/intro"
)

func main() {
	intro.Env.Check()

	in, out, err := intro.DoSSH(intro.Env["node"], intro.Env["ssh_port"])
	if err != nil {
		log.Fatalf("SSH failed to %v with Error: %v", intro.Env["node"], err)
	}
	// session and client close

	<-out //ignore the shell output
	in <- "ping -c 1 localhost"
	fmt.Printf("%s\n", <-out)
	in <- "date"
	fmt.Printf("%s\n", <-out)
	/*
		in <- "date"
		fmt.Printf("whoami: %s\n", <-out)
	*/
	in <- "exit"
	// s.Wait()

}
