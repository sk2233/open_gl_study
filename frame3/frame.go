/*
@author: sk
@date: 2023/5/7
*/
package frame3

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func Run(game IGame, width, height int) {
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
	window, err := glfw.CreateWindow(width, height, "Game2D", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
	window.SetSizeCallback(onSizeChange)               // 必须回掉设置
	gl.Enable(gl.BLEND)                                // 打开混合函数
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // 混合全部颜色
	InitData()
	game.Init(window)
	for !window.ShouldClose() {
		glfw.PollEvents()
		// 必须先清颜色缓存  否则 会闪烁
		gl.Clear(gl.COLOR_BUFFER_BIT)
		game.Update(window)
		game.Draw()
		// 刷Buffer 响应用户事件
		window.SwapBuffers()
	}
}

func onSizeChange(w *glfw.Window, width int, height int) {
	// 设置在窗口显示的位置  这样设置就是全屏幕显示   实际点的位置 还是  -1 ~ 1
	gl.Viewport(0, 0, int32(width), int32(height))
}
