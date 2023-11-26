package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "./test.yaml.gz"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(filename, ".gz") {
		gr, err := gzip.NewReader(file)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		defer gr.Close()
		r = gr
	}
	sh := sha1.New()
	io.Copy(sh, r)
	shasum := sh.Sum(nil)
	fmt.Printf("%x", shasum)
}
