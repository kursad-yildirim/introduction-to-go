package main

import "fmt"

var x, y int = 4, 5

type cardface string

func main() {
	// const faceCount = 4 //character, string, boolean, or numeric values
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

	names := [4]string{
		"Ugo",
		"Ignacio",
		"Danielle",
		"Francesco",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, len(primes), cap(primes))
	slicePrimes := primes[:0]
	slicePrint(slicePrimes)
	slicePrimes = primes[:4]
	slicePrint(slicePrimes)
	slicePrimes = primes[1:5]
	slicePrint(slicePrimes)
	slicePrimes = primes[:5]
	slicePrint(slicePrimes)
	//	slicePrimes[5] = 13
	slicePrimes = append(slicePrimes, 13)
	slicePrint(slicePrimes)
	// slicePrimes[1] = 99
	slicePrimes = append(slicePrimes, 17)
	slicePrint(slicePrimes)
	slicePrimes = slicePrimes[:len(slicePrimes)+1]
	//	slicePrimes[1] = 99
	slicePrint(slicePrimes)
	slicePrint(slicePrimes)
	newPrimes := [7]int{2, 3, 5, 7, 11, 13, 17}
	slicePrimes = newPrimes[:]
	slicePrint(slicePrimes)
	slicePrimes = primes[0:3]
	newSlice := primes[2:5]
	slicePrint(slicePrimes)
	slicePrint(newSlice)
	copy(slicePrimes, newSlice)
	slicePrint(slicePrimes)
	slicePrint(newSlice)
	slicePrimes[2] = 98
	slicePrint(slicePrimes)
	slicePrint(newSlice)

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

func slicePrint(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
