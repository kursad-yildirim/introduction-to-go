package main

import "fmt"

type cardFace string
type cardValue string
type faceArray []cardValue

type singleFace struct {
	faceName  cardFace
	faceCards faceArray
}

type deck struct {
	games
	faces []singleFace
}

type games struct {
	gameSlice []string
}

func main() {
	faceSlice := []cardFace{"diamonds", "hearts", "clubs", "spades"}
	cardSlice := []cardValue{"one", "two", "three", "queen", "king"}

	currentFace := singleFace{}
	myDeck := deck{
		games: games{
			gameSlice: []string{"blackjack", "poker"},
		},
		faces: []singleFace{},
	}
	for _, face := range faceSlice {
		currentFace.faceName = face
		for _, card := range cardSlice {
			currentFace.faceCards = append(currentFace.faceCards, card)
		}
		myDeck.faces = append(myDeck.faces, currentFace)
		currentFace.faceCards = nil
	}

	fmt.Printf("%v\n", myDeck)
	fmt.Println(myDeck.games.gameSlice)
	fmt.Println(myDeck.games.gameSlice)
}
