package main

import "fmt"

var x, y int = 4, 5

type cardface string

type deck struct {
	faces []cardface
}

func main() {
	const faceCount = 4 //character, string, boolean, or numeric values
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

	blackfaces := faces[:2]
	redfaces := faces[2:]
	fmt.Println("Black faces", blackfaces, len(blackfaces), cap(blackfaces))
	fmt.Println("Red faces", redfaces, len(redfaces), cap(redfaces))
	var mydeck deck
	for _, face := range faces {
		mydeck.faces = append(mydeck.faces, face)
		fmt.Println(face, " is added to  the deck as a new face")
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
	/*
		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8

		rune // alias for int32
		     // represents a Unicode code point

		float32 float64

		complex64 complex128
	*/
}
