/*
@author: sk
@date: 2023/6/4
*/
package main

import (
	frame "openGL/frame2"
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	// 必须使用 不清除模式 运行 该例子
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	reimu, rect        *frame.Obj
	shader, rectShader *frame.Shader
	eye                *frame.Camera
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	temp := utils.LoadObj("test38/reimu.obj")
	m.reimu = frame.NewObj(temp, int32(len(temp)/8))
	m.reimu.VertexAttr(0, 3, 8*4, 0)
	m.reimu.VertexAttr(1, 2, 8*4, 3*4)
	m.reimu.BindTexture(gl.TEXTURE0, utils.LoadTexture("test38/reimu.png"))

	m.rect = frame.NewObj(screenVs, 6)
	m.rect.VertexAttr(0, 3, 5*4, 0)

	m.shader = frame.NewShader("test38/test")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])
	m.shader.Set1i("uTexture", 0)

	m.rectShader = frame.NewShader("test38/test1")

	m.eye = frame.NewCamera()

	// 混合 半透明
	gl.Enable(gl.BLEND)                                // 打开混合函数
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // 混合全部颜色
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	// 绘制时忽略深度 直接覆盖
	gl.Disable(gl.DEPTH_TEST)
	m.rectShader.Use()
	m.rect.Draw(nil)
	gl.Enable(gl.DEPTH_TEST)

	m.shader.Use()
	view := m.eye.GetView()
	m.shader.SetMat4("uView", &view[0])
	m.reimu.Draw(m.shader)
}
