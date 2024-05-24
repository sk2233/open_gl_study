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
	cube   *frame.Obj
	shader *frame.Shader
	eye    *frame.Camera
	rate   float32
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.cube = frame.NewObj(cubeVs, 36)
	m.cube.VertexAttr(0, 3, 8*4, 0)
	m.cube.VertexAttr(1, 3, 8*4, 3*4)
	m.cube.VertexAttr(2, 2, 8*4, 6*4)
	m.cube.BindTexture(gl.TEXTURE0, utils.LoadTexture("test43/Crate_Diffuse.jpg"))
	m.cube.BindTexture(gl.TEXTURE1, utils.LoadTexture("test43/burn_noise.png"))

	m.shader = frame.NewShader("test43/test")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])
	m.shader.Set1i("uTexture", 0)
	m.shader.Set1i("uNoise", 1)

	m.eye = frame.NewCamera()
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
	m.rate += 0.01
	if m.rate > 1 {
		m.rate = 0
	}
}

func (m *MainGame) Draw() {
	m.shader.Use()
	view := m.eye.GetView()
	m.shader.SetMat4("uView", &view[0])
	m.shader.Set1f("uRate", m.rate)
	m.cube.Draw(m.shader)
}
