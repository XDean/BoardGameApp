package lostcities

import (
	"github.com/thoas/go-funk"
	"math/rand"
)

const (
	DEVELOP_POINT = 20
	BONUS_COUNT   = 8
	BONUS_POINT   = 20

	PLAYER      = 2
	COLOR       = 5
	CARD        = 12
	CARD_DOUBLE = 3
	CARD_POINT  = CARD - CARD_DOUBLE
)

type (
	Card struct {
		int
	}

	Game struct {
		EventStream chan<- interface{}
		Current     int
		Deck        []Card     // [index] from 0 (top)
		Board       [][][]Card // [player][color][index] from 0 (oldest)
		Drop        [][]Card   // [color][index] from 0 (oldest)
		Hand        [][]Card   // [player][index] no order, by default 0 (oldest)
	}
)

var (
	EMPTY_CARD = Card{-1}
)

func NewStandardGame() *Game {
	deck := make([]Card, CARD*COLOR)
	for i := range deck {
		deck[i] = Card{i}
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
	g := &Game{
		Deck:  deck,
		Board: board,
		Drop:  drop,
		Hand:  hand,
	}
	g.DrawCard(0, 7)
	g.DrawCard(1, 7)
	return g
}

func (c Card) Color() int {
	return c.int / CARD
}

func (c Card) Point() int {
	if c.IsDouble() {
		return 0
	} else {
		return c.int - CARD_DOUBLE + 2
	}
}

func (c Card) IsDouble() bool {
	return c.int < CARD_DOUBLE
}

func (g *Game) hasCard(player int, card Card) bool {
	return funk.Contains(g.Hand[player], card)
}

func (g *Game) removeHandCard(player int, card Card) bool {
	index := funk.IndexOf(g.Hand[player], card)
	if index != -1 {
		g.Hand = append(g.Hand[:index], g.Hand[index+1:]...)
		return true
	}
	return false
}
