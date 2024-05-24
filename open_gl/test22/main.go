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
	cube, plane, screen         *frame.Obj
	obj, depth, screen0, shadow *frame.Shader
	lightPro, lightView, eyePro mgl32.Mat4
	screenBuf, screenTex        uint32
	camera                      *frame.Camera
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	gl.GenTextures(1, &m.screenTex)
	gl.BindTexture(gl.TEXTURE_2D, m.screenTex)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH_COMPONENT, 1024, 1024, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.GenFramebuffers(1, &m.screenBuf)
	gl.BindFramebuffer(gl.FRAMEBUFFER, m.screenBuf)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.TEXTURE_2D, m.screenTex, 0)
	gl.DrawBuffer(gl.NONE)
	gl.ReadBuffer(gl.NONE)

	m.cube = frame.NewObj(cubeVs, 36)
	m.cube.VertexAttr(0, 3, 3*4, 0)
	m.cube.Translate(0, 1, 0)
	m.cube.BindTexture(gl.TEXTURE0, m.screenTex)
	m.plane = frame.NewObj(planeVs, 6)
	m.plane.VertexAttr(0, 3, 3*4, 0)
	m.plane.BindTexture(gl.TEXTURE0, m.screenTex)
	m.screen = frame.NewObj(screenVs, 6)
	m.screen.VertexAttr(0, 3, 5*4, 0)
	m.screen.VertexAttr(1, 2, 5*4, 3*4)
	m.screen.BindTexture(gl.TEXTURE0, m.screenTex)

	m.obj = frame.NewShader("test22/obj")
	m.depth = frame.NewShader("test22/depth")
	m.screen0 = frame.NewShader("test22/screen")
	m.shadow = frame.NewShader("test22/shadow")
	m.lightPro = mgl32.Ortho(-10, 10, -10, 10, 0.1, 10)
	m.lightView = mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	m.eyePro = utils.GetDefaultPerspective()
	m.camera = frame.NewCamera()
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	gl.Viewport(0, 0, 1024, 1024)
	gl.BindFramebuffer(gl.FRAMEBUFFER, m.screenBuf)
	gl.Clear(gl.DEPTH_BUFFER_BIT)
	m.depth.Use()
	m.depth.SetMat4("uView", &m.lightView[0])
	m.depth.SetMat4("uProjection", &m.lightPro[0])
	m.cube.Draw(m.depth)
	m.plane.Draw(m.depth)

	gl.Viewport(0, 0, 1280*2, 720*2)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	m.shadow.Use()
	view := m.camera.GetView()
	m.shadow.SetMat4("uView", &view[0])
	m.shadow.SetMat4("uProjection", &m.eyePro[0])
	temp := m.lightPro.Mul4(m.lightView)
	m.shadow.SetMat4("uLight", &temp[0])
	m.cube.Draw(m.shadow)
	m.plane.Draw(m.shadow)
}
