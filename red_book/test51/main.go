/*
@author: sk
@date: 2023/7/1
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
	shader *frame.Shader
	eye    *frame.Camera
	obj    *PointObj
}

func (m *MainGame) Init(window *glfw.Window) {
	m.obj = NewPointObj(planVs, 6)
	m.obj.VertexAttr(0, 3, 3*4, 0)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, utils.LoadTexture("red_book/test51/bg.png"))

	m.shader = frame.NewShader("red_book/test51/obj")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])

	m.eye = frame.NewCamera()
	gl.Enable(gl.PROGRAM_POINT_SIZE)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	v := m.eye.GetView()
	m.shader.SetMat4("uView", &v[0])
	m.obj.Draw(m.shader)
}

func NewMainGame() *MainGame {
	return &MainGame{}
}
