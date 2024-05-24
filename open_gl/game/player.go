/*
@author: sk
@date: 2023/5/27
*/
package main

import (
	"openGL/frame3"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Player struct { // 128 * 32
	sprite *frame3.Sprite
	x      float32
}

func (p *Player) GetPos() complex64 {
	return complex(p.x, 720-32)
}

func (p *Player) GetSize() complex64 {
	return 128 + 32i
}

func (p *Player) Draw() {
	p.sprite.Draw(complex(p.x, 720-32), 1+1i, 0, frame3.ColorWhite)
}

func (p *Player) Update(window *glfw.Window) {
	p.x += frame3.GetAxle(window, glfw.KeyA, glfw.KeyD) * 3
	if p.x > 1280-128 {
		p.x = 1280 - 128
	} else if p.x < 0 {
		p.x = 0
	}
}

func NewPlayer(shader *frame3.Shader) *Player {
	return &Player{sprite: frame3.NewSprite(shader, frame3.NewTexture("game/img/player.png")), x: 0}
}
