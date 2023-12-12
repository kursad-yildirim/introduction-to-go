package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
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
	fmt.Printf("\n\nESS SSH Client connected to %#v started:\n\n", ess.Env["host"])
	<-out
	in <- "export IS_IT_ME=YES"
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", <-out)
	for reader.Scan() {
		if reader.Text() == "exit" {
			in <- "echo $IS_IT_ME"
			if strings.Contains(<-out, "YES") {
				break
			}
		}
		in <- reader.Text()
		fmt.Printf("%s", <-out)
	}
	fmt.Printf("ESS SSH session to %#v ended:\n", ess.Env["host"])
}
