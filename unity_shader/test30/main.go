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
	cube, plane          *frame.Obj
	depth, shadow        *frame.Shader
	screenBuf, screenTex uint32
	eye                  *frame.Camera
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.screenBuf, m.screenTex = utils.NewDepTextureBuff(1024, 1024)

	m.cube = frame.NewObj(cubeVs, 36)
	m.cube.VertexAttr(0, 3, 5*4, 0)
	m.cube.Translate(0, 1, 0)
	m.cube.BindTexture(gl.TEXTURE0, m.screenTex)
	m.cube.BindTexture(gl.TEXTURE1, utils.LoadTexture("unity_shader/test30/test.png"))

	m.plane = frame.NewObj(planeVs, 6)
	m.plane.VertexAttr(0, 3, 3*4, 0)
	m.plane.BindTexture(gl.TEXTURE0, m.screenTex)

	m.depth = frame.NewShader("unity_shader/test30/depth")
	m.depth.Use()
	lightView := mgl32.LookAtV(mgl32.Vec3{3, 2, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	m.depth.SetMat4("uView", &lightView[0])
	lightPro := mgl32.Ortho(-5, 5, -5, 5, 0.1, 30)
	m.depth.SetMat4("uProjection", &lightPro[0])
	m.depth.Set1i("uTexture", 1)

	m.shadow = frame.NewShader("unity_shader/test30/shadow")
	m.shadow.Use()
	eyePro := utils.GetDefaultPerspective()
	m.shadow.SetMat4("uProjection", &eyePro[0])
	temp := lightPro.Mul4(lightView)
	m.shadow.SetMat4("uLight", &temp[0])

	m.eye = frame.NewCamera()
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	gl.Viewport(0, 0, 1024, 1024)
	gl.BindFramebuffer(gl.FRAMEBUFFER, m.screenBuf)
	gl.Clear(gl.DEPTH_BUFFER_BIT)
	m.depth.Use()
	m.depth.Set1i("uOpen", 1)
	m.cube.Draw(m.depth)
	m.depth.Set1i("uOpen", 0)
	m.plane.Draw(m.depth)

	gl.Viewport(0, 0, 1280*2, 720*2)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	m.shadow.Use()
	view := m.eye.GetView()
	m.shadow.SetMat4("uView", &view[0])
	m.cube.Draw(m.shadow)
	m.plane.Draw(m.shadow)
}
