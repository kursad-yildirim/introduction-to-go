package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type example struct{}

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	// fmt.Println(resp)

	//rb := make([]byte, 20)
	//rs, err := resp.Body.Read(rb)

	//fmt.Println(string(rb))
	//fmt.Println("Number of bytes read from response:", rs)

	var exampleCom example
	//ws, err := exampleCom.Write(rb)

	io.Copy(exampleCom, resp.Body)

	//fmt.Println("Number of bytes written to console:", ws)

	/*
		rs, err := io.Copy(os.Stdout, resp.Body)
		fmt.Println("Number of bytes read from response:", rs)
	*/

}

func (example) Write(p []byte) (n int, err error) {
	fmt.Println("custom writer interface")
	fmt.Println(string(p))
	//fmt.Printf("%s", p)
	return len(p), nil
}
