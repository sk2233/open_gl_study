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
	camera                *frame.Camera
	skyShader, cubeShader *frame.Shader
	skyObj, cubeObj       *frame.Obj
}

func NewMainGame() *MainGame {
	return &MainGame{camera: frame.NewCamera()}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.skyShader = frame.NewShader("test11/sky")
	m.skyShader.Use()
	m.skyShader.Set1i("skybox", 0)
	m.cubeShader = frame.NewShader("test11/obj")
	m.cubeShader.Use()
	m.cubeShader.Set1i("skybox", 0)

	m.cubeObj = frame.NewObj(cubeVs, 36)
	m.cubeObj.VertexAttr(0, 3, 6*4, 0)
	m.cubeObj.VertexAttr(1, 3, 6*4, 3*4)
	temp := utils.LoadCubeTexture("test11/sky/right.jpg", "test11/sky/left.jpg", "test11/sky/top.jpg",
		"test11/sky/bottom.jpg", "test11/sky/front.jpg", "test11/sky/back.jpg")
	m.cubeObj.BindTextureCube(gl.TEXTURE0, temp)

	m.skyObj = frame.NewObj(skyVs, 36)
	m.skyObj.VertexAttr(0, 3, 3*4, 0)
	temp = utils.LoadCubeTexture("test11/sky/right.jpg", "test11/sky/left.jpg", "test11/sky/top.jpg",
		"test11/sky/bottom.jpg", "test11/sky/front.jpg", "test11/sky/back.jpg")
	m.skyObj.BindTextureCube(gl.TEXTURE0, temp)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	gl.DepthFunc(gl.LESS)
	m.cubeShader.Use()
	view := m.camera.GetView()
	projection := utils.GetDefaultPerspective()
	m.cubeShader.SetMat4("view", &view[0])
	m.cubeShader.SetMat4("projection", &projection[0])
	cameraPos := m.camera.GetPos()
	m.cubeShader.Set3fv("cameraPos", &cameraPos[0])
	m.cubeObj.Draw(m.cubeShader)

	gl.DepthFunc(gl.LEQUAL)
	m.skyShader.Use()
	view = m.camera.GetView().Mat3().Mat4()
	m.skyShader.SetMat4("view", &view[0])
	m.skyShader.SetMat4("projection", &projection[0])
	m.skyObj.Draw(nil)
}
