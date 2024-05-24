/*
@author: sk
@date: 2023/5/14
*/
package main

import (
	frame "openGL/frame2"

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
	glfw.WindowHint(glfw.Samples, 4)
	gl.Enable(gl.MULTISAMPLE)
	m.shader = frame.NewShader("test15/obj")
	m.obj = frame.NewObj(rVs, 6)
	m.obj.VertexAttr(0, 2, 5*4, 0)
	m.obj.VertexAttr(1, 3, 5*4, 2*4)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	m.obj.DrawInstanced(nil, 100000)
}
