/*
@author: sk
@date: 2023/5/14
*/
package frame

import (
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Obj struct {
	vao          uint32
	count        int32
	textures     map[uint32]uint32
	texture3Ds   map[uint32]uint32
	model        mgl32.Mat4
	textureCubes map[uint32]uint32
}

func (o *Obj) VertexAttr(index uint32, count int32, step int32, offset uintptr) {
	gl.BindVertexArray(o.vao)
	gl.EnableVertexAttribArray(index)
	gl.VertexAttribPointerWithOffset(index, count, gl.FLOAT, false, step, offset)
}

// 对应的 shader 必须 有 model 参数
func (o *Obj) Draw(shader *Shader) {
	gl.BindVertexArray(o.vao)
	for index, texture := range o.textures {
		gl.ActiveTexture(index)
		gl.BindTexture(gl.TEXTURE_2D, texture)
	}
	for index, textureCube := range o.textureCubes {
		gl.ActiveTexture(index)
		gl.BindTexture(gl.TEXTURE_CUBE_MAP, textureCube)
	}
	if shader != nil {
		shader.SetMat4("uModel", &o.model[0])
		mIT := o.model.Inv().Transpose()
		shader.SetMat4("uModelIT", &mIT[0])
	}
	gl.DrawArrays(gl.TRIANGLES, 0, o.count) // 可以考虑复用下面的
}

func (o *Obj) DrawInstanced(shader *Shader, count int32) {
	gl.BindVertexArray(o.vao)
	for index, texture := range o.textures {
		gl.ActiveTexture(index)
		gl.BindTexture(gl.TEXTURE_2D, texture)
	}
	for index, textureCube := range o.textureCubes {
		gl.ActiveTexture(index)
		gl.BindTexture(gl.TEXTURE_CUBE_MAP, textureCube)
	}
	for index, texture3D := range o.texture3Ds {
		gl.ActiveTexture(index)
		gl.BindTexture(gl.TEXTURE_3D, texture3D)
	}
	if shader != nil {
		shader.SetMat4("uModel", &o.model[0])
	}
	gl.DrawArrays(gl.TRIANGLES, 0, o.count)
	gl.DrawArraysInstanced(gl.TRIANGLES, 0, o.count, count)
}

func (o *Obj) BindTexture(index uint32, texture uint32) {
	o.textures[index] = texture
}

func (o *Obj) BindTexture3D(index uint32, texture uint32) {
	o.texture3Ds[index] = texture
}

func (o *Obj) Translate(x, y, z float32) {
	o.model = o.model.Mul4(mgl32.Translate3D(x, y, z))
}

func (o *Obj) Scale(x, y, z float32) {
	o.model = o.model.Mul4(mgl32.Scale3D(x, y, z))
}

func (o *Obj) Reset() {
	o.model = mgl32.Ident4()
}

func (o *Obj) BindTextureCube(index uint32, textureCube uint32) {
	o.textureCubes[index] = textureCube
}

func (o *Obj) GetModel() mgl32.Mat4 {
	return o.model
}

func NewObj(vs []float32, count int32) *Obj {
	vao := utils.CreateVertexArray()
	utils.BindBuffer(vs, gl.ARRAY_BUFFER)
	return &Obj{vao: vao, count: count, textures: make(map[uint32]uint32),
		textureCubes: make(map[uint32]uint32), model: mgl32.Ident4(), texture3Ds: make(map[uint32]uint32)}
}
