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
	faces      []singleFace
	playerName string
	player
}

type games struct {
	gameSlice  []string
	playerName string
}

type player struct {
	playerName  string
	playerLevel string
}

func main() {
	faceSlice := []cardFace{"diamonds", "hearts", "clubs", "spades"}
	cardSlice := []cardValue{"one", "two", "three", "queen", "king"}

	currentFace := singleFace{}
	myDeck := deck{
		games: games{
			gameSlice:  []string{"blackjack", "poker"},
			playerName: "kyildiri",
		},
		faces:      []singleFace{},
		playerName: "kursad",
		player: player{
			playerName:  "yildirim",
			playerLevel: "beginner",
		},
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
	fmt.Println(myDeck.gameSlice)
	fmt.Println(myDeck.games.gameSlice)
	fmt.Println(myDeck.playerName)
	fmt.Println(myDeck.games.playerName)

	myDeck.addCard("five", 3)
	myDeck.addCard("four", 3)
	myDeck.addCard("ten", 5)
	myDeck.addCard("jack", 6)
	fmt.Println(myDeck.faces)
}

func (d *deck) addCard(v cardValue, i int) {
	fmt.Println("ADD CARD", v, i)
	for index, cards := range d.faces {
		//fmt.Println(cards.faceCards)
		cards.faceCards = insertCard(cards.faceCards, v, i)
		//fmt.Println(cards.faceCards)
		d.faces[index] = cards
	}
}

func insertCard(f faceArray, c cardValue, ind int) faceArray {
	var s faceArray
	for i := 0; i < ind; i++ {
		s = append(s, f[i])
	}
	s = append(s, c)
	for i := ind; i < len(f); i++ {
		s = append(s, f[i])
	}

	return s
}
