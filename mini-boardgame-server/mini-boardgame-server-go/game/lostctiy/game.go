package lostctiy

import "math/rand"

type (
	Card int

	Config struct {
		DevelopPoint int
		BonusCount   int
		BonusPoint   int
	}

	Game struct {
		Deck  []Card     // [index]
		Board [][][]Card // [player][color][index]
		Drop  [][]Card   // [color][index]
		Hand  [][]Card   // [player][index]
	}
)

var (
	//DefaultConfig = Config{
	//	DevelopPoint: 20,
	//	BonusCount:   8,
	//	BonusPoint:   20,
	//}

	PLAYER      = 2
	COLOR       = 5
	CARD        = 12
	CARD_DOUBLE = 3
	CARD_POINT  = CARD - CARD_DOUBLE
)

func NewStandardGame() *Game {
	deck := make([]Card, CARD*COLOR)
	for i := range deck {
		deck[i] = Card(i)
	}
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	board := make([][][]Card, PLAYER)
	for i := range board {
		board[i] = make([][]Card, COLOR)
		for m := range board[i] {
			board[i][m] = make([]Card, 0)
		}
	}

	drop := make([][]Card, COLOR)
	for i := range drop {
		drop[i] = make([]Card, 0)
	}

	hand := make([][]Card, PLAYER)
	for i := range board {
		hand[i] = make([]Card, 0)
	}
	return &Game{
		Deck:  deck,
		Board: board,
		Drop:  drop,
		Hand:  hand,
	}
}

func (c Card) Color() int {
	return int(c) / CARD
}

func (c Card) Point() int {
	if c.IsDouble() {
		return 0
	} else {
		return int(c) - CARD_DOUBLE + 2
	}
}

func (c Card) IsDouble() bool {
	return int(c) < CARD_DOUBLE
}
