package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type player struct {
// 	deck []int
// 	discard []int
// }

func main() {
	/*
		Be able to run this 1000s of times and run weird scenarios to get probability
		Be able to import from seperate file a deck schema for each player
	*/

	d := []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 10, 10, 10, 10, 11, 11, 11, 11, 12, 12, 12, 12, 13, 13, 13, 13}

	fmt.Println(d)
	d = shuffle(d)
	fmt.Println(d)

	hands := Deal(d)
	p1 := hands[0]
	p2 := hands[1]

	count := 1
	for {
		fmt.Println("Round ", count)
		count++
		cardsToPlay := 1

		p1, p2 = Battle(p1, p2, cardsToPlay)

		if (len(p1) + len(p2) < 52) {
			fmt.Println("Card fell under the table")
			break;
		}

		if len(p1) == 0 {
				winner := "p2"
				fmt.Println("The Winner is ", winner)
				break
		}
		if len(p2) == 0 {
				winner := "p1"
				fmt.Println("The Winner is ", winner)
				break
		}

		if count > 10000 {
			fmt.Println("Go outside")
			break
		}
		fmt.Printf("\n")
	}
}

func shuffle(d []int) []int {
	shuffled := make([]int, len(d))

	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(d))

	for i, v := range perm {
			shuffled[v] = d[i]
	}
	return shuffled
}

func Deal(d []int) [][]int {
	players := make([][]int, 2)
	var evenIndexDeck []int
	var oddIndexDeck []int
	for i := range d {
		if i % 2 == 0 {
			evenIndexDeck = append(evenIndexDeck, d[i])
		} else {
			oddIndexDeck = append(oddIndexDeck, d[i])
		}
	}
	players[0] = evenIndexDeck
	players[1] = oddIndexDeck
	fmt.Printf("Dealing: \nPlayer 1: \t%v\nPlayer 2: \t%v\n", evenIndexDeck, oddIndexDeck)
	return players
}

func Play(d []int, n int) ([]int, []int) {
	var playedCards []int
	if len(d) < n {
		playedCards = append(playedCards, d...)
		return []int{}, playedCards
	}

	for i := 0; i < n; i++ {
		playedCards = append(playedCards, d[i])
	}
	return d[n:], playedCards
}

func Battle(p1 []int, p2 []int, n int, pc ...int) ([]int, []int) {
	// var winner string
	var playedCards []int
	playedCards = append(playedCards, pc...)

	fmt.Println("Player 1 and 2 battle!")
	p1, c1 := Play(p1, n)
	p2, c2 := Play(p2, n) 

	playedCards = append(playedCards, c1...)
	playedCards = append(playedCards, c2...)
	shuffle(playedCards)

	if len(p1) == 0 || len(p2) == 0 {
		if len(p1) == 0 {
			// winner = "p2"
			p2 = append(p2, playedCards...)
			fmt.Printf("Player 1 cards: %d\tPlayer 2 cards: %d\n", len(p1), len(p2))
			return p1, p2
		}
		if len(p2) == 0 {
			// winner = "p1"
			p1 = append(p1, playedCards...)
			fmt.Printf("Player 1 cards: %d\tPlayer 2 cards: %d\n", len(p1), len(p2))
			return p1, p2
		}
		
	}

	p1ShowedCard := c1[n - 1]
	p2ShowedCard := c2[n - 1]
	fmt.Printf("Player 1's Card: [%d]\tPlayer 2's Card: [%d]\n", p1ShowedCard, p2ShowedCard)

	switch {
	case p1ShowedCard > p2ShowedCard:
		p1 = append(p1, playedCards...)
		// winner = "p1"
	case p2ShowedCard > p1ShowedCard:
		p2 = append(p2, playedCards...)
		// winner = "p2"
	case p1ShowedCard == p2ShowedCard:
		// todo fix out of index error
		fmt.Println("**********************WAR**********************")
		n = 2
		p1, p2 = Battle(p1, p2, n, playedCards...)
	}
	// fmt.Println("Battle decided: ", winner)
	fmt.Printf("Player 1 cards: %d\tPlayer 2 cards: %d\n", len(p1), len(p2))
	return p1, p2
}
