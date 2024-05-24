/*
@author: sk
@date: 2023/5/27
*/
package main

import (
	"openGL/frame3"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	frame3.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	bg           *frame3.Sprite
	bricks       *BrickContainer
	player       *Player
	ball         *Ball
	postEffect   *frame3.PostEffect
	effectShader *frame3.Shader
	bgm          *frame3.Music
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	Init()
	base := frame3.NewShader("frame3/shader/base")
	m.bg = frame3.NewSprite(base, frame3.NewTexture("game/img/bg.png"))
	m.bricks = NewBrickContainer(BrickData, base)
	m.player = NewPlayer(base)
	m.ball = NewBall(base)
	// ????
	m.effectShader = frame3.NewShader("game/shader/effect")
	m.postEffect = frame3.NewPostEffect(1280*2, 720*2, 1)
	m.bgm = frame3.NewMusic("game/audio/breakout.mp3")
	m.bgm.Play()
}

func (m *MainGame) Update(window *glfw.Window) {
	m.bricks.Update()
	m.player.Update(window)
	m.ball.Update(m.bricks, m.player)
}

func (m *MainGame) Draw() {
	m.postEffect.Begin()
	m.bg.Draw(0, 1+1i, 0, frame3.ColorWhite)
	m.bricks.Draw()
	m.player.Draw()
	m.ball.Draw()
	m.postEffect.End()
	m.effectShader.Use()
	m.effectShader.Set1f("uTime", float32(glfw.GetTime()))
	m.postEffect.Draw()
}
