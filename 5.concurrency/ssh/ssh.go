package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"tuff.local/concurrency/ssh/ess"
)

func main() {
	signal.Ignore(syscall.SIGINT)
	err := ess.Env.Check()
	if err != nil {
		log.Fatalf("ERROR:\n%v", err)
	}
	in, out, err := ess.DoSSH(ess.Env["host"], ess.Env["port"])
	if err != nil {
		log.Fatalf("Error: SSH to %#v failed with %#v", ess.Env["host"], fmt.Sprintf("%s", err))
	}
	fmt.Printf("ESS SSH Client connected to %#v started:\n", ess.Env["host"])
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", <-out)
	for reader.Scan() {
		if reader.Text() == "exit" {
			break
		}
		in <- reader.Text()
		go func() {
			for resp := range out {
				fmt.Printf("%s", resp)
			}
		}()
	}
	in <- reader.Text()
	fmt.Printf("ESS SSH session to %#v ended:\n", ess.Env["host"])
}
