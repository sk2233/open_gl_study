/*
@author: sk
@date: 2023/5/7
*/
package frame

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func Run(game IGame) {
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
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	// 创建窗口
	width, height := game.Size()
	window, err := glfw.CreateWindow(width, height, "Game", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
	// 初始化shader
	program := game.GetProgram()
	modelLoc := gl.GetUniformLocation(program, gl.Str("aModel\x00"))
	viewLoc := gl.GetUniformLocation(program, gl.Str("aView\x00"))
	projectionLoc := gl.GetUniformLocation(program, gl.Str("aProjection\x00"))
	window.SetSizeCallback(onSizeChange) // 必须回掉设置
	gl.ClearColor(0, 0, 0, 1)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.Enable(gl.CULL_FACE)
	//gl.CullFace(gl.FRONT)
	//gl.FrontFace(gl.CCW)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	//gl.BlendEquation(gl.FUNC_SUBTRACT)
	//gl.BlendFuncSeparate(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA, gl.ONE, gl.ZERO)
	game.Init(window)
	camera := NewCamera()
	// 正交/透视矩阵   实现相机效果
	projection := mgl32.Perspective(45, float32(width)/float32(height), 0.1, 100)
	gl.UseProgram(program)
	gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])
	//gl.Enable(gl.STENCIL_TEST)
	//gl.StencilOp(gl.KEEP, gl.KEEP, gl.REPLACE)
	for !window.ShouldClose() {
		// 必须先清颜色缓存  否则 会闪烁
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		game.Update(window, camera)
		view := camera.GetView()
		gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])
		game.Draw(modelLoc, camera)
		// 刷Buffer 响应用户事件
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func onSizeChange(w *glfw.Window, width int, height int) {
	// 设置在窗口显示的位置  这样设置就是全屏幕显示   实际点的位置 还是  -1 ~ 1
	gl.Viewport(0, 0, int32(width), int32(height))
}
