/*
@author: sk
@date: 2023/6/10
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
	rect   *frame.Obj
	shader *frame.Shader
	eye    *frame.Camera
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.eye = frame.NewCamera()

	m.rect = frame.NewObj(rectVs, 6)
	m.rect.VertexAttr(0, 3, 5*4, 0)
	m.rect.VertexAttr(1, 2, 5*4, 3*4)
	m.rect.BindTexture(gl.TEXTURE0, utils.LoadTexture("test28/test.png"))

	m.shader = frame.NewShader("test28/test")
	m.shader.Use()
	pro := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &pro[0])

	gl.Enable(gl.BLEND)                                // 打开混合函数
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // 混合全部颜色
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	view := m.eye.GetView()
	m.shader.SetMat4("uView", &view[0])
	m.rect.Reset()
	m.rect.Draw(m.shader)
	m.rect.Translate(0.2, 0.2, 0.2)
	m.rect.Draw(m.shader)
	m.rect.Translate(0.4, 0.4, 0.4)
	m.rect.Draw(m.shader)
}
