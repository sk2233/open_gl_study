/*
@author: sk
@date: 2023/5/7
*/
package frame

import "github.com/go-gl/glfw/v3.3/glfw"

type IGame interface {
	Init(window *glfw.Window)
	Update(window *glfw.Window, camera *Camera)
	Draw(int32, *Camera)
	Size() (int, int)
	GetProgram() uint32 // 必须支持一个着色器  至少支持传入 mvp
}
