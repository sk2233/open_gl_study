/*
@author: sk
@date: 2023/5/27
*/
package main

import (
	"openGL/frame3"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	frame3.Run(NewHDRGame(), 1280, 720)
}

type HDRGame struct {
	sprite    *frame3.Sprite
	hdrColor  mgl32.Vec4
	hdrEffect *PostEffect
}

func NewHDRGame() *HDRGame {
	return &HDRGame{}
}

func (h *HDRGame) Init(window *glfw.Window) {
	shader := frame3.NewShader("test20/shader/base")
	h.sprite = frame3.NewSprite(shader, frame3.NewTexture("test20/test.png"))
	// 颜色超过1  需要开启 HDR
	h.hdrColor = mgl32.Vec4{1, 2, 1, 1}
	h.hdrEffect = NewPostEffect(1280*2, 720*2, "test20/shader/effect")
}

func (h *HDRGame) Update(window *glfw.Window) {

}

func (h *HDRGame) Draw() {
	h.hdrEffect.Begin()
	h.sprite.Draw(100+100i, 5+5i, 0, h.hdrColor)
	h.hdrEffect.End()
	h.hdrEffect.Draw()
}
