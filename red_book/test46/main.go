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
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	ebo, abo, vao uint32
	shader        *frame.Shader
	eye           *frame.Camera
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	gl.GenVertexArrays(1, &m.vao)
	gl.BindVertexArray(m.vao)

	gl.GenBuffers(1, &m.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(rectIndex)*4, gl.Ptr(rectIndex), gl.STATIC_DRAW)

	gl.GenBuffers(1, &m.abo)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.abo)
	gl.BufferData(gl.ARRAY_BUFFER, len(rectVs)*4, gl.Ptr(rectVs), gl.STATIC_DRAW)

	gl.VertexAttribPointerWithOffset(0, 4, gl.FLOAT, false, 8*4, 0)
	gl.VertexAttribPointerWithOffset(1, 4, gl.FLOAT, false, 8*4, 4*4)
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)

	m.shader = frame.NewShader("test46/obj")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uP", &p[0])
	m0 := mgl32.Ident4()
	m.shader.SetMat4("uM", &m0[0])
	m.eye = frame.NewCamera()

	gl.Enable(gl.PRIMITIVE_RESTART)
	gl.PrimitiveRestartIndex(0xffff)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	v := m.eye.GetView()
	m.shader.SetMat4("uV", &v[0])

	gl.BindVertexArray(m.vao)
	gl.DrawElements(gl.TRIANGLE_STRIP, 17, gl.UNSIGNED_INT, nil)
}
