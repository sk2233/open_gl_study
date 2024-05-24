// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3 and OpenGL 4.1 core forward-compatible profile.
package main // import "github.com/go-gl/example/gl41core-cube"

import (
	_ "image/png"
	"log"
	"openGL/utils"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const WindowWidth = 800
const WindowHeight = 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	// 初始化
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// 设置
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// 设置兼容 否则使用向前的API会报错
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// 创建窗口
	window, err := glfw.CreateWindow(WindowWidth, WindowHeight, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// 设置在窗口显示的位置  这样设置就是全屏幕显示   实际点的位置 还是  -1 ~ 1
	gl.Viewport(0, 0, WindowWidth, WindowHeight)

	program := utils.LoadProgram("test3/test")
	ourColor := gl.GetUniformLocation(program, gl.Str("ourColor\x00"))
	transform := gl.GetUniformLocation(program, gl.Str("transform\x00"))

	gl.UseProgram(program)
	testTexture := gl.GetUniformLocation(program, gl.Str("testTexture\x00"))
	gl.Uniform1i(testTexture, 1) // 设置 一次即可

	vertices := []float32{
		0.5, 0.5, 0.0, 1.0, 0, // 右上角
		0.5, -0.5, 0.0, 1.0, 1.0, // 右下角
		-0.5, -0.5, 0.0, 0, 1.0, // 左下角
		-0.5, 0.5, 0.0, 0, 0, // 左上角
	}
	indices := []uint32{
		// 注意索引从0开始!
		// 此例的索引(0,1,2,3)就是顶点数组vertices的下标，
		// 这样可以由下标代表顶点组合成矩形
		0, 1, 3, // 第一个三角形
		1, 2, 3, // 第二个三角形
	}

	// 绑定到  TEXTURE0   TEXTURE1
	_ = utils.LoadTexture("test3/3_1.png", gl.TEXTURE0)
	_ = utils.LoadTexture("test3/2_12.png", gl.TEXTURE1)

	// 为什么直接使用 vao就可以了呢???
	vao := utils.CreateVertexArray() // 里面已经 bind了  下面再操作操作的就是 vao
	_ = utils.BindBuffer(vertices, gl.ARRAY_BUFFER)
	_ = utils.BindBuffer(indices, gl.ELEMENT_ARRAY_BUFFER)

	//  这里指定代码中 layout (location = 0)  的传参
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 4*5, 0)
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 4*5, 4*3)
	gl.EnableVertexAttribArray(1)

	gl.ClearColor(1, 1, 0, 1)
	r := float32(0)

	var angle float32
	// 用户不关闭 就一直刷新
	for !window.ShouldClose() {
		// 必须先清颜色缓存  否则 会闪烁
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// 事件处理
		handleInput(window)
		r += 0.01
		if r > 1 {
			r = 0
		}

		gl.UseProgram(program)
		gl.Uniform4f(ourColor, r, 0, 1, 1)

		angle += 0.02
		// 非常好用的数学矩阵包 !!!!!
		matrix := mgl32.HomogRotate3DZ(angle).Mul4(mgl32.Translate3D(0.2, 0.2, 0.2))
		gl.UniformMatrix4fv(transform, 1, false, &matrix[0])
		gl.BindVertexArray(vao)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

		// 刷Buffer 响应用户事件
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func handleInput(window *glfw.Window) {
	// 处理按键事件
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}
