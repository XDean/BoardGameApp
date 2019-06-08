package lostctiy

import "errors"

type (
	// Events
	PlayCard struct {
		Card Card
		Drop bool
	}

	DrawCard struct {
		FromDeck bool // Or from drop
		Color    int  // Only available when from deck
	}
)

func (g *Game) Play(player int, play PlayCard, draw DrawCard) error {
	if g.Current != player {
		return errors.New("You are not current player")
	}
	if play.Drop && !draw.FromDeck && (play.Card.Color() == draw.Color) {
		return errors.New("You can't draw the drop card immediatly")
	}
	if !g.hasCard(player, play.Card) {
		return errors.New("You don't have the card to play")
	}
	cards := g.Board[player][play.Card.Color()]
	if !play.Drop && len(cards) > 0 && cards[0].Point() > play.Card.Point() {
		return errors.New("You can't play the card because you already have a lagrer one")
	}
	if !draw.FromDeck && len(g.Drop[draw.Color]) == 0 {
		return errors.New("No card to draw in this color's drop area")
	}
	// TODO Check deck has card
	if play.Drop {
		g.DropCard(player, play.Card)
	} else {
		g.PlayCard(player, play.Card)
	}
	if draw.FromDeck {
		g.DrawCard(player)
	} else {
		g.DrawDropCard(player, draw.Color)
	}
	return nil
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

func (g *Game) DrawCard(player int) Card {
	card := g.Deck[0]
	g.Deck = g.Deck[1:]
	g.Hand[player] = append(g.Hand[player], card)
	return card
}

func (g *Game) DrawDropCard(player int, color int) Card {
	drop := g.Drop[color]
	card := drop[len(drop)-1]
	g.Drop[color] = drop[:len(drop)-1]
	g.Hand[player] = append(g.Hand[player], card)
	return card
}
