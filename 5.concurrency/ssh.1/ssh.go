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
	cmd := "ping -c 1 localhost"
	fmt.Printf("%s %s", <-out, cmd)
	in <- cmd
	fmt.Printf("%s", <-out)
	in <- "exit"
}
