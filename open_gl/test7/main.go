/*
@author: sk
@date: 2023/5/7
*/
package main

import (
	"openGL/frame"
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	frame.Run(&Main{})
}

type Main struct {
	program                                                               uint32
	ambientLight                                                          int32
	lightPos                                                              int32
	eyePos                                                                int32
	materialAmbient, materialDiffuse, materialSpecular, materialShininess int32
	count                                                                 int32
}

func (m *Main) Update(window *glfw.Window, camera *frame.Camera) {
	offsetX := frame.GetAxis(window, glfw.KeyA, glfw.KeyD)
	if offsetX != 0 {
		camera.TranslateX(offsetX * 0.1)
	}
	offsetY := frame.GetAxis(window, glfw.KeyE, glfw.KeyQ)
	if offsetY != 0 {
		camera.TranslateY(offsetY * 0.1)
	}
	offsetZ := frame.GetAxis(window, glfw.KeyS, glfw.KeyW)
	if offsetZ != 0 {
		camera.TranslateZ(offsetZ * 0.1)
	}
	rotateX := frame.GetAxis(window, glfw.KeyRight, glfw.KeyLeft)
	if rotateX != 0 {
		camera.RotateX(rotateX * 0.01)
	}
	rotateY := frame.GetAxis(window, glfw.KeyDown, glfw.KeyUp)
	if rotateY != 0 {
		camera.RotateY(rotateY * 0.01)
	}
}

func (m *Main) Draw(modelLoc int32, camera *frame.Camera) {
	model := mgl32.Translate3D(0, 0, 0)
	gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])
	gl.Uniform3f(m.ambientLight, 0.5, 0.5, 0.5)
	gl.Uniform3f(m.lightPos, 5, 5, 5)
	pos := camera.GetPos()
	gl.Uniform3fv(m.eyePos, 1, &pos[0])
	// 设置材质参数
	gl.Uniform3f(m.materialAmbient, 1, 1, 1)
	gl.Uniform3f(m.materialSpecular, 0, 1, 0)
	gl.Uniform3f(m.materialDiffuse, 0, 0, 1)
	gl.Uniform1f(m.materialShininess, 64)
	gl.DrawElements(gl.TRIANGLES, m.count, gl.UNSIGNED_INT, nil)
}

func (m *Main) Size() (int, int) {
	return 1280, 720
}

func (m *Main) GetProgram() uint32 {
	m.program = utils.LoadProgram("test7/test")
	m.ambientLight = gl.GetUniformLocation(m.program, gl.Str("aAmbient\x00"))
	m.lightPos = gl.GetUniformLocation(m.program, gl.Str("aLightPos\x00"))
	m.eyePos = gl.GetUniformLocation(m.program, gl.Str("aEyePos\x00"))
	m.materialSpecular = gl.GetUniformLocation(m.program, gl.Str("aMaterial.specular\x00"))
	m.materialShininess = gl.GetUniformLocation(m.program, gl.Str("aMaterial.shininess\x00"))
	m.materialDiffuse = gl.GetUniformLocation(m.program, gl.Str("aMaterial.diffuse\x00"))
	m.materialAmbient = gl.GetUniformLocation(m.program, gl.Str("aMaterial.ambient\x00"))
	return m.program
}

func (m *Main) Init(window *glfw.Window) {
	vertices := utils.LoadObj("test7/reimu.obj")
	m.count = int32(len(vertices) / 8)
	indices := make([]uint32, m.count)
	for i := 0; i < len(indices); i++ {
		indices[i] = uint32(i)
	}
	_ = utils.CreateVertexArray() // 里面已经 bind了  下面再操作操作的就是 vao
	_ = utils.BindBuffer(vertices, gl.ARRAY_BUFFER)
	_ = utils.BindBuffer(indices, gl.ELEMENT_ARRAY_BUFFER)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 4*8, 0)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 4*8, 4*3)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(2, 3, gl.FLOAT, false, 4*8, 4*5)
	gl.EnableVertexAttribArray(2)
	utils.LoadTexture("test7/reimu.png")
}
