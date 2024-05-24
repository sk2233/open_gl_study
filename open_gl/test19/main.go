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
	objShader *frame.Shader
	obj       *frame.Obj
	offset    float32
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.objShader = frame.NewShader("test19/obj")
	m.objShader.Use()
	m.objShader.Set1i("uImage", 0)
	pro := mgl32.Ortho2D(0, 1280, 720, 0)
	m.objShader.SetMat4("uPro", &pro[0])

	m.obj = frame.NewObj(rectVs, 6)
	m.obj.VertexAttr(0, 3, 5*4, 0)
	m.obj.VertexAttr(1, 2, 5*4, 3*4)
	m.obj.BindTexture(gl.TEXTURE0, utils.LoadTexture("test19/strip9.png"))
}

func (m *MainGame) Update(window *glfw.Window) {
	if window.GetKey(glfw.KeyEnter) == glfw.Release {
		m.offset += 1.0 / 9
		if m.offset >= 1 {
			m.offset = 0
		}
	}
}

func (m *MainGame) Draw() {
	m.objShader.Use()
	m.objShader.Set1f("uOffset", m.offset)

	m.obj.Reset()
	m.obj.Translate(100, 100, 0)
	m.obj.Draw(m.objShader)
}
