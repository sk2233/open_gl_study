/*
@author: sk
@date: 2023/5/14
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
	camera                      *frame.Camera
	depShader, debugShader      *frame.Shader
	planObj, cubeObj, screenObj *frame.Obj
	lightPos                    mgl32.Vec3
	frameBuf, textureBuf        uint32
}

func NewMainGame() *MainGame {
	return &MainGame{camera: frame.NewCamera()}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.lightPos = mgl32.Vec3{4, 4, 4}

	m.depShader = frame.NewShader("open_gl/test16/dep")
	m.debugShader = frame.NewShader("open_gl/test16/debug")

	temp := utils.LoadTexture("open_gl/test16/img.png")
	m.frameBuf, m.textureBuf = utils.NewDepTextureBuff(1024, 1024)

	m.planObj = frame.NewObj(planVs, 6)
	m.planObj.VertexAttr(0, 3, 8*4, 0)
	m.planObj.VertexAttr(1, 3, 8*4, 3*4)
	m.planObj.VertexAttr(2, 2, 8*4, 6*4)
	m.planObj.BindTexture(gl.TEXTURE0, temp)

	m.cubeObj = frame.NewObj(cubeVs, 36)
	m.cubeObj.VertexAttr(0, 3, 8*4, 0)
	m.cubeObj.VertexAttr(1, 3, 8*4, 3*4)
	m.cubeObj.VertexAttr(2, 2, 8*4, 6*4)
	m.cubeObj.BindTexture(gl.TEXTURE0, temp)

	m.screenObj = frame.NewObj(screenVs, 6)
	m.screenObj.VertexAttr(0, 3, 5*4, 0)
	m.screenObj.VertexAttr(1, 2, 5*4, 3*4)
	m.screenObj.BindTexture(gl.TEXTURE0, m.textureBuf)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	gl.Viewport(0, 0, 1024, 1024)
	gl.BindFramebuffer(gl.FRAMEBUFFER, m.frameBuf)
	gl.Clear(gl.DEPTH_BUFFER_BIT)

	lightPro := mgl32.Ortho(-10, 10, -10, 10, 1, 7.5)
	lightView := mgl32.LookAtV(m.lightPos, utils.VecZero, utils.VecUp)
	m.depShader.Use()
	m.depShader.SetMat4("uView", &lightView[0])
	m.depShader.SetMat4("uPro", &lightPro[0])

	m.planObj.Draw(m.depShader)
	m.cubeObj.Reset()
	m.cubeObj.Translate(0, 1.5, 0)
	m.cubeObj.Scale(0.5, 0.5, 0.5)
	m.cubeObj.Draw(m.depShader)
	m.cubeObj.Reset()
	m.cubeObj.Translate(2, 0, 1)
	m.cubeObj.Scale(0.5, 0.5, 0.5)
	m.cubeObj.Draw(m.depShader)
	m.cubeObj.Reset()
	m.cubeObj.Translate(-1, 0, 2)
	m.cubeObj.Scale(0.25, 0.25, 0.25)
	m.cubeObj.Draw(m.depShader)

	gl.Viewport(0, 0, 1280, 720)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	m.debugShader.Use()

	m.screenObj.Draw(nil)
	//m.testDraw()
}

func (m *MainGame) testDraw() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	lightPro := mgl32.Ortho(-10, 10, -10, 10, 0.1, 1000)
	lightView := mgl32.LookAtV(m.lightPos, utils.VecZero, utils.VecUp)
	m.depShader.Use()
	m.depShader.SetMat4("uView", &lightView[0])
	m.depShader.SetMat4("uPro", &lightPro[0])

	m.planObj.Draw(m.depShader)
	m.cubeObj.Reset()
	m.cubeObj.Translate(0, 1.5, 0)
	m.cubeObj.Scale(0.5, 0.5, 0.5)
	m.cubeObj.Draw(m.depShader)
	m.cubeObj.Reset()
	m.cubeObj.Translate(2, 0, 1)
	m.cubeObj.Scale(0.5, 0.5, 0.5)
	m.cubeObj.Draw(m.depShader)
	m.cubeObj.Reset()
	m.cubeObj.Translate(-1, 0, 2)
	m.cubeObj.Scale(0.25, 0.25, 0.25)
	m.cubeObj.Draw(m.depShader)
}
