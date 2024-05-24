/*
@author: sk
@date: 2023/6/4
*/
package main

import (
	frame "openGL/frame2"
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	obj                  *frame.Obj
	shader, effectShader *frame.Shader
	eye                  *frame.Camera
	effect               *frame.PostEffect
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	vs := utils.LoadObj("red_book/test52/reisen.obj")
	m.obj = frame.NewObj(vs, int32(len(vs)/8))
	m.obj.VertexAttr(0, 3, 8*4, 0)
	m.obj.VertexAttr(1, 2, 8*4, 3*4)
	m.obj.BindTexture(gl.TEXTURE0, utils.LoadTexture("red_book/test52/reisen.png"))

	m.shader = frame.NewShader("red_book/test52/test")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])

	m.effect = frame.NewPostEffect(1280, 720, 1)
	m.effectShader = frame.NewShader("red_book/test52/effect")

	m.eye = frame.NewCamera()
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.effect.Begin()
	m.shader.Use()
	view := m.eye.GetView()
	m.shader.SetMat4("uView", &view[0])
	m.obj.Draw(m.shader)
	m.effect.End()

	m.effectShader.Use()
	m.effectShader.Set1f("uTime", float32(glfw.GetTime()))
	m.effect.Draw()
}
