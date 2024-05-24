/*
@author: sk
@date: 2023/5/27
*/
package main

import "openGL/frame3"

type Ball struct {
	sprite *frame3.Sprite
	pos    complex64
	speed  complex64
}

func (b *Ball) GetPos() complex64 {
	return b.pos
}

func (b *Ball) GetSize() complex64 {
	return 32 + 32i
}

func (b *Ball) Draw() {
	b.sprite.Draw(b.pos, 1+1i, 0, frame3.ColorWhite)
}

func (b *Ball) Update(bricks *BrickContainer, player *Player) {
	b.pos += b.speed
	b.checkBorder()
	b.checkBricks(bricks)
	b.checkPlayer(player)
}

func (b *Ball) checkBorder() {
	x, y := frame3.Vec2Float(b.pos)
	if x < 0 || x > 1280-32 {
		b.speed = frame3.VecMul(b.speed, -1+1i)
	}
	if y < 0 || y > 720-32 {
		b.speed = frame3.VecMul(b.speed, 1-1i)
	}
}

func (b *Ball) checkBricks(bricks *BrickContainer) {
	brick := bricks.Collision(b)
	if brick == nil {
		return
	}
	b.pos -= b.speed * 2 // 回退
	brick.Hurt()
	if b.overlayX(brick) {
		b.speed = frame3.VecMul(b.speed, 1-1i)
	} else {
		b.speed = frame3.VecMul(b.speed, -1+1i)
	}
}

func (b *Ball) checkPlayer(player *Player) {
	if frame3.CollisionRect(b, player) {
		b.speed = frame3.VecMul(b.speed, 1-1i)
	}
}

func (b *Ball) overlayX(rect frame3.IRect) bool {
	if real(b.pos) > real(rect.GetPos()+rect.GetSize()) {
		return false
	}
	if real(rect.GetPos()) > real(b.pos+b.GetSize()) {
		return false
	}
	return true
}

func NewBall(shader *frame3.Shader) *Ball {
	return &Ball{sprite: frame3.NewSprite(shader, frame3.NewTexture("game/img/ball.png")), pos: 1280/2 + 720i/2,
		speed: 1 + 1i}
}
