/*
@author: sk
@date: 2023/6/11
*/
package main

import (
	frame "openGL/frame2"
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	frame.Run(NewMainApp(), 1280, 720)
}

type MainApp struct {
	obj          *frame.Obj
	shader, post *frame.Shader
	effect       *frame.PostEffect
}

func (m *MainApp) Init(window *glfw.Window) {
	m.obj = frame.NewObj(rectVs, 6)
	m.obj.VertexAttr(0, 3, 5*4, 0)
	m.obj.VertexAttr(1, 2, 5*4, 3*4)
	m.obj.BindTexture(gl.TEXTURE0, utils.LoadTexture("test35/sakura0.jpg"))

	m.shader = frame.NewShader("test35/test")
	m.post = frame.NewShader("test35/effect")

	m.effect = frame.NewPostEffect(1280, 720, 1)
}

func (m *MainApp) Update(window *glfw.Window) {
}

func (m *MainApp) Draw() {
	m.effect.Begin()
	m.shader.Use()
	m.obj.Draw(nil)
	m.effect.End()

	m.post.Use()
	m.effect.Draw()
}

func NewMainApp() *MainApp {
	return &MainApp{}
}
