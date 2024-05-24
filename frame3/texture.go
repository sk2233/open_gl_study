/*
@author: sk
@date: 2023/5/27
*/
package frame3

import "github.com/go-gl/gl/v4.1-core/gl"

type Texture struct {
	width, height float32
	texture       uint32
}

func (t *Texture) initTexture(name string) {
	gl.GenTextures(1, &t.texture)
	gl.BindTexture(gl.TEXTURE_2D, t.texture)

	img := LoadImage(name)
	size := img.Bounds()
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(size.Dx()), int32(size.Dy()), 0, gl.RGBA,
		gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	t.width = float32(size.Dx())
	t.height = float32(size.Dy())

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
}

func (t *Texture) Use() {
	gl.BindTexture(gl.TEXTURE_2D, t.texture)
}

func (t *Texture) GetWidth() float32 {
	return t.width
}

func (t *Texture) GetHeight() float32 {
	return t.height
}

func NewTexture(name string) *Texture {
	res := &Texture{}
	res.initTexture(name)
	return res
}
