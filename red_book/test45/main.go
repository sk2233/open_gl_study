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
	ebo, abo, vao uint32
	shader        *frame.Shader
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

	m.shader = frame.NewShader("test45/obj")
}

func (m *MainGame) Update(window *glfw.Window) {

}

func (m *MainGame) Draw() {
	m.shader.Use()

	gl.BindVertexArray(m.vao)
	gl.DrawElements(gl.TRIANGLES, 3, gl.UNSIGNED_INT, nil)
}
