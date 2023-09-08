package main

import "fmt"

var x, y int = 4, 5

type cardface string

type deck struct {
	faces []cardface
}

func main() {
	var i, j int
	i = 1
	j = 2
	s, t := 3, 6
	num1, str1, istrue := 42, "forty two", true
	fmt.Println("i,j,s,x,y,z,t =", i, j, s, x, y, t)
	fmt.Printf("%v is %v, which is %v\n", num1, str1, istrue)
	fmt.Printf("%#v is %#v, which is %#v\n", num1, str1, istrue)

	var diamondface cardface = "diamonds"
	faces := []cardface{"spades", "clubs", "hearts"}
	fmt.Println("Missing one face:", faces, len(faces), cap(faces))
	faces = append(faces, diamondface)
	fmt.Println("Appended face", faces, len(faces), cap(faces))

	blackfaces := faces[0:2]
	fmt.Println("Black faces", blackfaces, len(blackfaces), cap(blackfaces))
	var mydeck deck
	for _, face := range faces {
		mydeck.faces = append(mydeck.faces, face)
		fmt.Println(face, " is added to  the deck as a new face")
	}
	fmt.Println(mydeck)
	fmt.Printf("my deck is %v\n", mydeck)
	fmt.Printf("my deck is %#v\n", mydeck)
}
