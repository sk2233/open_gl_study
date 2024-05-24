/*
@author: sk
@date: 2023/6/24
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
	eye                  *frame.Camera
	obj                  *frame.Obj
	shader, effectShader *frame.Shader
	effect               *frame.PostEffect
}

func NewMainGame() *MainGame {
	return &MainGame{eye: frame.NewCamera()}
}

func (m *MainGame) Init(window *glfw.Window) {
	vs := utils.LoadObj("rabbit_cocos/test3/reisen.obj")
	m.obj = frame.NewObj(vs, int32(len(vs)/8))
	m.obj.VertexAttr(0, 3, 4*8, 0)
	m.obj.VertexAttr(1, 2, 4*8, 4*3)
	m.obj.BindTexture(gl.TEXTURE0, utils.LoadTexture("rabbit_cocos/test3/reisen.png"))

	m.shader = frame.NewShader("rabbit_cocos/test3/test")
	m.shader.Use()
	p := utils.GetDefaultPerspective()
	m.shader.SetMat4("uProjection", &p[0])

	m.effect = frame.NewPostEffect(1280, 720, 2)
	m.effectShader = frame.NewShader("rabbit_cocos/test3/effect")
	m.effectShader.Use()
	m.effectShader.Set1i("uTexture", 0)
	m.effectShader.Set1i("uDep", 1)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.effect.Begin()
	m.shader.Use()
	v := m.eye.GetView()
	m.shader.SetMat4("uView", &v[0])
	m.obj.Draw(m.shader)
	m.effect.End()

	m.effectShader.Use()
	m.effect.Draw()
}
