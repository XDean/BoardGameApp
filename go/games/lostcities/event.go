package lostcities

import (
	"errors"
	"github.com/xdean/miniboardgame/go/server/game"
)

type (
	// Events
	Event struct {
		game.BaseEvent

		Play bool
		// Play card
		Card Card
		Drop bool
		// Draw card
		FromDeck bool // Or from drop
		Color    int  // Only available when from deck
	}
)

func (g *Game) Play(event Event) game.Response {
	seat := event.GetSeat()
	if g.Current != seat {
		return errors.New("You are not current player")
	}
	if event.Drop && !event.FromDeck && (event.Card.Color() == event.Color) {
		return errors.New("You can't draw the drop card immediatly")
	}
	if !g.hasCard(seat, event.Card) {
		return errors.New("You don't have the card to play")
	}
	cards := g.Board[seat][event.Card.Color()]
	if !event.Drop && len(cards) > 0 && cards[0].Point() > event.Card.Point() {
		return errors.New("You can't play the card because you already have a lagrer one")
	}
	if !event.FromDeck && len(g.Drop[event.Color]) == 0 {
		return errors.New("No card to draw in this color's drop area")
	}
	// TODO Check deck has card
	if event.Play {
		if event.Drop {
			g.DropCard(seat, event.Card)
			return "Drop Success"
		} else {
			g.PlayCard(seat, event.Card)
			return "Play Success"
		}
	} else {
		if event.FromDeck {
			return g.DrawCard(seat, 1)[0]
		} else {
			return g.DrawDropCard(seat, event.Color)
		}
	}
}

func (g *Game) PlayCard(player int, card Card) {
	if g.removeHandCard(player, card) {
		g.Board[player][card.Color()] = append(g.Board[player][card.Color()], card)
	}
}

func (g *Game) DropCard(player int, card Card) {
	if g.removeHandCard(player, card) {
		g.Drop[card.Color()] = append(g.Drop[card.Color()], card)
	}
}

func (g *Game) DrawCard(player int, count int) []Card {
	card := g.Deck[:count]
	g.Deck = g.Deck[count:]
	g.Hand[player] = append(g.Hand[player], card...)
	return card
}

func (g *Game) DrawDropCard(player int, color int) Card {
	drop := g.Drop[color]
	card := drop[len(drop)-1]
	g.Drop[color] = drop[:len(drop)-1]
	g.Hand[player] = append(g.Hand[player], card)
	return card
}
