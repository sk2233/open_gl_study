/*
@author: sk
@date: 2023/6/3
*/
package main

import (
	frame "openGL/frame2"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	rect   *frame.Obj
	shader *frame.Shader
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.rect = frame.NewObj(screenVs, 6)
	m.rect.VertexAttr(0, 3, 5*4, 0)
	m.rect.VertexAttr(1, 2, 5*4, 3*4)
	m.shader = frame.NewShader("test24/base")
}

func (m *MainGame) Update(window *glfw.Window) {

}

func (m *MainGame) Draw() {
	m.shader.Use()
	m.rect.Draw(nil)
}
