/*
@author: sk
@date: 2023/5/7
*/
package frame

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	pos, dir   mgl32.Vec3
	dirX, dirY float32 // 不能绕Z旋转
}

func NewCamera() *Camera {
	return &Camera{pos: mgl32.Vec3{5, 5, 5}, dir: mgl32.Vec3{-1, -1, -1}}
}

func (c *Camera) GetView() mgl32.Mat4 {
	return mgl32.LookAtV(c.pos, c.pos.Add(c.dir), VecUp)
}

func (c *Camera) TranslateX(value float32) {
	c.pos = c.pos.Add(c.dir.Cross(VecUp).Normalize().Mul(value))
}

func (c *Camera) TranslateY(value float32) {
	c.pos = c.pos.Add(c.dir.Cross(VecRight).Normalize().Mul(value))
}

func (c *Camera) TranslateZ(value float32) {
	c.pos = c.pos.Add(c.dir.Normalize().Mul(value))
}

func (c *Camera) RotateX(value float32) { // 左右看 沿着 Y轴
	//c.dirY = Clamp(c.dirY+value, -math.Pi/2, math.Pi/2)
	//c.updateDir()
	c.dir = mgl32.Rotate3DY(value).Mul3x1(c.dir)
}

func (c *Camera) RotateY(value float32) { // 上下看 沿着  X轴
	//c.dirX = Clamp(c.dirX+value, -math.Pi/2, math.Pi/2)
	//c.updateDir()
	c.dir = mgl32.Rotate3DX(value).Mul3x1(c.dir)
}

func (c *Camera) updateDir() {
	c.dir = mgl32.Rotate3DX(c.dirX).Mul3(mgl32.Rotate3DY(c.dirY)).Mul3x1(VecFront)
}

func (c *Camera) GetPos() mgl32.Vec3 {
	return c.pos
}
