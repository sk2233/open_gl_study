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
	vao     uint32
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
	angle += 0.01
	gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])
	gl.DrawElements(gl.TRIANGLES, 36, gl.UNSIGNED_INT, nil)
}

var (
	angle = float32(0.0)
)

func (m *Main) Size() (int, int) {
	return 1280, 720
}

func (m *Main) GetProgram() uint32 {
	m.program = utils.LoadProgram("test8/test")
	return m.program
}

func (m *Main) Init(window *glfw.Window) {
	vertices := []float32{
		// positions        // texture coords
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}
	var indices []uint32
	for i := 0; i < 36; i++ {
		indices = append(indices, uint32(i))
	}
	m.vao = utils.CreateVertexArray() // 里面已经 bind了  下面再操作操作的就是 vao
	_ = utils.BindBuffer(vertices, gl.ARRAY_BUFFER)
	_ = utils.BindBuffer(indices, gl.ELEMENT_ARRAY_BUFFER)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 4*5, 0)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 4*5, 4*3)
	gl.EnableVertexAttribArray(1)
	utils.LoadTexture("test8/img_1.png", gl.TEXTURE0)
	// 创建一个帧缓存
	var fbo uint32
	gl.GenFramebuffers(1, &fbo)
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo) // 读写都会搞到这个buf上  READ_FRAMEBUFFER 可也以指定只读只写
	//gl.BindFramebuffer(gl.FRAMEBUFFER,0)// 默认是渲染到 平面0上
	//gl.DeleteFramebuffers(1, &fbo) // 删除 表面
	// 创建一个纹理 附件 无需设置初始数据
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, 1280, 720, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	// 绑定附件   COLOR_ATTACHMENT0  这里仅绑定一个附件  TEXTURE_2D是指纹理类型  level多级纹理
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texture, 0)
	// 创建一个渲染缓冲
	var rbo uint32
	gl.GenRenderbuffers(1, &rbo)
	gl.BindRenderbuffer(gl.RENDERBUFFER, rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, 1280, 720)
	//gl.BindRenderbuffer(gl.RENDERBUFFER, 0)// 绑定的原对象
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, rbo)

}

func Test11() {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	// 深度与模拟信息可以共用一个  深度  24 + 模版 8  = 32
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH24_STENCIL8, 1280, 720, 0,
		gl.DEPTH_STENCIL, gl.UNSIGNED_INT_24_8, nil)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.TEXTURE_2D, texture, 0)
}
