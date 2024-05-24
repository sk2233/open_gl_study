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
	obj, plane, screen          *frame.Obj
	depth, shadow, effectShader *frame.Shader
	screenBuf, screenTex        uint32
	eye                         *frame.Camera
	effect                      *frame.PostEffect
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.screenBuf, m.screenTex = utils.NewDepTextureBuff(1024, 1024)

	vs := utils.LoadObj("rabbit_cocos/test5/reisen.obj")
	m.obj = frame.NewObj(vs, int32(len(vs)/8))
	m.obj.VertexAttr(0, 3, 4*8, 0)
	m.obj.VertexAttr(1, 2, 4*8, 4*3)
	m.obj.VertexAttr(2, 3, 4*8, 4*5)
	m.obj.BindTexture(gl.TEXTURE0, utils.LoadTexture("rabbit_cocos/test5/reisen.png"))
	m.obj.BindTexture(gl.TEXTURE1, m.screenTex)

	m.plane = frame.NewObj(planeVs, 6)
	m.plane.VertexAttr(0, 3, 4*8, 0)
	m.plane.VertexAttr(1, 2, 4*8, 4*3)
	m.plane.VertexAttr(2, 3, 4*8, 4*5)
	m.plane.BindTexture(gl.TEXTURE0, utils.LoadTexture("rabbit_cocos/test5/plane.jpg"))
	m.plane.BindTexture(gl.TEXTURE1, m.screenTex)

	m.depth = frame.NewShader("rabbit_cocos/test5/depth")
	m.depth.Use()
	lightView := mgl32.LookAtV(mgl32.Vec3{3, 2, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	m.depth.SetMat4("uView", &lightView[0])
	lightPro := mgl32.Ortho(-5, 5, -5, 5, 0.1, 30)
	m.depth.SetMat4("uProjection", &lightPro[0])

	m.shadow = frame.NewShader("rabbit_cocos/test5/shadow")
	m.shadow.Use()
	eyePro := utils.GetDefaultPerspective()
	m.shadow.SetMat4("uProjection", &eyePro[0])
	temp := lightPro.Mul4(lightView)
	m.shadow.SetMat4("uLight", &temp[0])
	m.shadow.Set1i("uTexture", 0)
	m.shadow.Set1i("uShadow", 1)

	m.eye = frame.NewCamera()
	m.effect = frame.NewPostEffect(1280, 720, 2)
	m.effectShader = frame.NewShader("rabbit_cocos/test5/effect")
	m.effectShader.Use()
	m.effectShader.Set1i("uTexture1", 0)
	m.effectShader.Set1i("uTexture1", 1)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	gl.Viewport(0, 0, 1024, 1024)
	gl.BindFramebuffer(gl.FRAMEBUFFER, m.screenBuf)
	gl.Clear(gl.DEPTH_BUFFER_BIT)
	m.depth.Use()
	m.obj.Draw(m.depth)
	m.plane.Draw(m.depth)

	gl.Viewport(0, 0, 1280*2, 720*2)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	m.effect.Begin()

	m.shadow.Use()
	view := m.eye.GetView()
	m.shadow.SetMat4("uView", &view[0])
	pos := m.eye.GetPos()
	m.shadow.Set3fv("uEye", &pos[0])
	m.shadow.Set1i("uOpen", 1)
	m.obj.Draw(m.shadow)
	m.shadow.Set1i("uOpen", 0)
	m.plane.Draw(m.shadow)

	m.effect.End()

	m.effectShader.Use()
	m.effect.Draw()
}
