/*
@author: sk
@date: 2023/5/27
*/
package main

import (
	"openGL/frame3"

	"github.com/go-gl/mathgl/mgl32"
)

// 宽 80  高  45
type Brick struct {
	sprite *frame3.Sprite
	hp     int
	pos    complex64
	color  mgl32.Vec4
}

func (b *Brick) GetPos() complex64 {
	return b.pos
}

func (b *Brick) GetSize() complex64 {
	return 80 + 45i
}

func NewBrick(sprite *frame3.Sprite, hp int, pos complex64) *Brick {
	res := &Brick{sprite: sprite, hp: hp, pos: pos, color: frame3.ColorWhite}
	res.updateColor()
	return res
}

func (b *Brick) Draw() {
	b.sprite.Draw(b.pos, 1+1i, 0, b.color)
}

func (b *Brick) updateColor() {
	if b.hp > 0 {
		b.color = BrickColor[b.hp-1]
	}
}

func (b *Brick) Hurt() {
	if b.hp == 0 {
		return
	}
	b.hp--
	b.updateColor()
}

func (b *Brick) IsDie() bool {
	return b.hp == 0
}

type BrickContainer struct {
	bricks []*Brick
}

func (c *BrickContainer) Draw() {
	for _, brick := range c.bricks {
		brick.Draw()
	}
}

func (c *BrickContainer) Collision(rect frame3.IRect) *Brick {
	for _, brick := range c.bricks {
		if frame3.CollisionRect(rect, brick) {
			return brick
		}
	}
	return nil
}

func (c *BrickContainer) Update() {
	temp := make([]*Brick, 0, len(c.bricks))
	for _, brick := range c.bricks {
		if !brick.IsDie() {
			temp = append(temp, brick)
		}
	}
	c.bricks = temp
}

func NewBrickContainer(data [][]int, shader *frame3.Shader) *BrickContainer {
	res := &BrickContainer{bricks: make([]*Brick, 0)}
	brick := frame3.NewSprite(shader, frame3.NewTexture("game/img/brick.png"))
	wall := frame3.NewSprite(shader, frame3.NewTexture("game/img/wall.png"))
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] > 0 { // 砖块
				res.bricks = append(res.bricks, NewBrick(brick, data[i][j],
					complex(float32(i*80), float32(j*45))))
			} else if data[i][j] < 0 { // 铁墙
				res.bricks = append(res.bricks, NewBrick(wall, data[i][j],
					complex(float32(i*80), float32(j*45))))
			}
		}
	}
	return res
}
