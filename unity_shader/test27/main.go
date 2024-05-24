/*
@author: sk
@date: 2023/6/10
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
	rect     *frame.Obj
	shader   *frame.Shader
	lightPos mgl32.Vec3
	eye      *frame.Camera
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.lightPos = mgl32.Vec3{5, 5, 5}
	m.eye = frame.NewCamera()

	m.rect = frame.NewObj(rectVs, 6)
	m.rect.VertexAttr(0, 3, 5*4, 0)
	m.rect.VertexAttr(1, 2, 5*4, 3*4)
	m.rect.BindTexture(gl.TEXTURE0, utils.LoadTexture("test27/diffuse.jpg"))
	m.rect.BindTexture(gl.TEXTURE1, utils.LoadTexture("test27/normal.jpg"))

	m.shader = frame.NewShader("test27/test")
	m.shader.Use()
	m.shader.Set1i("uTexture", 0)
	m.shader.Set1i("uNormal", 1)
	m.shader.Set3fv("uLightPos", &m.lightPos[0])
	pro := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &pro[0])
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	view := m.eye.GetView()
	m.shader.SetMat4("uView", &view[0])
	pos := m.eye.GetPos()
	m.shader.Set3fv("uEyePos", &pos[0])
	m.rect.Draw(m.shader)
}
