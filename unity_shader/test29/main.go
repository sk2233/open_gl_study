/*
@author: sk
@date: 2023/6/4
*/
package main

import (
	"fmt"
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
	reimu              *frame.Obj
	shader             *frame.Shader
	eye                *frame.Camera
	lightPos, lightDir mgl32.Vec3
	threshold          float32
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.lightPos = mgl32.Vec3{5, 5, 5}
	m.lightDir = mgl32.Vec3{-1, -1, -1}
	m.threshold = 0.98

	temp := utils.LoadObj("test29/reimu.obj")
	m.reimu = frame.NewObj(temp, int32(len(temp)/8))
	m.reimu.VertexAttr(0, 3, 8*4, 0)
	m.reimu.VertexAttr(1, 2, 8*4, 3*4)
	m.reimu.VertexAttr(2, 3, 8*4, 5*4)
	m.reimu.BindTexture(gl.TEXTURE0, utils.LoadTexture("test29/reimu.png"))

	m.shader = frame.NewShader("test29/test")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])
	m.shader.Set3fv("uLightPos", &m.lightPos[0])
	m.shader.Set3fv("uLightDir", &m.lightDir[0])

	m.eye = frame.NewCamera()
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
	m.threshold += frame.GetAxis(window, glfw.KeyJ, glfw.KeyK) * 0.001
	fmt.Println(m.threshold)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	m.shader.Set1f("uThreshold", m.threshold)
	view := m.eye.GetView()
	m.shader.SetMat4("uView", &view[0])
	pos := m.eye.GetPos()
	m.shader.Set3fv("uEyePos", &pos[0])
	m.reimu.Draw(m.shader)
}
