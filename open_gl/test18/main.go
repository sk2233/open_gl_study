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
	camera    *frame.Camera
	objShader *frame.Shader
	obj       *frame.Obj
	lightPos  mgl32.Vec3
}

func NewMainGame() *MainGame {
	return &MainGame{camera: frame.NewCamera()}
}

func (m *MainGame) Init(window *glfw.Window) {
	//fmt.Println(gl.GetError())
	m.lightPos = mgl32.Vec3{4, 4, 4}

	m.objShader = frame.NewShader("test18/obj")
	m.objShader.Use()
	m.objShader.Set1i("uImage", 0)
	pro := utils.GetDefaultPerspective()
	m.objShader.SetMat4("uPro", &pro[0])
	m.objShader.Set3fv("uLightPos", &m.lightPos[0])

	m.obj = frame.NewObj(rectVs, 6)
	m.obj.VertexAttr(0, 3, 5*4, 0)
	m.obj.VertexAttr(1, 2, 5*4, 3*4)
	m.obj.BindTexture(gl.TEXTURE0, utils.LoadTexture("test18/cube.png"))
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	m.objShader.Use()
	view := m.camera.GetView()
	m.objShader.SetMat4("uView", &view[0])
	pos := m.camera.GetPos()
	m.objShader.Set3fv("uEyePos", &pos[0])

	m.obj.Draw(nil)
}
