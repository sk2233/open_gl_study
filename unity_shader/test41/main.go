/*
@author: sk
@date: 2023/6/17
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
	reimu                     *frame.Obj
	reimuShader, effectShader *frame.Shader
	eye                       *frame.Camera
	effect                    *frame.PostEffect
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	temp := utils.LoadObj("test41/reisen.obj")
	m.reimu = frame.NewObj(temp, int32(len(temp)/8))
	m.reimu.VertexAttr(0, 3, 8*4, 0)
	m.reimu.VertexAttr(1, 2, 8*4, 3*4)
	m.reimu.VertexAttr(2, 3, 8*4, 5*4)
	m.reimu.BindTexture(gl.TEXTURE0, utils.LoadTexture("test41/reisen.png"))

	m.reimuShader = frame.NewShader("test41/test")
	m.reimuShader.Use()
	p := utils.GetDefaultPerspective()
	m.reimuShader.SetMat4("uProjection", &p[0])

	m.effectShader = frame.NewShader("test41/effect")
	m.effectShader.Use()
	m.effectShader.Set1i("uTexture0", 0)
	m.effectShader.Set1i("uTexture1", 1)

	m.eye = frame.NewCamera()
	m.effect = frame.NewPostEffect(1280, 720, 2)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.effect.Begin()

	m.reimuShader.Use()
	view := m.eye.GetView()
	m.reimuShader.SetMat4("uView", &view[0])
	m.reimu.Draw(m.reimuShader)

	m.effect.End()
	m.effectShader.Use()
	m.effect.Draw()
}
