package lostctiy

func (g *Game) Start() {
	go g.GameLoop()
}

func (g *Game) GameLoop() {
	g.DrawCard(0, 7)
	g.DrawCard(1, 7)

}
