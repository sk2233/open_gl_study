/*
@author: sk
@date: 2023/5/14
*/
package utils

import "github.com/go-gl/mathgl/mgl32"

func GetDefaultPerspective() mgl32.Mat4 {
	return mgl32.Perspective(45, 1280.0/720.0, 0.1, 100)
}

func GetDefaultOrtho2D() mgl32.Mat4 {
	return mgl32.Ortho2D(-640, 640, 360, -360)
}

var (
	VecZero = mgl32.Vec3{0, 0, 0}
	VecUp   = mgl32.Vec3{0, 1, 0}
)
