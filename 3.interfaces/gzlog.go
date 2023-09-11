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
	fileName := "test.yaml.gz"
	file, err := os.Open(fileName)
	if err != nil {
		/*		fmt.Println("Error:", err)
				os.Exit(1)*/
		log.Fatalf("Error: %s", err)
	}
	defer file.Close()
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		defer gz.Close()
		r = gz
	}
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		log.Fatalf("Error: %s", err)
	}
	sig := w.Sum(nil)

	fmt.Printf("%x", sig)
}
