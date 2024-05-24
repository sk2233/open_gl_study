/*
@author: sk
@date: 2023/5/27
*/
package frame

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	rectVs = []float32{
		// 位置     // 纹理
		0.0, 1.0, 0.0, 1.0,
		1.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 0.0,

		0.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 0.0,
	}
	screenVs = []float32{
		// Pos        // Tex
		-1.0, -1.0, 0.0, 0.0,
		1.0, 1.0, 1.0, 1.0,
		-1.0, 1.0, 0.0, 1.0,

		-1.0, -1.0, 0.0, 0.0,
		1.0, -1.0, 1.0, 0.0,
		1.0, 1.0, 1.0, 1.0,
	}
)

var (
	RectVao   uint32
	ScreenVao uint32
	ProjMat   mgl32.Mat4
)

var (
	ColorWhite = mgl32.Vec4{1, 1, 1, 1}
)

func InitData() {
	// RectVao
	gl.GenVertexArrays(1, &RectVao)
	gl.BindVertexArray(RectVao)
	// bindBuff
	var rectBuff uint32
	gl.GenBuffers(1, &rectBuff)
	gl.BindBuffer(gl.ARRAY_BUFFER, rectBuff)
	gl.BufferData(gl.ARRAY_BUFFER, len(rectVs)*4, gl.Ptr(rectVs), gl.STATIC_DRAW)
	// setAttr
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 2, gl.FLOAT, false, 4*4, 0)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 4*4, 2*4)
	// ScreenVao
	gl.GenVertexArrays(1, &ScreenVao)
	gl.BindVertexArray(ScreenVao)
	// bindBuff
	var screenBuff uint32
	gl.GenBuffers(1, &screenBuff)
	gl.BindBuffer(gl.ARRAY_BUFFER, screenBuff)
	gl.BufferData(gl.ARRAY_BUFFER, len(screenVs)*4, gl.Ptr(screenVs), gl.STATIC_DRAW)
	// setAttr
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 2, gl.FLOAT, false, 4*4, 0)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 4*4, 2*4)
	// ProjMat
	ProjMat = mgl32.Ortho2D(0, 1280, 720, 0)
}

type Rect struct {
	pos, size complex64
}

func (r *Rect) GetPos() complex64 {
	return r.pos
}

func (r *Rect) GetSize() complex64 {
	return r.size
}
