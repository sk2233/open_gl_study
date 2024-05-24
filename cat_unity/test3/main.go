/*
@author: sk
@date: 2023/6/3
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
	rect   *frame.Obj
	shader *frame.Shader
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.rect = frame.NewObj(screenVs, 6)
	m.rect.VertexAttr(0, 3, 5*4, 0)
	m.rect.VertexAttr(1, 2, 5*4, 3*4)
	m.rect.BindTexture(gl.TEXTURE0, utils.LoadTexture("cat_unity/test2/player_stop_0.png"))

	m.shader = frame.NewShader("cat_unity/test3/base")
	m.shader.Use()
	p := utils.GetDefaultOrtho2D()
	m.shader.SetMat4("uProjection", &p[0])
	m0 := mgl32.Scale3D(43*2, 63*2, 1)
	m.shader.SetMat4("uModel", &m0[0])
}

func (m *MainGame) Update(window *glfw.Window) {

}

func (m *MainGame) Draw() {
	m.shader.Use()
	m.shader.Set1f("uTime", float32(glfw.GetTime()))
	m.rect.Draw(nil)
}
