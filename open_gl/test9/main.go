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
	program uint32
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
	gl.ActiveTexture(gl.TEXTURE0)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
}

func (m *Main) Size() (int, int) {
	return 1280, 720
}

func (m *Main) GetProgram() uint32 {
	m.program = utils.LoadProgram("test9/test")
	return m.program
}

func (m *Main) Init(window *glfw.Window) {
	vertices := []float32{
		// positions
		-1.0, 1.0, -1.0,
		-1.0, -1.0, -1.0,
		1.0, -1.0, -1.0,
		1.0, -1.0, -1.0,
		1.0, 1.0, -1.0,
		-1.0, 1.0, -1.0,

		-1.0, -1.0, 1.0,
		-1.0, -1.0, -1.0,
		-1.0, 1.0, -1.0,
		-1.0, 1.0, -1.0,
		-1.0, 1.0, 1.0,
		-1.0, -1.0, 1.0,

		1.0, -1.0, -1.0,
		1.0, -1.0, 1.0,
		1.0, 1.0, 1.0,
		1.0, 1.0, 1.0,
		1.0, 1.0, -1.0,
		1.0, -1.0, -1.0,

		-1.0, -1.0, 1.0,
		-1.0, 1.0, 1.0,
		1.0, 1.0, 1.0,
		1.0, 1.0, 1.0,
		1.0, -1.0, 1.0,
		-1.0, -1.0, 1.0,

		-1.0, 1.0, -1.0,
		1.0, 1.0, -1.0,
		1.0, 1.0, 1.0,
		1.0, 1.0, 1.0,
		-1.0, 1.0, 1.0,
		-1.0, 1.0, -1.0,

		-1.0, -1.0, -1.0,
		-1.0, -1.0, 1.0,
		1.0, -1.0, -1.0,
		1.0, -1.0, -1.0,
		-1.0, -1.0, 1.0,
		1.0, -1.0, 1.0,
	}
	var indices []uint32
	for i := 0; i < 36; i++ {
		indices = append(indices, uint32(i))
	}
	utils.CreateVertexArray() // 里面已经 bind了  下面再操作操作的就是 vao
	utils.BindBuffer(vertices, gl.ARRAY_BUFFER)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 4*3, 0)
	gl.EnableVertexAttribArray(0)
	utils.LoadCubeTexture("test9/img.png")
}
