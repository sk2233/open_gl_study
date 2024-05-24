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
)

func main() {
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	camera                                *frame.Camera
	objShader, screenShader               *frame.Shader
	cubeObj, planObj, screenObj           *frame.Obj
	cubeTexture, planTexture, buffTexture uint32
	buff                                  uint32
}

func NewMainGame() *MainGame {
	return &MainGame{camera: frame.NewCamera()}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.objShader = frame.NewShader("test10/obj")
	m.screenShader = frame.NewShader("test10/screen")
	cubeTexture := utils.LoadTexture("test10/cube.png")
	planTexture := utils.LoadTexture("test10/plan.png")
	m.cubeObj = frame.NewObj(cubeVs, 36)
	m.cubeObj.VertexAttr(0, 3, 5*4, 0)
	m.cubeObj.VertexAttr(1, 2, 5*4, 3*4)
	m.cubeObj.BindTexture(gl.TEXTURE0, cubeTexture)
	m.planObj = frame.NewObj(planeVs, 6)
	m.planObj.VertexAttr(0, 3, 5*4, 0)
	m.planObj.VertexAttr(1, 2, 5*4, 3*4)
	m.planObj.BindTexture(gl.TEXTURE0, planTexture)
	m.buff, m.buffTexture = utils.NewTextureBuff(1280, 720)
	m.screenObj = frame.NewObj(quadVs, 6)
	m.screenObj.VertexAttr(0, 2, 4*4, 0)
	m.screenObj.VertexAttr(1, 2, 4*4, 2*4)
	m.screenObj.BindTexture(gl.TEXTURE0, m.buffTexture)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, m.buff) // 先渲染到 buff
	//gl.Enable(gl.DEPTH_TEST)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) // 因为切换绘制的原因  原来的清除失效
	// 绘制普通对象
	view := m.camera.GetView()
	projection := utils.GetDefaultPerspective()
	m.objShader.Use()
	m.objShader.SetMat4("view", &view[0])
	m.objShader.SetMat4("projection", &projection[0])
	m.cubeObj.Reset()
	m.cubeObj.Translate(-1, 0, -1)
	m.cubeObj.Draw(m.objShader)
	m.cubeObj.Reset()
	m.cubeObj.Translate(1, 0, 1)
	m.cubeObj.Draw(m.objShader)
	//m.planObj.Draw(m.objShader)
	// 绘制缓存
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	//gl.Disable(gl.DEPTH_TEST)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	m.screenShader.Use()
	m.screenObj.Draw(nil)
}
