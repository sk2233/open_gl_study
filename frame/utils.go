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
