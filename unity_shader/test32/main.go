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
	// 天空盒必须使用 小于等于  因为天空盒 会固定最后渲染 齐次化后位于  zBuffer 1 位于齐次空间的边缘，若是超出齐次空间将取消绘制
	gl.DepthFunc(gl.LEQUAL)

	pro := utils.GetDefaultPerspective()

	m.skyShader = frame.NewShader("test32/sky")
	m.skyShader.Use()
	m.skyShader.Set1i("uSkybox", 0)
	m.skyShader.SetMat4("uProjection", &pro[0])

	m.cubeShader = frame.NewShader("test32/obj")
	m.cubeShader.Use()
	m.cubeShader.Set1i("uSkybox", 0)
	m.cubeShader.Set1i("uNor", 1)
	m.cubeShader.Set1i("uTex", 2)
	m.cubeShader.SetMat4("uProjection", &pro[0])

	temp := utils.LoadCubeTexture("test32/sky/posx.jpg", "test32/sky/negx.jpg", "test32/sky/posy.jpg", "test32/sky/negy.jpg",
		"test32/sky/posz.jpg", "test32/sky/negz.jpg")

	m.cubeObj = frame.NewObj(cubeVs, 36)
	m.cubeObj.VertexAttr(0, 3, 5*4, 0)
	m.cubeObj.VertexAttr(1, 2, 5*4, 3*4)
	m.cubeObj.BindTextureCube(gl.TEXTURE0, temp)
	m.cubeObj.BindTexture(gl.TEXTURE1, utils.LoadTexture("test32/normal.jpg"))
	m.cubeObj.BindTexture(gl.TEXTURE2, utils.LoadTexture("test32/diffuse.jpg"))

	m.skyObj = frame.NewObj(skyVs, 36)
	m.skyObj.VertexAttr(0, 3, 3*4, 0)
	m.skyObj.BindTextureCube(gl.TEXTURE0, temp)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	m.cubeShader.Use()
	view := m.camera.GetView()
	m.cubeShader.SetMat4("uView", &view[0])
	eyePos := m.camera.GetPos()
	m.cubeShader.Set3fv("uEyePos", &eyePos[0])
	m.cubeObj.Draw(m.cubeShader)

	m.skyShader.Use()
	view = view.Mat3().Mat4()
	m.skyShader.SetMat4("uView", &view[0])
	m.skyObj.Draw(nil)
}
