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
	reimu        *frame.Obj
	shader, post *frame.Shader
	eye          *frame.Camera
	effect       *frame.PostEffect
	index        int32
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	temp := utils.LoadObj("test37/reimu.obj")
	m.reimu = frame.NewObj(temp, int32(len(temp)/8))
	m.reimu.VertexAttr(0, 3, 8*4, 0)
	m.reimu.VertexAttr(1, 2, 8*4, 3*4)
	m.reimu.BindTexture(gl.TEXTURE0, utils.LoadTexture("test37/reimu.png"))

	m.shader = frame.NewShader("test37/test")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])
	m.shader.Set1i("uTexture", 0)

	m.post = frame.NewShader("test37/effect")
	m.post.Set1i("uImage0", 0)
	m.post.Set1i("uImage1", 1)
	m.post.Set1i("uImage2", 2)
	m.post.Set1i("uImage3", 3)
	m.post.Set1i("uImage4", 4)

	m.eye = frame.NewCamera()
	m.effect = frame.NewPostEffect(1280*2, 720*2, 2)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
	m.index = (m.index + 1) % 5
}

func (m *MainGame) Draw() {
	m.effect.Begin()

	m.shader.Use()
	m.shader.Set1i("uIndex", m.index)
	view := m.eye.GetView()
	m.shader.SetMat4("uView", &view[0])
	m.reimu.Draw(m.shader)

	m.effect.End()
	m.post.Use()
	m.post.Set1i("uIndex", m.index)
	m.effect.Draw()
}
