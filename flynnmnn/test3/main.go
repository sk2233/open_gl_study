/*
@author: sk
@date: 2023/5/29
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
	cube, plane    *frame.Obj
	shadow, effect *frame.Shader
	eye            *frame.Camera
	postEffect     *frame.PostEffect
	pos            mgl32.Vec2
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	texture0 := utils.LoadTexture("flynnmnn/test3/test.png")
	texture1 := utils.LoadTexture("flynnmnn/test3/bg.png")

	m.cube = frame.NewObj(cubeVs, 36)
	m.cube.VertexAttr(0, 3, 5*4, 0)
	m.cube.VertexAttr(1, 2, 5*4, 3*4)
	m.cube.Translate(0, 1, 0)
	m.cube.BindTexture(gl.TEXTURE0, texture0)
	m.cube.BindTexture(gl.TEXTURE1, texture1)

	m.plane = frame.NewObj(planeVs, 6)
	m.plane.VertexAttr(0, 3, 5*4, 0)
	m.plane.VertexAttr(1, 2, 5*4, 3*4)
	m.plane.BindTexture(gl.TEXTURE0, texture0)
	m.cube.BindTexture(gl.TEXTURE1, texture1)

	m.shadow = frame.NewShader("flynnmnn/test3/shadow")
	m.shadow.Use()
	eyePro := utils.GetDefaultPerspective()
	m.shadow.SetMat4("uProjection", &eyePro[0])
	m.shadow.Set1i("uImage", 0)
	m.shadow.Set1i("uMask", 1)
	m.effect = frame.NewShader("flynnmnn/test3/effect")

	m.eye = frame.NewCamera()
	m.postEffect = frame.NewPostEffect(1280, 720, 2)
	m.pos = mgl32.Vec2{0.0, 0.0} // 向外 扩展 0.5

	gl.Enable(gl.BLEND)                                // 打开混合函数
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // 混合全部颜色
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
	x := frame.GetAxis(window, glfw.KeyI, glfw.KeyK) * 0.1
	y := frame.GetAxis(window, glfw.KeyL, glfw.KeyJ) * 0.1
	m.pos = m.pos.Add(mgl32.Vec2{x, y})
}

func (m *MainGame) Draw() {
	//m.postEffect.Begin()
	m.shadow.Use()
	view := m.eye.GetView()
	m.shadow.SetMat4("uView", &view[0])
	// 同样使用 世界坐标 这里 使用拦截墙面效果
	m.shadow.Set2fv("uPos", &m.pos[0])
	m.shadow.Set1f("uLen", float32(glfw.GetTime()))
	m.cube.Draw(m.shadow)
	m.plane.Draw(m.shadow)
	//m.postEffect.End()
	//
	//m.effect.Use()
	//m.postEffect.Draw()
}
