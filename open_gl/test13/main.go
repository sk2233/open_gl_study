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
}

func NewMainGame() *MainGame {
	return &MainGame{camera: frame.NewCamera()}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.shader = frame.NewShader("test13/obj")
	utils.CreateVertexArray()
	utils.BindBuffer(pVs, gl.ARRAY_BUFFER)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 2, gl.FLOAT, false, 5*4, 0)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 5*4, 2*4)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	gl.DrawArrays(gl.POINTS, 0, 4)
}
