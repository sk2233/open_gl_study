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
	ebo, vao uint32
	shader   *frame.Shader
	counter  *frame.Counter
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	gl.GenVertexArrays(1, &m.vao)
	gl.BindVertexArray(m.vao)

	//gl.GenBuffers(1, &m.ebo)
	//gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ebo)
	//gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(rectIndex)*4, gl.Ptr(rectIndex), gl.STATIC_DRAW)

	var temp uint32
	gl.GenBuffers(1, &temp)
	gl.BindBuffer(gl.ARRAY_BUFFER, temp)
	gl.BufferData(gl.ARRAY_BUFFER, len(posVs)*4, gl.Ptr(posVs), gl.STATIC_DRAW)
	gl.VertexAttribPointerWithOffset(0, 4, gl.FLOAT, false, 4*4, 0)
	gl.EnableVertexAttribArray(0)

	gl.GenBuffers(1, &temp)
	gl.BindBuffer(gl.ARRAY_BUFFER, temp)
	gl.BufferData(gl.ARRAY_BUFFER, len(colVs)*4, gl.Ptr(colVs), gl.STATIC_DRAW)
	gl.VertexAttribPointerWithOffset(1, 4, gl.FLOAT, false, 4*4, 0)
	gl.EnableVertexAttribArray(1)

	gl.GenBuffers(2, &temp)
	gl.BindBuffer(gl.ARRAY_BUFFER, temp)
	gl.BufferData(gl.ARRAY_BUFFER, len(offset)*4, gl.Ptr(offset), gl.STATIC_DRAW)
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, 2*4, 0)
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribDivisor(2, 1)

	m.shader = frame.NewShader("test47/obj")
	m.counter = frame.NewCounter()
}

func (m *MainGame) Update(window *glfw.Window) {
	m.counter.Update()
}

func (m *MainGame) Draw() {
	m.shader.Use()

	gl.BindVertexArray(m.vao)
	gl.DrawArraysInstanced(gl.TRIANGLE_STRIP, 0, 4, 3)
}
