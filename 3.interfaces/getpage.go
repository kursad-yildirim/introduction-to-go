package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	//bSlice := make([]byte, 99999)
	//resp.Body.Read(bSlice)
	//	fmt.Println(string(bSlice))

	io.Copy(os.Stdout, resp.Body)

}
