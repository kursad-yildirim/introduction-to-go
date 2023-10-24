package main

import "fmt"

type cardface string

type deck struct {
	faces []cardface
}

func main() {
	const faceCount = 4

	faces := []cardface{"spades", "clubs", "hearts", "diamonds"}

	var mydeck deck
	for i, face := range faces {
		mydeck.faces = append(mydeck.faces, face)
		fmt.Println(i, ": ", face, " is added to  the deck as a new face")
	}
	fmt.Println(mydeck)
	fmt.Printf("my deck is %v\n", mydeck)
	fmt.Printf("my deck is %#v\n", mydeck)

	if len(mydeck.faces) < faceCount {
		fmt.Println("missing faces")
	} else if len(mydeck.faces) > faceCount {
		fmt.Println("Excessive faces")
	} else {
		fmt.Println("Correct faces")
	}
}
