/*
@author: sk
@date: 2023/5/7
*/
package frame3

import "github.com/go-gl/glfw/v3.3/glfw"

type IGame interface {
	Init(window *glfw.Window)
	Update(window *glfw.Window)
	Draw()
}

type IRect interface {
	GetPos() complex64
	GetSize() complex64
}
