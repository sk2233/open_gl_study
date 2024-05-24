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
	vao                                                                   uint32
	ambientLight                                                          int32
	lightPos                                                              int32
	eyePos                                                                int32
	materialAmbient, materialDiffuse, materialSpecular, materialShininess int32
	specTexture                                                           int32
	lightDir                                                              int32
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
	gl.Uniform3f(m.ambientLight, 1, 1, 1)
	gl.Uniform3f(m.lightPos, 5, 5, 5)
	pos := camera.GetPos()
	gl.Uniform3fv(m.eyePos, 1, &pos[0])
	// 设置材质参数
	gl.Uniform3f(m.materialAmbient, 1, 1, 1)
	gl.Uniform3f(m.materialSpecular, 0, 1, 0)
	gl.Uniform3f(m.materialDiffuse, 0, 0, 1)
	gl.Uniform1f(m.materialShininess, 64)
	gl.Uniform1i(m.specTexture, 1)
	gl.Uniform3f(m.lightDir, -5, -5, -5)
	gl.DrawElements(gl.TRIANGLES, 36, gl.UNSIGNED_INT, nil)
	for i := 0; i < 5; i++ {
		angle += 0.01
		model = mgl32.Translate3D(float32(i-2)*1.5, -2, 0).Mul4(mgl32.HomogRotate3DY(angle))
		gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])
		gl.DrawElements(gl.TRIANGLES, 36, gl.UNSIGNED_INT, nil)
	}
}

var (
	angle = float32(0.0)
)

func (m *Main) Size() (int, int) {
	return 1280, 720
}

func (m *Main) GetProgram() uint32 {
	m.program = utils.LoadProgram("test6/test")
	m.ambientLight = gl.GetUniformLocation(m.program, gl.Str("aAmbient\x00"))
	m.lightPos = gl.GetUniformLocation(m.program, gl.Str("aLightPos\x00"))
	m.eyePos = gl.GetUniformLocation(m.program, gl.Str("aEyePos\x00"))
	m.materialSpecular = gl.GetUniformLocation(m.program, gl.Str("aMaterial.specular\x00"))
	m.materialShininess = gl.GetUniformLocation(m.program, gl.Str("aMaterial.shininess\x00"))
	m.materialDiffuse = gl.GetUniformLocation(m.program, gl.Str("aMaterial.diffuse\x00"))
	m.materialAmbient = gl.GetUniformLocation(m.program, gl.Str("aMaterial.ambient\x00"))
	m.specTexture = gl.GetUniformLocation(m.program, gl.Str("specTexture\x00"))
	m.lightDir = gl.GetUniformLocation(m.program, gl.Str("aLightDir\x00"))
	return m.program
}

func (m *Main) Init(window *glfw.Window) {
	vertices := []float32{
		// positions       // normals        // texture coords
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,
		0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,
		0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,

		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, -0.5, -1.0, 0.0, 0.0, 1.0, 1.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, -1.0, 0.0, 0.0, 0.0, 0.0,
		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 0.0, 0.0, 1.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0, 0.0, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
	}
	var indices []uint32
	for i := 0; i < 36; i++ {
		indices = append(indices, uint32(i))
	}
	m.vao = utils.CreateVertexArray() // 里面已经 bind了  下面再操作操作的就是 vao
	_ = utils.BindBuffer(vertices, gl.ARRAY_BUFFER)
	_ = utils.BindBuffer(indices, gl.ELEMENT_ARRAY_BUFFER)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 4*8, 0)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 4*8, 4*3)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, 4*8, 4*6)
	gl.EnableVertexAttribArray(2)
	utils.LoadTexture("test6/img.png", gl.TEXTURE0)
	utils.LoadTexture("test6/img_1.png", gl.TEXTURE1)
}
