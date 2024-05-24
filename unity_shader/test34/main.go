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
	obj    *frame.Obj
	shader *frame.Shader
	eye    *frame.Camera
}

func (m *MainApp) Init(window *glfw.Window) {
	m.eye = frame.NewCamera()

	m.obj = frame.NewObj(rectVs, 6)
	m.obj.VertexAttr(0, 3, 5*4, 0)
	m.obj.VertexAttr(1, 2, 5*4, 3*4)
	m.obj.BindTexture(gl.TEXTURE0, utils.LoadTexture("test34/star.png"))

	m.shader = frame.NewShader("test34/test")
	m.shader.Use()
	pro := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &pro[0])

	gl.Enable(gl.BLEND)                                // 打开混合函数
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // 混合全部颜色
}

func (m *MainApp) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainApp) Draw() {
	m.shader.Use()
	view := m.eye.GetView()
	m.shader.SetMat4("uView", &view[0])
	// 获取相机 在模型空间的位置 传入 shader
	pos := m.eye.GetPos().Vec4(1)
	model := m.obj.GetModel()
	pos = model.Inv().Mul4x1(pos)
	temp := pos.Vec3()
	m.shader.Set3fv("uEye", &temp[0])
	m.obj.Draw(m.shader)
}

func NewMainApp() *MainApp {
	return &MainApp{}
}
