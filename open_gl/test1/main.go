// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3 and OpenGL 4.1 core forward-compatible profile.
package main // import "github.com/go-gl/example/gl41core-cube"

import (
	_ "image/png"
	"log"
	frame "openGL/frame2"
	"openGL/utils"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
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
	// 设置窗口变化回调
	//window.SetSizeCallback(func(w *glfw.Window, width int, height int) {
	//
	//})

	vShader := utils.LoadShader("test1/v.vert", gl.VERTEX_SHADER)
	fShader := utils.LoadShader("test1/f.frag", gl.FRAGMENT_SHADER)
	program := gl.CreateProgram()
	gl.AttachShader(program, vShader)
	gl.AttachShader(program, fShader)
	gl.LinkProgram(program)
	gl.UseProgram(program)
	gl.DeleteShader(vShader)
	gl.DeleteShader(fShader)

	vertices := []float32{
		0.5, 0.5, 0.0, // 右上角
		0.5, -0.5, 0.0, // 右下角
		-0.5, -0.5, 0.0, // 左下角
	}

	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 4*3, 0)
	gl.EnableVertexAttribArray(0)

	counter := frame.NewCounter()

	// 用户不关闭 就一直刷新
	for !window.ShouldClose() {
		// 必须先清颜色缓存  否则 会闪烁
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.ClearColor(1, 1, 0, 1)

		// 事件处理
		handleInput(window)

		gl.UseProgram(program)
		gl.BindVertexArray(vao)
		gl.DrawArraysInstanced(gl.TRIANGLES, 0, 3, 10000)
		counter.Update()

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
