/*
@author: sk
@date: 2023/6/17
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
	reimu       *frame.Obj
	reimuShader *frame.Shader
	eye         *frame.Camera
	lastView    mgl32.Mat4
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	temp := utils.LoadObj("test40/reimu.obj")
	m.reimu = frame.NewObj(temp, int32(len(temp)/8))
	m.reimu.VertexAttr(0, 3, 8*4, 0)
	m.reimu.VertexAttr(1, 2, 8*4, 3*4)
	m.reimu.BindTexture(gl.TEXTURE0, utils.LoadTexture("test40/reimu.png"))
	m.reimu.BindTexture(gl.TEXTURE1, utils.LoadTexture("test40/fog_noise.jpg"))

	m.reimuShader = frame.NewShader("test40/test")
	m.reimuShader.Use()
	p := utils.GetDefaultPerspective()
	m.reimuShader.SetMat4("uProjection", &p[0])
	m.reimuShader.Set1i("uNoise", 1)

	m.eye = frame.NewCamera()
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.reimuShader.Use()
	view := m.eye.GetView()
	m.reimuShader.SetMat4("uView", &view[0])
	// 通过记录传入 上一次的 相机信息 获取一个点 上一次 与相机的相对位置
	m.reimuShader.SetMat4("uLastView", &m.lastView[0])
	m.reimu.Draw(m.reimuShader)
	m.lastView = view
}
