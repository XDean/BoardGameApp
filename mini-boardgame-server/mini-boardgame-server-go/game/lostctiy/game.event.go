package lostctiy

import "errors"

type (
	// Events
	Event struct {
		// Play card
		Card Card
		Drop bool
		// Draw card
		FromDeck bool // Or from drop
		Color    int  // Only available when from deck
	}
)

func (g *Game) Play(player int, event Event) (Card, error) {
	if g.Current != player {
		return EMPTY_CARD, errors.New("You are not current player")
	}
	if event.Drop && !event.FromDeck && (event.Card.Color() == event.Color) {
		return EMPTY_CARD, errors.New("You can't draw the drop card immediatly")
	}
	if !g.hasCard(player, event.Card) {
		return EMPTY_CARD, errors.New("You don't have the card to play")
	}
	cards := g.Board[player][event.Card.Color()]
	if !event.Drop && len(cards) > 0 && cards[0].Point() > event.Card.Point() {
		return EMPTY_CARD, errors.New("You can't play the card because you already have a lagrer one")
	}
	if !event.FromDeck && len(g.Drop[event.Color]) == 0 {
		return EMPTY_CARD, errors.New("No card to draw in this color's drop area")
	}
	// TODO Check deck has card
	if event.Drop {
		g.DropCard(player, event.Card)
	} else {
		g.PlayCard(player, event.Card)
	}
	if event.FromDeck {
		return g.DrawCard(player, 1)[0], nil
	} else {
		return g.DrawDropCard(player, event.Color), nil
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
