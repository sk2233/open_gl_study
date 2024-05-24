/*
@author: sk
@date: 2023/6/23
*/
package main

import (
	frame "openGL/frame2"
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type PointObj struct {
	vao   uint32
	count int32
	model mgl32.Mat4
}

func (o *PointObj) VertexAttr(index uint32, count int32, step int32, offset uintptr) {
	gl.BindVertexArray(o.vao)
	gl.EnableVertexAttribArray(index)
	gl.VertexAttribPointerWithOffset(index, count, gl.FLOAT, false, step, offset)
}

// 对应的 shader 必须 有 model 参数
func (o *PointObj) Draw(shader *frame.Shader) {
	gl.BindVertexArray(o.vao)
	if shader != nil {
		shader.SetMat4("uModel", &o.model[0])
	}
	gl.DrawArrays(gl.POINTS, 0, o.count) // 可以考虑复用下面的
}

func (o *PointObj) Translate(x, y, z float32) {
	o.model = o.model.Mul4(mgl32.Translate3D(x, y, z))
}

func (o *PointObj) Scale(x, y, z float32) {
	o.model = o.model.Mul4(mgl32.Scale3D(x, y, z))
}

func (o *PointObj) Reset() {
	o.model = mgl32.Ident4()
}

func (o *PointObj) GetModel() mgl32.Mat4 {
	return o.model
}

func (o *PointObj) SetModel(model mgl32.Mat4) {
	o.model = model
}

func NewPointObj(vs []float32, count int32) *PointObj {
	vao := utils.CreateVertexArray()
	utils.BindBuffer(vs, gl.ARRAY_BUFFER)
	return &PointObj{vao: vao, count: count, model: mgl32.Ident4()}
}
