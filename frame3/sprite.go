/*
@author: sk
@date: 2023/5/27
*/
package frame3

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Sprite struct {
	Shader  *Shader
	Texture *Texture
}

func NewSprite(shader *Shader, texture *Texture) *Sprite {
	return &Sprite{Shader: shader, Texture: texture}
}

func (s *Sprite) Draw(pos, scale complex64, angle float32, color mgl32.Vec4) {
	s.Shader.Use()
	gl.ActiveTexture(gl.TEXTURE0) // 默认只激活第一张贴图即可
	model := mgl32.Translate3D(real(pos), imag(pos), 0).Mul4(mgl32.HomogRotate3DZ(angle)).
		Mul4(mgl32.Scale3D(real(scale)*s.Texture.GetWidth(), imag(scale)*s.Texture.GetHeight(), 1))
	s.Shader.SetMat4("uModel", &model[0])
	s.Shader.SetMat4("uProj", &ProjMat[0])
	s.Shader.Set4fv("uColor", &color[0])
	s.Texture.Use()
	gl.BindVertexArray(RectVao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}
