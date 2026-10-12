package assets

import (
	"embed"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var PlayerSprite *ebiten.Image = mustLoadSprite("assets/player.png")

func mustLoadSprite(name string) *ebiten.Image {
	//打开文件
	f, err := os.Open(name)
	if err != nil {
		log.Fatal("mustLoadSprite os.open error:", err.Error())
	}
	defer f.Close()
	//转换格式
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("mustLoadSprite image.Decode error:", err.Error())
	}
	return ebiten.NewImageFromImage(img)
	//返回
}
