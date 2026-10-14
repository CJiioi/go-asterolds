package goasteroids

import "github.com/hajimehoshi/ebiten/v2"

//游戏场景
type GameScene struct {
	player *Player
}

//创建新游戏场景
func NewGameScene() *GameScene {
	var g = &GameScene{}
	g.player = NewPlayer(g)
	return g
}

func (g *GameScene) Update(state *State) error {
	g.player.Update()
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}
