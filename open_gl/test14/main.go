/*
@author: sk
@date: 2023/5/14
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
	camera *frame.Camera
	shader *frame.Shader
	obj    *frame.Obj
}

func NewMainGame() *MainGame {
	return &MainGame{camera: frame.NewCamera()}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.shader = frame.NewShader("test14/obj")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])
	vs := utils.LoadObj("test14/reimu.obj")
	m.obj = frame.NewObj(vs, int32(len(vs)/8))
	m.obj.VertexAttr(0, 3, 8*4, 0)
	m.obj.VertexAttr(1, 2, 8*4, 3*4)
	m.obj.VertexAttr(2, 3, 8*4, 5*4)
	temp := utils.LoadTexture("test14/reimu.png")
	m.obj.BindTexture(gl.TEXTURE0, temp)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	view := m.camera.GetView()
	m.shader.SetMat4("uView", &view[0])
	//m.shader.Set1f("uTime", float32(glfw.GetTime()))
	m.obj.Draw(m.shader)
}
