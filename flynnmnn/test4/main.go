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
)

func main() {
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	cube, plane    *frame.Obj
	shadow, effect *frame.Shader
	eye            *frame.Camera
	postEffect     *frame.PostEffect
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.eye = frame.NewCamera()
	m.postEffect = frame.NewPostEffect(1280, 720, 1)

	texture := utils.LoadTexture("flynnmnn/test4/test.png")
	eyePro := utils.GetDefaultPerspective()

	m.cube = frame.NewObj(cubeVs, 36)
	m.cube.VertexAttr(0, 3, 5*4, 0)
	m.cube.VertexAttr(1, 2, 5*4, 3*4)
	m.cube.BindTexture(gl.TEXTURE0, texture)
	m.cube.BindTexture(gl.TEXTURE1, m.postEffect.GetTexture(0)) // 深度纹理

	m.plane = frame.NewObj(planeVs, 6)
	m.plane.VertexAttr(0, 3, 5*4, 0)
	m.plane.VertexAttr(1, 2, 5*4, 3*4)
	m.plane.BindTexture(gl.TEXTURE0, texture)

	m.shadow = frame.NewShader("flynnmnn/test4/shadow")
	m.shadow.Use()
	m.shadow.SetMat4("uProjection", &eyePro[0])

	m.effect = frame.NewShader("flynnmnn/test4/effect")
	m.effect.Use()
	m.effect.SetMat4("uProjection", &eyePro[0])
	m.effect.Set1i("uImage", 0)
	m.effect.Set1i("uDep", 1)

	gl.Enable(gl.BLEND)                                // 打开混合函数
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // 混合全部颜色
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	view := m.eye.GetView()

	// 获取场景中除Cube外物体的深度
	m.postEffect.Begin()
	m.shadow.Use()
	m.shadow.SetMat4("uView", &view[0])
	m.plane.Draw(m.shadow)
	m.postEffect.End()

	m.effect.Use()
	m.effect.SetMat4("uView", &view[0])
	m.effect.Set1i("uOpen", gl.FALSE)
	m.plane.Draw(m.effect)
	m.effect.Set1i("uOpen", gl.TRUE)
	m.cube.Draw(m.effect)
}
