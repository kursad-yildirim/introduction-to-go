package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := "./test.yaml.gz"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	//	defer file.Close()
	defer func() {
		file.Close()
	}()
	/*
		var r io.Reader = file

		if strings.HasSuffix(filename, ".gz") {
			gr, err := gzip.NewReader(file)
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			defer gr.Close()
			r = gr
		}
	*/

	r, err := gzip.NewReader(file)
	defer r.Close()

	//	io.Copy(os.Stdout, gr)
	sh := sha1.New()

	if _, err := io.Copy(sh, r); err != nil {
		log.Fatalf("Error: %s", err)
	}

	sha1sum := sh.Sum(nil)

	fmt.Printf("%x", sha1sum)
}
