/*
@author: sk
@date: 2023/5/7
*/
package frame

import (
	"fmt"
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
	// 设置兼容 否则使用向前的API会报错
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	// 创建窗口
	window, err := glfw.CreateWindow(width, height, "Game", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.MakeContextCurrent()
	window.SetSizeCallback(onSizeChange) // 必须回掉设置
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.ClearColor(0.3, 0.3, 0.3, 1)
	gl.Enable(gl.DEPTH_TEST)
	InitData()
	game.Init(window)
	for !window.ShouldClose() {
		// 必须先清颜色缓存  否则 会闪烁
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) // 正常模式
		//gl.Clear(gl.DEPTH_BUFFER_BIT) // 不清除模式
		game.Update(window)
		game.Draw()
		// 刷Buffer 响应用户事件
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func onSizeChange(window *glfw.Window, width int, height int) {
	// 设置在窗口显示的位置  这样设置就是全屏幕显示   实际点的位置 还是  -1 ~ 1
	gl.Viewport(0, 0, int32(width), int32(height))
	fmt.Printf("width:%d,height:%d", width, height)
}
