/*
@author: sk
@date: 2023/5/7
*/
package frame

import "github.com/go-gl/glfw/v3.3/glfw"

func Clamp(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func PressKey(window *glfw.Window, key glfw.Key) bool {
	return window.GetKey(key) == glfw.Press
}

func GetAxis(window *glfw.Window, min, max glfw.Key) float32 {
	if PressKey(window, min) {
		return -1
	}
	if PressKey(window, max) {
		return 1
	}
	return 0
}

func ApplyInput(camera *Camera, window *glfw.Window) {
	offsetX := GetAxis(window, glfw.KeyA, glfw.KeyD)
	if offsetX != 0 {
		camera.TranslateX(offsetX * 0.1)
	}
	offsetY := GetAxis(window, glfw.KeyE, glfw.KeyQ)
	if offsetY != 0 {
		camera.TranslateY(offsetY * 0.1)
	}
	offsetZ := GetAxis(window, glfw.KeyS, glfw.KeyW)
	if offsetZ != 0 {
		camera.TranslateZ(offsetZ * 0.1)
	}
	rotateX := GetAxis(window, glfw.KeyRight, glfw.KeyLeft)
	if rotateX != 0 {
		camera.RotateX(rotateX * 0.01)
	}
	rotateY := GetAxis(window, glfw.KeyDown, glfw.KeyUp)
	if rotateY != 0 {
		camera.RotateY(rotateY * 0.01)
	}
}
