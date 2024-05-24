/*
@author: sk
@date: 2023/6/17
*/
package main

import (
	frame "openGL/frame2"
	"openGL/utils"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	reimu       *frame.Obj
	reimuShader *frame.Shader
	eye         *frame.Camera
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	temp := utils.LoadObj("test39/reimu.obj")
	m.reimu = frame.NewObj(temp, int32(len(temp)/8))
	m.reimu.VertexAttr(0, 3, 8*4, 0)
	m.reimu.VertexAttr(1, 3, 8*4, 5*4)

	m.reimuShader = frame.NewShader("test39/test")
	m.reimuShader.Use()
	p := utils.GetDefaultPerspective()
	m.reimuShader.SetMat4("uProjection", &p[0])

	m.eye = frame.NewCamera()
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.reimuShader.Use()
	view := m.eye.GetView()
	m.reimuShader.SetMat4("uView", &view[0])
	m.reimu.Draw(m.reimuShader)
}
