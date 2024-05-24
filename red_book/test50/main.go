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
	obj    *frame.Obj
	shader *frame.Shader
	eye    *frame.Camera
}

func (m *MainGame) Init(window *glfw.Window) {
	m.obj = frame.NewObj(planVs, 6)
	m.obj.VertexAttr(0, 3, 8*4, 0)
	m.obj.VertexAttr(1, 2, 8*4, 6*4)
	m.obj.BindTexture3D(gl.TEXTURE0, utils.Load3DTexture("red_book/test50/sanying_bg.jpg", "red_book/test50/zhulin_bg.jpg"))

	m.shader = frame.NewShader("red_book/test50/obj")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])

	m.eye = frame.NewCamera()
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
