package main

import (
	"go-asterrolds/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSpeed          = math.Pi   //每秒旋转的角度
	maxAcceleration           = 8.0       //最大加速度
	screenWidth, screenHeight = 1280, 720 //16:9
)

var curAcceleration float64 //当前加速度

type Player struct {
	game           *Game
	sprite         *ebiten.Image
	rotation       float64
	position       Vector  //位置
	playerVelocity float64 //玩家速度

}

func NewPlayer(g *Game) *Player {
	var p = &Player{}

	p.sprite = assets.PlayerSprite
	p.game = g
	//玩家居中屏幕
	bounds := p.sprite.Bounds()
	halfW, halfH := float64(bounds.Dx()/2), float64(bounds.Dy()/2)
	p.position = Vector{
		X: screenWidth/2 - halfW, //-halfW 可以让图片中心居中
		Y: screenHeight/2 - halfH,
	}
	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	bounds := p.sprite.Bounds()
	halfW, halfH := bounds.Dx()/2, bounds.Dy()/2 //角色精灵图的中心坐标

	op.GeoM.Translate(-float64(halfW), -float64(halfH))
	//旋转角色
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(float64(halfW), float64(halfH))
	//移动角色
	op.GeoM.Translate(p.position.X, p.position.Y)
	//绘制角色
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Update() {
	speed := rotationPerSpeed / float64(ebiten.TPS()) //每帧旋转的角度
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
	//角色加速
	p.accelerate()
}

//角色加速
func (p *Player) accelerate() {
	//如果按下了上键 则加速
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		//逐步加速 | 限制边界
		p.keepScreen()
		if curAcceleration < maxAcceleration {
			curAcceleration = p.playerVelocity + 4
		}
		if curAcceleration > 8 {
			curAcceleration = 8
		}
		p.playerVelocity = curAcceleration

		//计算dx和dy | dx是移动量？
		dx := math.Sin(p.rotation) * curAcceleration  //角色带方向移动
		dy := math.Cos(p.rotation) * -curAcceleration //这里为什么要取负？ 因为角色是面向上的，所以y轴是向下的，所以要取负号

		//更新位置
		p.position.X += dx
		p.position.Y += dy
	}
}

//角色保持在窗口内
func (p *Player) keepScreen() {
	//如果角色超出边界，则把角色放置到该边界对面出现
	if p.position.X >= screenWidth {
		p.position.X = 0
	} else if p.position.X <= 0 {
		p.position.X = screenWidth
	}
	if p.position.Y >= screenHeight {
		p.position.Y = 0
	} else if p.position.Y <= 0 {
		p.position.Y = screenHeight
	}
}
