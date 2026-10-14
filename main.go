package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

//渲染绘制
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

//分辨率
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	var g = &Game{}
	g.player = NewPlayer(g)
	ebiten.SetWindowTitle("Go Asterolds")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
